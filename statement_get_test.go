package monoacquiring

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetStatement(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/statement", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "application/json", req.Header.Get("Accept"))
		assert.Equal(t, "x-api-token-test", req.Header.Get("X-Token"))
		assert.Equal(t, "cms-test", req.Header.Get("X-Cms"))
		assert.Equal(t, "0.0.1", req.Header.Get("X-Cms-Version"))
		assert.Equal(t, "from=1755692087", req.URL.RawQuery)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "list": [
    {
      "invoiceId": "2205175v4MfatvmUL2oR",
      "status": "success",
      "maskedPan": "444403******1902",
      "date": null,
      "paymentScheme": "full",
      "amount": 4200,
      "profitAmount": 4100,
      "ccy": 980,
      "approvalCode": "662476",
      "rrn": "060189181768",
      "reference": "84d0070ee4e44667b31371d8f8813947",
      "shortQrId": "OBJE",
      "destination": "Покупка щастя",
      "cancelList": [
        {
          "amount": 4200,
          "ccy": 980,
          "date": null,
          "approvalCode": "662476",
          "rrn": "060189181768",
          "maskedPan": "444403******1902"
        }
      ]
    }
  ]
}`)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewClient(
		Config{APIKey: "x-api-token-test", BaseURL: srv.URL, CMS: "cms-test", CMSVersion: "0.0.1"},
		srv.Client(),
		nil,
	)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	res, err := client.GetStatement(context.Background(), GetStatementRequest{
		From: time.Unix(1755692087, 0),
	})

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.List, 1)

	assert.Equal(t, "2205175v4MfatvmUL2oR", res.List[0].InvoiceID)
	assert.Equal(t, "444403******1902", res.List[0].MaskedPan)
	assert.True(t, res.List[0].Status.IsSuccess())
	assert.True(t, res.List[0].PaymentScheme.IsFull())
	assert.Len(t, res.List[0].CancelList, 1)
}

func TestGetStatement_Validation(t *testing.T) {
	client, err := NewClient(Config{APIKey: "test", BaseURL: DefaultBaseURL}, nil, nil)

	assert.NoError(t, err)

	res, err := client.GetStatement(context.Background(), GetStatementRequest{})

	assert.Error(t, err)
	assert.Nil(t, res)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"GetStatementRequest.From": "required",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

func TestGetStatement_NetworkError(t *testing.T) {

	from := time.Unix(1755692087, 0)
	to := time.Unix(1755692314, 0)

	tests := map[string]struct {
		From       time.Time
		Err        error
		Code       *string
		To         *time.Time
		ErrCode    string
		ErrMessage string
		RawQuery   string
		StatusCode int
	}{
		"bad request": {
			ErrCode:    "BAD_REQUEST",
			ErrMessage: "empty 'invoiceId'",
			StatusCode: http.StatusBadRequest,
			Err:        ErrBadRequestHTTPStatus,
			From:       from,
			RawQuery:   "from=1755692087",
		},
		"forbidden": {
			ErrCode:    "FORBIDDEN",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusForbidden,
			Err:        ErrForbiddenHTTPStatus,
			From:       from,
			Code:       util.Pointer("test-code"),
			RawQuery:   "code=test-code&from=1755692087",
		},
		"not found": {
			ErrCode:    "NOT_FOUND",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusNotFound,
			Err:        ErrNotFoundHTTPStatus,
			From:       from,
			To:         &to,
			RawQuery:   "from=1755692087&to=1755692314",
		},
		"too many requests": {
			ErrCode:    "TOO_MANY_REQUESTS",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusTooManyRequests,
			Err:        ErrTooManyRequestsHTTPStatus,
			From:       from,
			To:         &to,
			Code:       util.Pointer("test-code"),
			RawQuery:   "code=test-code&from=1755692087&to=1755692314",
		},
		"internal server": {
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: "",
			StatusCode: http.StatusInternalServerError,
			Err:        ErrInternalHTTPStatus,
			From:       from,
			RawQuery:   "from=1755692087",
		},
		"method not allowed": {
			ErrCode:    "METHOD_NOT_ALLOWED",
			ErrMessage: "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
			Err:        ErrMethodNotAllowedStatus,
			From:       from,
			RawQuery:   "from=1755692087",
		},
		"proxy auth required": {
			ErrCode:    "",
			ErrMessage: "",
			StatusCode: http.StatusProxyAuthRequired,
			Err:        ErrUnexpectedHTTPStatus,
			From:       from,
			RawQuery:   "from=1755692087",
		},
	}
	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/api/merchant/statement", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, val.RawQuery, req.URL.RawQuery)

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			res, err := client.GetStatement(context.Background(), GetStatementRequest{
				From: val.From,
				To:   val.To,
				Code: val.Code,
			})

			assert.Error(t, err)
			assert.ErrorIs(t, err, val.Err)
			assert.Nil(t, res)

			var reqErr *RequestError

			assert.True(t, errors.As(err, &reqErr), "*RequestError")

			assert.Equal(t, val.ErrCode, reqErr.Code)
			assert.Equal(t, val.ErrMessage, reqErr.Message)
		})
	}
}

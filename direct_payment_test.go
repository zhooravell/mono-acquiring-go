package monoacquiring

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDirectPayment(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/invoice/payment-direct", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "application/json", req.Header.Get("Accept"))
		assert.Equal(t, "x-api-token-test", req.Header.Get("X-Token"))
		assert.Equal(t, "cms-test", req.Header.Get("X-Cms"))
		assert.Equal(t, "0.0.1", req.Header.Get("X-Cms-Version"))

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "invoiceId": "2210012MPLYwJjVUzchj",
  "tdsUrl": "https://example.com/tds/url",
  "status": "success",
  "failureReason": "Неправильний CVV код",
  "amount": 4200,
  "ccy": 980,
  "createdDate": null,
  "modifiedDate": null
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

	req := DirectPaymentRequest{
		Card: DirectPaymentCard{
			PAN:        "4111111111111111",
			Expiration: "1228",
			CVV:        "123",
		},
		Amount:         1000,
		Currency:       util.Pointer(980),
		InitiationKind: util.Pointer(InitiationKindMerchant),
	}

	res, err := client.DirectPayment(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "2210012MPLYwJjVUzchj", res.InvoiceID)
	assert.True(t, res.Status.IsSuccess())
}

func TestDirectPayment_Validation(t *testing.T) {
	client, err := NewClient(Config{APIKey: "test", BaseURL: DefaultBaseURL}, nil, nil)

	assert.NoError(t, err)

	req := DirectPaymentRequest{}

	res, err := client.DirectPayment(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, res)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"DirectPaymentRequest.Card.PAN":        "required",
		"DirectPaymentRequest.Card.Expiration": "required",
		"DirectPaymentRequest.Card.CVV":        "required",
		"DirectPaymentRequest.Amount":          "required",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

func TestDirectPayment_NetworkError(t *testing.T) {
	tests := map[string]struct {
		Err        error
		ErrCode    string
		ErrMessage string
		StatusCode int
	}{
		"bad request": {
			ErrCode:    "BAD_REQUEST",
			ErrMessage: "empty 'invoiceId'",
			StatusCode: http.StatusBadRequest,
			Err:        ErrBadRequestHTTPStatus,
		},
		"forbidden": {
			ErrCode:    "FORBIDDEN",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusForbidden,
			Err:        ErrForbiddenHTTPStatus,
		},
		"not found": {
			ErrCode:    "NOT_FOUND",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusNotFound,
			Err:        ErrNotFoundHTTPStatus,
		},
		"too many requests": {
			ErrCode:    "TOO_MANY_REQUESTS",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusTooManyRequests,
			Err:        ErrTooManyRequestsHTTPStatus,
		},
		"internal server": {
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: "",
			StatusCode: http.StatusInternalServerError,
			Err:        ErrInternalHTTPStatus,
		},
		"method not allowed": {
			ErrCode:    "METHOD_NOT_ALLOWED",
			ErrMessage: "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
			Err:        ErrMethodNotAllowedStatus,
		},
		"proxy auth required": {
			ErrCode:    "",
			ErrMessage: "",
			StatusCode: http.StatusProxyAuthRequired,
			Err:        ErrUnexpectedHTTPStatus,
		},
	}
	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/api/merchant/invoice/payment-direct", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodPost, req.Method)

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			ctx := context.Background()

			client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			req := DirectPaymentRequest{
				Card: DirectPaymentCard{
					PAN:        "4111111111111111",
					Expiration: "1228",
					CVV:        "123",
				},
				Amount:         1000,
				Currency:       util.Pointer(980),
				InitiationKind: util.Pointer(InitiationKindMerchant),
			}
			res, err := client.DirectPayment(ctx, req)

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

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

func TestGetReceipt(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/invoice/receipt", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "application/json", req.Header.Get("Accept"))
		assert.Equal(t, "x-api-token-test", req.Header.Get("X-Token"))
		assert.Equal(t, "cms-test", req.Header.Get("X-Cms"))
		assert.Equal(t, "0.0.1", req.Header.Get("X-Cms-Version"))
		assert.Equal(t, "email=test%40monobank.ua&invoiceId=p2_9ZgpZVsl3", req.URL.RawQuery)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "file": "CJFVBERi0xLj4QKJaq..."
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

	res, err := client.GetReceipt(context.Background(), GetReceiptRequest{
		InvoiceID: "p2_9ZgpZVsl3",
		Email:     util.Pointer("test@monobank.ua"),
	})

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "CJFVBERi0xLj4QKJaq...", res.File)
}

func TestGetReceipt_Validation(t *testing.T) {
	client, err := NewClient(Config{APIKey: "test", BaseURL: DefaultBaseURL}, nil, nil)

	assert.NoError(t, err)

	res, err := client.GetReceipt(context.Background(), GetReceiptRequest{
		Email: util.Pointer("1234"),
	})

	assert.Error(t, err)
	assert.Nil(t, res)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"GetReceiptRequest.InvoiceID": "required",
		"GetReceiptRequest.Email":     "email",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

func TestGetReceipt_NetworkError(t *testing.T) {
	tests := map[string]struct {
		Err        error
		ErrCode    string
		ErrMessage string
		InvoiceID  string
		StatusCode int
	}{
		"bad request": {
			ErrCode:    "BAD_REQUEST",
			ErrMessage: "empty 'invoiceId'",
			StatusCode: http.StatusBadRequest,
			Err:        ErrBadRequestHTTPStatus,
			InvoiceID:  "test-1",
		},
		"forbidden": {
			ErrCode:    "FORBIDDEN",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusForbidden,
			Err:        ErrForbiddenHTTPStatus,
			InvoiceID:  "test-2",
		},
		"not found": {
			ErrCode:    "NOT_FOUND",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusNotFound,
			Err:        ErrNotFoundHTTPStatus,
			InvoiceID:  "test-3",
		},
		"too many requests": {
			ErrCode:    "TOO_MANY_REQUESTS",
			ErrMessage: "invalid 'qrId'",
			StatusCode: http.StatusTooManyRequests,
			Err:        ErrTooManyRequestsHTTPStatus,
			InvoiceID:  "test-4",
		},
		"internal server": {
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: "",
			StatusCode: http.StatusInternalServerError,
			Err:        ErrInternalHTTPStatus,
			InvoiceID:  "test-5",
		},
		"method not allowed": {
			ErrCode:    "METHOD_NOT_ALLOWED",
			ErrMessage: "Method not allowed",
			StatusCode: http.StatusMethodNotAllowed,
			Err:        ErrMethodNotAllowedStatus,
			InvoiceID:  "test-6",
		},
		"proxy auth required": {
			ErrCode:    "",
			ErrMessage: "",
			StatusCode: http.StatusProxyAuthRequired,
			Err:        ErrUnexpectedHTTPStatus,
			InvoiceID:  "test-7",
		},
	}
	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/api/merchant/invoice/receipt", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "invoiceId="+val.InvoiceID, req.URL.RawQuery)

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			res, err := client.GetReceipt(context.Background(), GetReceiptRequest{
				InvoiceID: val.InvoiceID,
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

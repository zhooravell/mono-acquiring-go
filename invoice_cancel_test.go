package monoacquiring

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCancelInvoice(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/invoice/cancel", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodPost, req.Method)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "status": "processing",
  "createdDate": "2025-07-17T12:00:00+03:00",
  "modifiedDate": "2025-07-17T14:00:00+03:00"
}`)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := CancelInvoiceRequest{
		InvoiceID: "p2_9ZgpZVsl3",
	}
	res, err := client.CancelInvoice(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "processing", res.Status)
	assert.Equal(t, "2025-07-17T12:00:00+03:00", res.CreatedDate)
	assert.Equal(t, "2025-07-17T14:00:00+03:00", res.ModifiedDate)
}

func TestCancelInvoice_Validation(t *testing.T) {
	client, err := NewClient(Config{APIKey: "test", BaseURL: DefaultBaseURL}, nil, nil)

	assert.NoError(t, err)

	res, err := client.CancelInvoice(context.Background(), CancelInvoiceRequest{})

	assert.Error(t, err)
	assert.Nil(t, res)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"CancelInvoiceRequest.InvoiceID": "required",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

func TestCancelInvoice_NetworkError(t *testing.T) {
	tests := map[string]struct {
		ErrCode    string
		ErrMessage string
		StatusCode int
		Err        error
		InvoiceID  string
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
			mux.HandleFunc("/api/merchant/invoice/cancel", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodPost, req.Method)

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			ctx := context.Background()

			client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			res, err := client.CancelInvoice(ctx, CancelInvoiceRequest{InvoiceID: val.InvoiceID})

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

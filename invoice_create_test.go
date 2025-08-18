package monoacquiring

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceCreate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/invoice/create", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "invoiceId": "p2_9ZgpZVsl3",
  "pageUrl": "https://pay.mbnk.biz/p2_9ZgpZVsl3"
}`)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := InvoiceCreateRequest{
		Amount: 100,
	}
	res, err := client.CreateInvoice(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "p2_9ZgpZVsl3", res.InvoiceID)
	assert.Equal(t, "https://pay.mbnk.biz/p2_9ZgpZVsl3", res.PageURL)
}

func TestInvoiceCreate_Validation(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(Config{APIKey: "test", BaseURL: DefaultBaseURL}, nil, nil)

	assert.NoError(t, err)

	req := InvoiceCreateRequest{
		WebHookURL:  util.Pointer("test"),
		RedirectURL: util.Pointer("test"),
		PaymentType: "test",
	}

	_, err = client.CreateInvoice(ctx, req)

	assert.Error(t, err)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"InvoiceCreateRequest.Amount":      "required",
		"InvoiceCreateRequest.RedirectURL": "http_url",
		"InvoiceCreateRequest.WebHookURL":  "http_url",
		"InvoiceCreateRequest.PaymentType": "oneof",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

func TestInvoiceCreate_NetworkError(t *testing.T) {
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
			mux.HandleFunc("/api/merchant/invoice/create", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodPost, req.Method)

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			ctx := context.Background()

			client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			res, err := client.CreateInvoice(ctx, InvoiceCreateRequest{Amount: 100})

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

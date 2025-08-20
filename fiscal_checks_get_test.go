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

func TestGetFiscalChecks(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/invoice/fiscal-checks", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "application/json", req.Header.Get("Accept"))
		assert.Equal(t, "x-api-token-test", req.Header.Get("X-Token"))
		assert.Equal(t, "cms-test", req.Header.Get("X-Cms"))
		assert.Equal(t, "0.0.1", req.Header.Get("X-Cms-Version"))
		assert.Equal(t, "invoiceId=p2_9ZgpZVsl3", req.URL.RawQuery)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "checks": [
    {
      "id": "a2fd4aef-cdb8-4e25-9b36-b6d4672c554d",
      "type": "sale",
      "status": "done",
      "statusDescription": null,
      "taxUrl": null,
      "file": null,
      "fiscalizationSource": "monopay"
    },
	{
      "id": "12d10651-8105-4e2c-811f-ae6e32a2a588",
      "type": "return",
      "status": "failed",
      "fiscalizationSource": "checkbox"
    },
	{
      "id": "a2a13e1f-7373-4642-838c-be3bc8b67819",
      "type": "return",
      "status": "process",
      "statusDescription": "test description",
      "taxUrl": "https://cabinet.tax.gov.ua/cashregs/check",
      "file": "CJFVBERi0xL....",
      "fiscalizationSource": "vchasnokasa"
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

	res, err := client.GetFiscalChecks(context.Background(), GetFiscalChecksRequest{
		InvoiceID: "p2_9ZgpZVsl3",
	})

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Checks, 3)

	assert.Equal(t, "a2fd4aef-cdb8-4e25-9b36-b6d4672c554d", res.Checks[0].ID)
	assert.True(t, res.Checks[0].Type.IsSale())
	assert.True(t, res.Checks[0].Status.IsDone())
	assert.True(t, res.Checks[0].FiscalizationSource.IsMonoPay())
	assert.Nil(t, res.Checks[0].StatusDescription)
	assert.Nil(t, res.Checks[0].TaxURL)
	assert.Nil(t, res.Checks[0].File)

	assert.Equal(t, "12d10651-8105-4e2c-811f-ae6e32a2a588", res.Checks[1].ID)
	assert.True(t, res.Checks[1].Type.IsReturn())
	assert.True(t, res.Checks[1].Status.IsFailed())
	assert.True(t, res.Checks[1].FiscalizationSource.IsCheckBox())
	assert.Nil(t, res.Checks[1].StatusDescription)
	assert.Nil(t, res.Checks[1].TaxURL)
	assert.Nil(t, res.Checks[1].File)

	assert.Equal(t, "a2a13e1f-7373-4642-838c-be3bc8b67819", res.Checks[2].ID)
	assert.True(t, res.Checks[2].Type.IsReturn())
	assert.True(t, res.Checks[2].Status.IsProcess())
	assert.True(t, res.Checks[2].FiscalizationSource.IsVchasnoKasa())
	assert.Equal(t, "test description", *res.Checks[2].StatusDescription)
	assert.Equal(t, "https://cabinet.tax.gov.ua/cashregs/check", *res.Checks[2].TaxURL)
	assert.Equal(t, "CJFVBERi0xL....", *res.Checks[2].File)
}

func TestGetFiscalChecks_Validation(t *testing.T) {
	client, err := NewClient(Config{APIKey: "test", BaseURL: DefaultBaseURL}, nil, nil)

	assert.NoError(t, err)

	res, err := client.GetFiscalChecks(context.Background(), GetFiscalChecksRequest{})

	assert.Error(t, err)
	assert.Nil(t, res)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"GetFiscalChecksRequest.InvoiceID": "required",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

func TestGetFiscalChecks_NetworkError(t *testing.T) {
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
			mux.HandleFunc("/api/merchant/invoice/fiscal-checks", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "invoiceId="+val.InvoiceID, req.URL.RawQuery)

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			res, err := client.GetFiscalChecks(context.Background(), GetFiscalChecksRequest{
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

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

func TestSyncPayment(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/invoice/sync-payment", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "application/json", req.Header.Get("Accept"))
		assert.Equal(t, "x-api-token-test", req.Header.Get("X-Token"))
		assert.Equal(t, "cms-test", req.Header.Get("X-Cms"))
		assert.Equal(t, "0.0.1", req.Header.Get("X-Cms-Version"))

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "invoiceId": "p2_9ZgpZVsl3",
  "status": "success",
  "failureReason": "Неправильний CVV код",
  "errCode": "59",
  "amount": 4200,
  "ccy": 980,
  "finalAmount": 4200,
  "createdDate": null,
  "modifiedDate": null,
  "reference": "84d0070ee4e44667b31371d8f8813947",
  "destination": "Покупка щастя",
  "cancelList": [
    {
      "status": null,
      "amount": 4200,
      "ccy": 980,
      "createdDate": null,
      "modifiedDate": null,
      "approvalCode": "662476",
      "rrn": "060189181768",
      "extRef": "635ace02599849e981b2cd7a65f417fe"
    }
  ],
  "paymentInfo": {
    "maskedPan": "444403******1902",
    "approvalCode": "662476",
    "rrn": "060189181768",
    "tranId": "13194036",
    "terminal": "MI001088",
    "bank": "Універсал Банк",
    "paymentSystem": "visa",
    "paymentMethod": null,
    "fee": null,
    "country": "804",
    "agentFee": null
  },
  "walletData": {
    "cardToken": "67XZtXdR4NpKU3",
    "walletId": "c1376a611e17b059aeaf96b73258da9c",
    "status": null
  },
  "tipsInfo": {
    "employeeId": null,
    "amount": 4200
  }
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

	res, err := client.SyncPayment(context.Background(), SyncPaymentRequest{Amount: 1000, Currency: 980})

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "p2_9ZgpZVsl3", res.InvoiceID)
	assert.True(t, res.Status.IsSuccess())
}

func TestSyncPayment_Validation(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(Config{APIKey: "test", BaseURL: DefaultBaseURL}, nil, nil)

	assert.NoError(t, err)

	req := SyncPaymentRequest{
		GooglePay: &GooglePay{
			Expiration: "test",
		},
		ApplePay:            &ApplePay{},
		SyncPaymentCard:     &SyncPaymentCard{},
		MerchantPaymentInfo: &MerchantPaymentInfo{},
	}

	res, err := client.SyncPayment(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"SyncPaymentRequest.Currency":                     "required",
		"SyncPaymentRequest.Amount":                       "required",
		"SyncPaymentRequest.GooglePay.Token":              "required",
		"SyncPaymentRequest.GooglePay.Expiration":         "card_exp",
		"SyncPaymentRequest.GooglePay.EciIndicator":       "required",
		"SyncPaymentRequest.ApplePay.Token":               "required",
		"SyncPaymentRequest.ApplePay.Expiration":          "required",
		"SyncPaymentRequest.ApplePay.EciIndicator":        "required",
		"SyncPaymentRequest.SyncPaymentCard.PAN":          "required",
		"SyncPaymentRequest.SyncPaymentCard.Type":         "required",
		"SyncPaymentRequest.SyncPaymentCard.Expiration":   "required",
		"SyncPaymentRequest.SyncPaymentCard.EciIndicator": "required",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

func TestSyncPayment_NetworkError(t *testing.T) {
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
			mux.HandleFunc("/api/merchant/invoice/sync-payment", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodPost, req.Method)

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			ctx := context.Background()

			client, err := NewClient(Config{APIKey: "test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			req := SyncPaymentRequest{Amount: 1000, Currency: 980}
			res, err := client.SyncPayment(ctx, req)

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

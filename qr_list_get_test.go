package monoacquiring

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetQRList(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/merchant/qr/list", func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "application/json", req.Header.Get("Accept"))
		assert.Equal(t, "x-api-token-test", req.Header.Get("X-Token"))
		assert.Equal(t, "cms-test", req.Header.Get("X-Cms"))
		assert.Equal(t, "0.0.1", req.Header.Get("X-Cms-Version"))

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{
  "list": [
    {
      "shortQrId": "1",
      "qrId": "id-1",
      "amountType": "merchant",
      "pageUrl": "https://pay.mbnk.biz/1"
    },
	{
      "shortQrId": "2",
      "qrId": "id-2",
      "amountType": "client",
      "pageUrl": "https://pay.mbnk.biz/2"
    },
	{
      "shortQrId": "3",
      "qrId": "id-3",
      "amountType": "fix",
      "pageUrl": "https://pay.mbnk.biz/3"
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

	res, err := client.GetQRList(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.List, 3)

	assert.Equal(t, "1", res.List[0].ShortQrID)
	assert.Equal(t, "id-1", res.List[0].QrID)
	assert.Equal(t, "https://pay.mbnk.biz/1", res.List[0].PageURL)
	assert.True(t, res.List[0].AmountType.IsMerchant())
	assert.Equal(t, "merchant", res.List[0].AmountType.String())

	assert.Equal(t, "2", res.List[1].ShortQrID)
	assert.Equal(t, "id-2", res.List[1].QrID)
	assert.Equal(t, "https://pay.mbnk.biz/2", res.List[1].PageURL)
	assert.True(t, res.List[1].AmountType.IsClient())
	assert.Equal(t, "client", res.List[1].AmountType.String())

	assert.Equal(t, "3", res.List[2].ShortQrID)
	assert.Equal(t, "id-3", res.List[2].QrID)
	assert.Equal(t, "https://pay.mbnk.biz/3", res.List[2].PageURL)
	assert.True(t, res.List[2].AmountType.IsFix())
	assert.Equal(t, "fix", res.List[2].AmountType.String())
}

func TestGetQRList_NetworkError(t *testing.T) {
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
			mux.HandleFunc("/api/merchant/qr/list", func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "application/json", req.Header.Get("Accept"))
				assert.Equal(t, "x-api-token-test", req.Header.Get("X-Token"))
				assert.Equal(t, "golang", req.Header.Get("X-Cms"))
				assert.Equal(t, runtime.GOOS+" "+runtime.GOARCH+" "+runtime.Version(), req.Header.Get("X-Cms-Version"))

				w.WriteHeader(val.StatusCode)
				_, _ = fmt.Fprint(w, `{"errCode": "`+val.ErrCode+`","errText": "`+val.ErrMessage+`"}`)
			})

			srv := httptest.NewServer(mux)
			defer srv.Close()

			client, err := NewClient(Config{APIKey: "x-api-token-test", BaseURL: srv.URL}, srv.Client(), nil)

			assert.NoError(t, err)

			res, err := client.GetQRList(context.Background())

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

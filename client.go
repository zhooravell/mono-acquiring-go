package monoacquiring

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"runtime"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type (
	Config struct {
		APIKey     string `validate:"required"`
		BaseURL    string `validate:"required,url"`
		CMS        string
		CMSVersion string
	}

	Client struct {
		httpClient *http.Client
		validator  *validator.Validate
		cnf        Config
	}
)

const (
	DefaultBaseURL = "https://api.monobank.ua/"

	invoiceStatusPath        = "/api/merchant/invoice/status"
	invoiceCancelPath        = "/api/merchant/invoice/cancel"
	invoiceCreatePath        = "/api/merchant/invoice/create"
	invoiceRemovePath        = "/api/merchant/invoice/remove"
	getPublicKeyPath         = "/api/merchant/pubkey"
	getEmployeeListPath      = "/api/merchant/employee/list"
	getMerchantDetailsPath   = "/api/merchant/details"
	getQRListPath            = "/api/merchant/qr/list"
	getWalletCardListPath    = "/api/merchant/wallet"
	getSplitReceiverListPath = "/api/merchant/split-receiver/list"
	removeWalletCardPath     = "/api/merchant/wallet/card"
	getQrDetailsPath         = "/api/merchant/qr/details"
	qrResetAmountPath        = "/api/merchant/qr/reset-amount"
	getSubMerchantListPath   = "/api/merchant/submerchant/list"
	getReceiptPath           = "/api/merchant/invoice/receipt"
	geFiscalChecksPath       = "/api/merchant/invoice/fiscal-checks"
	getStatementPath         = "/api/merchant/statement"
	finalizeHoldPath         = "/api/merchant/invoice/finalize"
)

var statusToError = map[int]error{
	http.StatusBadRequest:          ErrBadRequestHTTPStatus,
	http.StatusForbidden:           ErrForbiddenHTTPStatus,
	http.StatusNotFound:            ErrNotFoundHTTPStatus,
	http.StatusTooManyRequests:     ErrTooManyRequestsHTTPStatus,
	http.StatusInternalServerError: ErrInternalHTTPStatus,
	http.StatusMethodNotAllowed:    ErrMethodNotAllowedStatus,
}

func NewClient(config Config, httpClient *http.Client, validate *validator.Validate) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if validate == nil {
		validate = validator.New()
	}

	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return &Client{cnf: config, httpClient: httpClient, validator: validate}, nil
}

func (c *Client) addHeaders(req *http.Request) *http.Request {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Token", c.cnf.APIKey)
	req.Header.Add("X-Cms", util.Ternary(c.cnf.CMS == "", "golang", c.cnf.CMS))
	req.Header.Add(
		"X-Cms-Version",
		util.Ternary(c.cnf.CMSVersion == "", runtime.GOOS+" "+runtime.GOARCH+" "+runtime.Version(), c.cnf.CMSVersion),
	)

	return req
}

func (c *Client) baseURL() (*url.URL, error) {
	baseURL, err := url.Parse(c.cnf.BaseURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return baseURL, nil
}

func (c *Client) newRequest(
	ctx context.Context,
	method, path string,
	query map[string]string,
	body io.Reader,
) (*http.Request, error) {
	var (
		err     error
		req     *http.Request
		baseURL *url.URL
	)

	if baseURL, err = c.baseURL(); err != nil {
		return nil, err
	}

	baseURL.Path = path

	if len(query) > 0 {
		q := baseURL.Query()

		for k, v := range query {
			q.Add(k, v)
		}

		baseURL.RawQuery = q.Encode()
	}

	if req, err = http.NewRequestWithContext(ctx, method, baseURL.String(), body); err != nil {
		return nil, errors.WithStack(err)
	}

	return c.addHeaders(req), nil
}

func (c *Client) doReq(req *http.Request, result any) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		var errorData errorData

		err = json.NewDecoder(res.Body).Decode(&errorData)

		if err != nil {
			errorData.Message = err.Error()
		}

		if errCode, ok := statusToError[res.StatusCode]; ok {
			return newRequestError(errCode, errorData.Code, errorData.Message)
		}

		return newRequestError(ErrUnexpectedHTTPStatus, errorData.Code, errorData.Message)
	}

	if result != nil {
		if err := json.NewDecoder(res.Body).Decode(result); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

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
		cnf        Config
		httpClient *http.Client
		validator  *validator.Validate
	}
)

const (
	DefaultBaseURL = "https://api.monobank.ua/"
)

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
	var (
		err     error
		res     *http.Response
		resBody []byte
	)

	if res, err = c.httpClient.Do(req); err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if resBody, err = io.ReadAll(res.Body); err != nil {
		return errors.WithStack(err)
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		var errorData errorData

		_ = json.Unmarshal(resBody, &errorData)

		switch res.StatusCode {
		case http.StatusBadRequest:
			return newRequestError(ErrBadRequestHTTPStatus, errorData.Code, errorData.Message)
		case http.StatusForbidden:
			return newRequestError(ErrForbiddenHTTPStatus, errorData.Code, errorData.Message)
		case http.StatusNotFound:
			return newRequestError(ErrNotFoundHTTPStatus, errorData.Code, errorData.Message)
		case http.StatusTooManyRequests:
			return newRequestError(ErrTooManyRequestsHTTPStatus, errorData.Code, errorData.Message)
		case http.StatusInternalServerError:
			return newRequestError(ErrInternalHTTPStatus, errorData.Code, errorData.Message)
		case http.StatusMethodNotAllowed:
			return newRequestError(ErrMethodNotAllowedStatus, errorData.Code, errorData.Message)
		default:
			return newRequestError(ErrUnexpectedHTTPStatus, errorData.Code, errorData.Message)
		}
	}

	if err = json.Unmarshal(resBody, result); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

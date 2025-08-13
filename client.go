package monoacquiring

import (
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/pkg/errors"
)

type (
	Config struct {
		APIKey     string
		BaseURL    string
		CMS        string
		CMSVersion string
	}

	Client struct {
		cnf        Config
		httpClient *http.Client
	}
)

func (cnf Config) Validate() error {
	apiKey := strings.TrimSpace(cnf.APIKey)

	if apiKey == "" {
		return ErrBlankAPIKey
	}

	baseURL := strings.TrimSpace(cnf.BaseURL)

	if baseURL == "" {
		return ErrBlankBaseURL
	}

	if _, err := url.Parse(baseURL); err != nil {
		return errors.Wrap(ErrInvalidBaseURL, err.Error())
	}

	return nil
}

const (
	DefaultBaseURL = "https://api.monobank.ua/"
)

func NewClient(config Config, httpClient *http.Client) (*Client, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &Client{cnf: config, httpClient: httpClient}, nil
}

func (c *Client) addHeaders(req *http.Request) {
	req.Header.Add("X-Token", c.cnf.APIKey)
	req.Header.Add("X-Cms", util.Ternary(c.cnf.CMS == "", "golang", c.cnf.CMS))
	req.Header.Add(
		"X-Cms-Version",
		util.Ternary(c.cnf.CMSVersion == "", runtime.GOOS+" "+runtime.GOARCH+" "+runtime.Version(), c.cnf.CMSVersion),
	)
}

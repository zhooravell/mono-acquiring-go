package monoacquiring

import (
	"context"
	"net/http"
)

type GetMerchantDetailsResponse struct {
	MerchantID   string `json:"merchantId"`
	MerchantName string `json:"merchantName"`
	Edrpou       string `json:"edrpou"`
}

func (c *Client) GetMerchantDetails(ctx context.Context) (*GetMerchantDetailsResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, getMerchantDetailsPath, nil, nil)
	if err != nil {
		return nil, err
	}

	var result GetMerchantDetailsResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

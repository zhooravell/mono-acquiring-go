package monoacquiring

import (
	"context"
	"net/http"
)

type GetSubMerchantList struct {
	List []struct {
		Code   string `json:"code"`
		Edrpou string `json:"edrpou"`
		Iban   string `json:"iban"`
		Owner  string `json:"owner"`
	} `json:"list"`
}

func (c *Client) GetSubMerchantList(ctx context.Context) (*GetSubMerchantList, error) {
	req, err := c.newRequest(ctx, http.MethodGet, getSubMerchantListPath, nil, nil)
	if err != nil {
		return nil, err
	}

	var result GetSubMerchantList

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

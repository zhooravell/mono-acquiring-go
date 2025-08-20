//
// https://monobank.ua/api-docs/acquiring/instrumenty-rozrobky/webhooks/get--api--merchant--pubkey
//

package monoacquiring

import (
	"context"
	"net/http"
)

type GetPublicKeyResponse struct {
	Key string `json:"key"`
}

func (c *Client) GetPublicKey(ctx context.Context) (*GetPublicKeyResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, getPublicKeyPath, nil, nil)
	if err != nil {
		return nil, err
	}

	var result GetPublicKeyResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

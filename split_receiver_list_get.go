//
// https://monobank.ua/api-docs/acquiring/metody/rozshcheplennia/get--api--merchant--split-receiver--list
//

package monoacquiring

import (
	"context"
	"net/http"
)

type GetSplitReceiverListResponse struct {
	List []struct {
		SplitReceiverID string `json:"splitReceiverId"`
		Name            string `json:"name"`
	} `json:"list"`
}

func (c *Client) GetSplitReceiverList(ctx context.Context) (*GetSplitReceiverListResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, getSplitReceiverListPath, nil, nil)
	if err != nil {
		return nil, err
	}

	var result GetSplitReceiverListResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

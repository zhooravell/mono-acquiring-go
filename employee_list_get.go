package monoacquiring

import (
	"context"
	"net/http"
)

type GetEmployeeListResponse struct {
	List []struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
		ExternalReference string `json:"extRef"`
	} `json:"list"`
}

func (c *Client) GetEmployeeList(ctx context.Context) (*GetEmployeeListResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, getEmployeeListPath, nil, nil)
	if err != nil {
		return nil, err
	}

	var result GetEmployeeListResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

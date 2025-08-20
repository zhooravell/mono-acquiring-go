package monoacquiring

import (
	"context"
	"net/http"
)

type GetQRListResponse struct {
	List []struct {
		ShortQrID  string       `json:"shortQrId"`
		QrID       string       `json:"qrId"`
		AmountType QRAmountType `json:"amountType"`
		PageURL    string       `json:"pageUrl"`
	} `json:"list"`
}

func (c *Client) GetQRList(ctx context.Context) (*GetQRListResponse, error) {
	req, err := c.newRequest(ctx, http.MethodGet, getQRListPath, nil, nil)
	if err != nil {
		return nil, err
	}

	var result GetQRListResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

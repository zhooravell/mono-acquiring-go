package monoacquiring

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type GetWalletCardListRequest struct {
	WalletID string `validate:"required"`
}

type GetWalletCardListResponse struct {
	Wallet []struct {
		CardToken string `json:"cardToken"`
		MaskedPan string `json:"maskedPan"`
		Country   string `json:"country"`
	} `json:"wallet"`
}

func (c *Client) GetWalletCardList(
	ctx context.Context,
	payload GetWalletCardListRequest,
) (*GetWalletCardListResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := make(map[string]string, 1)
	query["walletId"] = payload.WalletID

	req, err := c.newRequest(ctx, http.MethodGet, getWalletCardListPath, query, nil)
	if err != nil {
		return nil, err
	}

	var result GetWalletCardListResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

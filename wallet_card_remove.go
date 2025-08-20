//
// https://monobank.ua/api-docs/acquiring/dodatkova-funktsionalnist/tokenizatsiia/delete--api--merchant--wallet--card
//

package monoacquiring

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type RemoveWalletCardRequest struct {
	CardToken string `validate:"required"`
}

func (c *Client) RemoveWalletCard(ctx context.Context, payload RemoveWalletCardRequest) error {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return errors.WithStack(err)
	}

	query := make(map[string]string, 1)
	query["cardToken"] = payload.CardToken

	req, err := c.newRequest(ctx, http.MethodDelete, removeWalletCardPath, query, nil)
	if err != nil {
		return err
	}

	return c.doReq(req, nil)
}

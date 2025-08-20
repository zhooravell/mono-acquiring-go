//
// https://monobank.ua/api-docs/acquiring/metody/internet-ekvairynh/get--api--merchant--invoice--receipt
// https://monobank.ua/api-docs/acquiring/metody/oplata-v-zastosunku/get--api--merchant--invoice--receipt
// https://monobank.ua/api-docs/acquiring/metody/rozshcheplennia/get--api--merchant--invoice--receipt
//

package monoacquiring

import (
	"context"
	"net/http"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/pkg/errors"
)

type GetReceiptRequest struct {
	Email     *string `validate:"omitempty,email"`
	InvoiceID string  `validate:"required"`
}

type GetReceiptResponse struct {
	File string `json:"file"`
}

func (c *Client) GetReceipt(ctx context.Context, payload GetReceiptRequest) (*GetReceiptResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := make(map[string]string, 2)
	query["invoiceId"] = payload.InvoiceID

	if payload.Email != nil {
		query["email"] = util.PointerValue(payload.Email)
	}

	req, err := c.newRequest(ctx, http.MethodGet, getReceiptPath, query, nil)
	if err != nil {
		return nil, err
	}

	var result GetReceiptResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

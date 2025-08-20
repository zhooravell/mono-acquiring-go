//
// https://monobank.ua/api-docs/acquiring/dodatkova-funktsionalnist/tokenizatsiia/post--api--merchant--invoice--remove
// https://monobank.ua/api-docs/acquiring/metody/rozshcheplennia/post--api--merchant--invoice--remove
//

package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type RemoveInvoiceRequest struct {
	InvoiceID string `json:"invoiceId" validate:"required"`
}

func (c *Client) RemoveInvoice(ctx context.Context, payload RemoveInvoiceRequest) error {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(payload); err != nil {
		return errors.Wrap(err, "failed to marshal remove invoice request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, invoiceRemovePath, nil, buf)
	if err != nil {
		return err
	}

	return c.doReq(req, nil)
}

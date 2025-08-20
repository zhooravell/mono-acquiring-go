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

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal remove invoice request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, invoiceRemovePath, nil, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	return c.doReq(req, nil)
}

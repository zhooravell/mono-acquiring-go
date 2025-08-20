package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type QrResetAmountRequest struct {
	QrID string `json:"qrId" validate:"required"`
}

func (c *Client) QrResetAmount(ctx context.Context, payload QrResetAmountRequest) error {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(payload); err != nil {
		return errors.Wrap(err, "failed to marshal qr reset amount request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, qrResetAmountPath, nil, buf)
	if err != nil {
		return err
	}

	return c.doReq(req, nil)
}

//
// https://monobank.ua/api-docs/acquiring/metody/internet-ekvairynh/get--api--merchant--invoice--status
//

package monoacquiring

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type GetInvoiceStatusRequest struct {
	InvoiceID string `validate:"required"`
}

type GetInvoiceStatusResponse struct {
	InvoiceID     string           `json:"invoiceId"`
	Status        string           `json:"status"`
	FailureReason *string          `json:"failureReason,omitempty"`
	ErrCode       *string          `json:"errCode,omitempty"`
	Amount        int64            `json:"amount"`
	Currency      int              `json:"ccy"`
	FinalAmount   *int             `json:"finalAmount,omitempty"`
	CreatedDate   *string          `json:"createdDate,omitempty"`
	ModifiedDate  *string          `json:"modifiedDate,omitempty"`
	Reference     *string          `json:"reference,omitempty"`
	Destination   *string          `json:"destination,omitempty"`
	CancelList    []CancelListItem `json:"cancelList,omitempty"`
	PaymentInfo   *PaymentInfo     `json:"paymentInfo"`
	WalletData    *WalletData      `json:"walletData,omitempty"`
	TipsInfo      *TipsInfo        `json:"tipsInfo,omitempty"`
}

func (c *Client) GetInvoiceStatus(
	ctx context.Context,
	payload GetInvoiceStatusRequest,
) (*GetInvoiceStatusResponse, error) {
	var (
		err    error
		req    *http.Request
		result GetInvoiceStatusResponse

		path  = "/api/merchant/invoice/status"
		query = make(map[string]string, 1)
	)

	if err = c.validator.StructCtx(ctx, payload); err != nil {
		return nil, errors.WithStack(err)
	}

	query["invoiceId"] = payload.InvoiceID

	if req, err = c.newRequest(ctx, http.MethodGet, path, query, nil); err != nil {
		return nil, err
	}

	if err = c.doReq(req, &result); err != nil {
		return nil, errors.WithStack(err)
	}

	return &result, nil
}

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
	Destination   *string          `json:"destination,omitempty"`
	TipsInfo      *TipsInfo        `json:"tipsInfo,omitempty"`
	FinalAmount   *int             `json:"finalAmount,omitempty"`
	CreatedDate   *string          `json:"createdDate,omitempty"`
	ModifiedDate  *string          `json:"modifiedDate,omitempty"`
	Reference     *string          `json:"reference,omitempty"`
	ErrCode       *string          `json:"errCode,omitempty"`
	PaymentInfo   *PaymentInfo     `json:"paymentInfo"`
	FailureReason *string          `json:"failureReason,omitempty"`
	WalletData    *WalletData      `json:"walletData,omitempty"`
	InvoiceID     string           `json:"invoiceId"`
	Status        string           `json:"status"`
	CancelList    []CancelListItem `json:"cancelList,omitempty"`
	Amount        int64            `json:"amount"`
	Currency      int              `json:"ccy"`
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

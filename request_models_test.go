package monoacquiring

import (
	"strings"
	"testing"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMerchantPaymentInfo_Validation(t *testing.T) {
	vld := validator.New()

	data := MerchantPaymentInfo{
		Destination:    util.Pointer(strings.Repeat("A", 285)),
		Comment:        util.Pointer(strings.Repeat("B ", 300)),
		CustomerEmails: []string{"test"},
		Discounts: []Discount{
			{
				Type:  "test",
				Mode:  "test",
				Value: 0.00000001,
			},
		},
		BasketOrder: []BasketOrder{
			{},
		},
	}

	err := vld.Struct(data)

	assert.Error(t, err)

	var errs validator.ValidationErrors

	assert.True(t, errors.As(err, &errs), "validator.ValidationErrors")

	expectedErrors := map[string]string{
		"MerchantPaymentInfo.Destination":         "max",
		"MerchantPaymentInfo.Comment":             "max",
		"MerchantPaymentInfo.CustomerEmails[0]":   "email",
		"MerchantPaymentInfo.Discounts[0].Type":   "oneof",
		"MerchantPaymentInfo.Discounts[0].Mode":   "oneof",
		"MerchantPaymentInfo.Discounts[0].Value":  "min",
		"MerchantPaymentInfo.BasketOrder[0].Name": "required",
		"MerchantPaymentInfo.BasketOrder[0].Qty":  "required",
		"MerchantPaymentInfo.BasketOrder[0].Sum":  "required",
		"MerchantPaymentInfo.BasketOrder[0].Code": "required",
	}

	assert.Len(t, errs, len(expectedErrors))

	for _, e := range errs {
		tag, ok := expectedErrors[e.Namespace()]
		assert.True(t, ok, "Unexpected error: %v", e)
		assert.Equal(t, tag, e.Tag(), "Wrong tag for %s", e.Namespace())
	}
}

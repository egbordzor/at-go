package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	pay "github.com/wondenge/at-go/payments"
	"go.uber.org/zap"
)

// Collect money into your payment wallet.
func (c *Client) BankCheckout(ctx context.Context, p *pay.BankCheckoutPayload) (res *pay.BankCheckoutResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/checkout/charge"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Validate a bank checkout charge requestFormBody
func (c *Client) BankCheckoutValidate(ctx context.Context, p *pay.BankCheckoutValidatePayload) (res *pay.BankCheckoutValidateResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/checkout/validate"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Initiate a bank transfer requestFormBody.
func (c *Client) BankTransfer(ctx context.Context, p *pay.BankTransferPayload) (res *pay.BankTransferResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/transfer"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

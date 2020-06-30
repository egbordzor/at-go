package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	pay "github.com/wondenge/at-go/payments"
	"go.uber.org/zap"
)

// Mobile Checkout APIs allow you to initiate Customer to Business (C2B) payments
// on a mobile subscriber’s device.
// This allows for a smoother checkout experience, since the client will no longer
// need to remember the amount or an account number to complete the transaction.
func (c *Client) MobileCheckout(ctx context.Context, p *pay.MobileCheckoutPayload) (res *pay.MobileCheckoutResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/mobile/checkout/request"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Mobile Business To Consumer (B2C) APIs allow you to send payments to mobile
// subscribers from your Payment Wallet.
func (c *Client) MobileB2C(ctx context.Context, p *pay.MobileB2CPayload) (res *pay.MobileB2CResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/mobile/b2c/request"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Mobile Business To Business (B2B) APIs allow you to send payments to businesses
// e.g banks from your Payment Wallet.
func (c *Client) MobileB2B(ctx context.Context, p *pay.MobileB2BPayload) (res *pay.MobileB2BResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/mobile/b2b/requestFormBody"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

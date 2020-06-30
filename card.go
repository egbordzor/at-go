package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	pay "github.com/wondenge/at-go/payments"
	"go.uber.org/zap"
)

// Collect money into your Payment Wallet by initiating transactions that deduct
// money from a customers Debit or Credit Card.
// These APIs are currently only available in Nigeria on MasterCard and Verve cards.
func (c *Client) CardCheckout(ctx context.Context, p *pay.CardCheckoutPayload) (res *pay.CardCheckoutResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/card/checkout/charge"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Allows your application to validate card checkout charge requests.
func (c *Client) CardCheckoutValidate(ctx context.Context, p *pay.CardCheckoutValidatePayload) (res *pay.CardCheckoutValidateResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/card/checkout/validate"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

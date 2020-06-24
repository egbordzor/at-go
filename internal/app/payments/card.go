package payments

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/pkg/gen/africastalking"
)

// Collect money into your Payment Wallet
func (c *Client) CardCheckout(ctx context.Context, p *africastalking.CardCheckoutPayload) (res *africastalking.CardCheckoutResponse, err error) {
	res = &africastalking.CardCheckoutResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.CardCheckout"))
	return
}

// Validate card checkout charge requests
func (c *Client) CardCheckoutValidate(ctx context.Context, p *africastalking.CardCheckoutValidatePayload) (res *africastalking.CardCheckoutValidateResponse, err error) {
	res = &africastalking.CardCheckoutValidateResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.CardCheckoutValidate"))
	return
}

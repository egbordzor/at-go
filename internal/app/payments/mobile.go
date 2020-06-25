package payments

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Initiate Mobile C2B payments on a mobile subscriber’s device.
func (c *Client) MobileCheckout(ctx context.Context, p *africastalking.MobileCheckoutPayload) (res *africastalking.MobileCheckoutResponse, err error) {
	res = &africastalking.MobileCheckoutResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MobileCheckout"))
	return
}

// Initiate Mobile B2C request
func (c *Client) MobileB2C(ctx context.Context, p *africastalking.MobileB2CPayload) (res *africastalking.MobileB2CResponse, err error) {
	res = &africastalking.MobileB2CResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MobileB2C"))
	return
}

// Initiate Mobile B2B request
func (c *Client) MobileB2B(ctx context.Context, p *africastalking.MobileB2BPayload) (res *africastalking.MobileB2BResponse, err error) {
	res = &africastalking.MobileB2BResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MobileB2B"))
	return
}

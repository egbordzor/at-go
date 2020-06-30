package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/wondenge/at-go/sms"
	"go.uber.org/zap"
)

// Send Bulk SMS
// Payload attributes `p *BulkPayload` passed as key value args.
func (c *Client) SendBulkSMS(ctx context.Context, args map[string]string) (res *sms.BulkResponse, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http requestFormBody: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Incrementally fetch messages from application inbox.
func (c *Client) FetchSMS(ctx context.Context, p *sms.FetchMsgPayload) (res *sms.FetchMsgResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Send Premium SMS
// Payload attributes `p *PremiumPayload` passed as key value args.
func (c *Client) SendPremiumSMS(ctx context.Context, args map[string]string) (res *sms.PremiumSMSResponse, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Generate a checkout token
// Payload attributes `p *CheckoutTokenPayload` passed as key value args.
func (c *Client) NewCheckoutToken(ctx context.Context, args map[string]string) (res *sms.CheckoutTokenResponse, err error) {

	if err := c.checkoutTokenBody(ctx, "POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/checkout/token/create"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Subscribe a phone number
// Payload attributes `p *NewSubPayload` passed as key value args.
func (c *Client) NewPremiumSubscription(ctx context.Context, args map[string]string) (res *sms.NewSubResponse, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription/create"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Incrementally fetch your premium sms subscriptions.
func (c *Client) FetchPremiumSubscription(ctx context.Context, p *sms.FetchSubPayload) (res *sms.FetchSubResponse, err error) {

	if err := c.requestJSONBody(ctx, "GET", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Delete a Premium SMS Subscription
// Payload attributes `p *PurgeSubPayload` passed as key value args.
func (c *Client) PurgePremiumSubscription(ctx context.Context, args map[string]string) (res *sms.PurgeSubResponse, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription/delete"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

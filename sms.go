package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/gen/africastalking"
)

// Send Bulk SMS
func (s *africastalkingsrvc) SendBulkSMS(ctx context.Context, p *africastalking.BulkPayload) (res *africastalking.BulkResponse, err error) {
	res = &africastalking.BulkResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.SendBulkSMS"))
	return
}

// Send Premium SMS
func (s *africastalkingsrvc) SendPremiumSMS(ctx context.Context, p *africastalking.PremiumPayload) (res *africastalking.PremiumSMSResponse, err error) {
	res = &africastalking.PremiumSMSResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.SendPremiumSMS"))
	return
}

// Incrementally fetch messages from application inbox.
func (s *africastalkingsrvc) FetchSMS(ctx context.Context, p *africastalking.FetchMsgPayload) (res *africastalking.FetchMsgResponse, err error) {
	res = &africastalking.FetchMsgResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.FetchSMS"))
	return
}

// Generate a checkout token
func (s *africastalkingsrvc) NewCheckoutToken(ctx context.Context, p *africastalking.CheckoutTokenPayload) (res *africastalking.CheckoutTokenResponse, err error) {
	res = &africastalking.CheckoutTokenResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.NewCheckoutToken"))
	return
}

// Subscribe a phone number
func (s *africastalkingsrvc) NewPremiumSubscription(ctx context.Context, p *africastalking.NewSubPayload) (res *africastalking.NewSubResponse, err error) {
	res = &africastalking.NewSubResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.NewPremiumSubscription"))
	return
}

// Incrementally fetch your premium sms subscriptions.
func (s *africastalkingsrvc) FetchPremiumSubscription(ctx context.Context, p *africastalking.FetchSubPayload) (res *africastalking.FetchSubResponse, err error) {
	res = &africastalking.FetchSubResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.FetchPremiumSubscription"))
	return
}

// Delete a Premium SMS Subscription
func (s *africastalkingsrvc) PurgePremiumSubscription(ctx context.Context, p *africastalking.PurgeSubPayload) (res *africastalking.PurgeSubResponse, err error) {
	res = &africastalking.PurgeSubResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.PurgePremiumSubscription"))
	return
}

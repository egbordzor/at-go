package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/gen/africastalking"
)

// Initiate Mobile C2B payments on a mobile subscriber’s device.
func (s *africastalkingsrvc) MobileCheckout(ctx context.Context, p *africastalking.MobileCheckoutPayload) (res *africastalking.MobileCheckoutResponse, err error) {
	res = &africastalking.MobileCheckoutResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.MobileCheckout"))
	return
}

// Initiate Mobile B2C request
func (s *africastalkingsrvc) MobileB2C(ctx context.Context, p *africastalking.MobileB2CPayload) (res *africastalking.MobileB2CResponse, err error) {
	res = &africastalking.MobileB2CResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.MobileB2C"))
	return
}

// Initiate Mobile B2B request
func (s *africastalkingsrvc) MobileB2B(ctx context.Context, p *africastalking.MobileB2BPayload) (res *africastalking.MobileB2BResponse, err error) {
	res = &africastalking.MobileB2BResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.MobileB2B"))
	return
}

// Collect money into your payment wallet.
func (s *africastalkingsrvc) BankCheckout(ctx context.Context, p *africastalking.BankCheckoutPayload) (res *africastalking.BankCheckoutResponse, err error) {
	res = &africastalking.BankCheckoutResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.Bank Checkout"))
	return
}

// Validate a bank checkout charge request
func (s *africastalkingsrvc) BankCheckoutValidate(ctx context.Context, p *africastalking.BankCheckoutValidatePayload) (res *africastalking.BankCheckoutValidateResponse, err error) {
	res = &africastalking.BankCheckoutValidateResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.BankCheckoutValidate"))
	return
}

// Initiate a bank transfer request.
func (s *africastalkingsrvc) BankTransfer(ctx context.Context, p *africastalking.BankTransferPayload) (res *africastalking.BankTransferResponse, err error) {
	res = &africastalking.BankTransferResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.BankTransfer"))
	return
}

// Collect money into your Payment Wallet
func (s *africastalkingsrvc) CardCheckout(ctx context.Context, p *africastalking.CardCheckoutPayload) (res *africastalking.CardCheckoutResponse, err error) {
	res = &africastalking.CardCheckoutResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.CardCheckout"))
	return
}

// Validate card checkout charge requests
func (s *africastalkingsrvc) CardCheckoutValidate(ctx context.Context, p *africastalking.CardCheckoutValidatePayload) (res *africastalking.CardCheckoutValidateResponse, err error) {
	res = &africastalking.CardCheckoutValidateResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.CardCheckoutValidate"))
	return
}

// Transfer money from one product to another.
func (s *africastalkingsrvc) WalletTransfer(ctx context.Context, p *africastalking.WalletTransferPayload) (res *africastalking.WalletTransferResponse, err error) {
	res = &africastalking.WalletTransferResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.WalletTransfer"))
	return
}

// Move money from a Payment Product to an application stash.
func (s *africastalkingsrvc) TopupStash(ctx context.Context, p *africastalking.TopupStashPayload) (res *africastalking.TopupStashResponse, err error) {
	res = &africastalking.TopupStashResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.TopupStash"))
	return
}

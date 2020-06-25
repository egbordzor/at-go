package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Initiate Mobile C2B payments on a mobile subscriber’s device.
func (c *Client) mobileCheckout(ctx context.Context, p *africastalking.MobileCheckoutPayload) (res *africastalking.MobileCheckoutResponse, err error) {
	res = &africastalking.MobileCheckoutResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MobileCheckout"))
	return
}

// Initiate Mobile B2C request
func (c *Client) mobileB2C(ctx context.Context, p *africastalking.MobileB2CPayload) (res *africastalking.MobileB2CResponse, err error) {
	res = &africastalking.MobileB2CResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MobileB2C"))
	return
}

// Initiate Mobile B2B request
func (c *Client) mobileB2B(ctx context.Context, p *africastalking.MobileB2BPayload) (res *africastalking.MobileB2BResponse, err error) {
	res = &africastalking.MobileB2BResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MobileB2B"))
	return
}

// Collect money into your payment wallet.
func (c *Client) bankCheckout(ctx context.Context, p *africastalking.BankCheckoutPayload) (res *africastalking.BankCheckoutResponse, err error) {
	res = &africastalking.BankCheckoutResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.Bank Checkout"))
	return
}

// Validate a bank checkout charge request
func (c *Client) bankCheckoutValidate(ctx context.Context, p *africastalking.BankCheckoutValidatePayload) (res *africastalking.BankCheckoutValidateResponse, err error) {
	res = &africastalking.BankCheckoutValidateResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.BankCheckoutValidate"))
	return
}

// Initiate a bank transfer request.
func (c *Client) bankTransfer(ctx context.Context, p *africastalking.BankTransferPayload) (res *africastalking.BankTransferResponse, err error) {
	res = &africastalking.BankTransferResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.BankTransfer"))
	return
}

// Collect money into your Payment Wallet
func (c *Client) cardCheckout(ctx context.Context, p *africastalking.CardCheckoutPayload) (res *africastalking.CardCheckoutResponse, err error) {
	res = &africastalking.CardCheckoutResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.CardCheckout"))
	return
}

// Validate card checkout charge requests
func (c *Client) cardCheckoutValidate(ctx context.Context, p *africastalking.CardCheckoutValidatePayload) (res *africastalking.CardCheckoutValidateResponse, err error) {
	res = &africastalking.CardCheckoutValidateResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.CardCheckoutValidate"))
	return
}

// Transfer money from one product to another.
func (c *Client) walletTransfer(ctx context.Context, p *africastalking.WalletTransferPayload) (res *africastalking.WalletTransferResponse, err error) {
	res = &africastalking.WalletTransferResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.WalletTransfer"))
	return
}

// Move money from a Payment Product to an application stash.
func (c *Client) topupStash(ctx context.Context, p *africastalking.TopupStashPayload) (res *africastalking.TopupStashResponse, err error) {
	res = &africastalking.TopupStashResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.TopupStash"))
	return
}

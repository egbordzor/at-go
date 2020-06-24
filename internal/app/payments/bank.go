package payments

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/pkg/gen/africastalking"
)

// Collect money into your payment wallet.
func (c *Client) BankCheckout(ctx context.Context, p *africastalking.BankCheckoutPayload) (res *africastalking.BankCheckoutResponse, err error) {
	res = &africastalking.BankCheckoutResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.Bank Checkout"))
	return
}

// Validate a bank checkout charge request
func (c *Client) BankCheckoutValidate(ctx context.Context, p *africastalking.BankCheckoutValidatePayload) (res *africastalking.BankCheckoutValidateResponse, err error) {
	res = &africastalking.BankCheckoutValidateResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.BankCheckoutValidate"))
	return
}

// Initiate a bank transfer request.
func (c *Client) BankTransfer(ctx context.Context, p *africastalking.BankTransferPayload) (res *africastalking.BankTransferResponse, err error) {
	res = &africastalking.BankTransferResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.BankTransfer"))
	return
}

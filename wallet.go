package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	Wallet interface {
		// Transfer money from one Payment Product to another Payment Product hosted on Africa’s Talking.
		walletTransfer(ctx context.Context, p *at.WalletTransferPayload) (res *at.WalletTransferResponse, err error)
	}
)

// Transfer money from one Payment Product to another Payment Product hosted on Africa’s Talking.
func (c *Client) walletTransfer(ctx context.Context, p *at.WalletTransferPayload) (res *at.WalletTransferResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/transfer/wallet"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.WalletTransferResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	WalletTransferLiveURL    = "https://payments.africastalking.com/transfer/wallet"
	WalletTransferSandboxURL = "https://payments.sandbox.africastalking.com/transfer/wallet"
)

// Transfer money from one product to another.
func (c *Client) walletTransfer(ctx context.Context, p *africastalking.WalletTransferPayload) (res *africastalking.WalletTransferResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/transfer/wallet", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.WalletTransferResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

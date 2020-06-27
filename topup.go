package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	TopupStashLiveURL    = "https://payments.africastalking.com/topup/stash"
	TopupStashSandboxURL = "https://payments.sandbox.africastalking.com/topup/stash"
)

// Move money from a Payment Product to an application stash.
// An application stash is the wallet that funds your service usage expenses.
func (c *Client) topupStash(ctx context.Context, p *africastalking.TopupStashPayload) (res *africastalking.TopupStashResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/topup/stash", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.TopupStashResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

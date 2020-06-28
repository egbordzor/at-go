package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	TopUp interface {

		// Move money from a Payment Product to an application stash.
		// An application stash is the wallet that funds your service usage expenses.
		topupStash(ctx context.Context, p *at.TopupStashPayload) (res *at.TopupStashResponse, err error)
	}
)

// Move money from a Payment Product to an application stash.
// An application stash is the wallet that funds your service usage expenses.
func (c *ATClient) topupStash(ctx context.Context, p *at.TopupStashPayload) (res *at.TopupStashResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/topup/stash"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.TopupStashResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

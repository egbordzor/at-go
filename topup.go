package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	pay "github.com/wondenge/at-go/payments"
	"go.uber.org/zap"
)

// Move money from a Payment Product to an application stash.
// An application stash is the wallet that funds your service usage expenses.
func (c *Client) TopupStash(ctx context.Context, p *pay.TopupStashPayload) (res *pay.TopupStashResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/topup/stash"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

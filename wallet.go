package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	pay "github.com/wondenge/at-go/payments"
	"go.uber.org/zap"
)

// Transfer money from one Payment Product to another Payment Product hosted on Africa’s Talking.
func (c *Client) WalletTransfer(ctx context.Context, p *pay.WalletTransferPayload) (res *pay.WalletTransferResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/transfer/wallet"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

package payments

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Transfer money from one product to another.
func (c *Client) WalletTransfer(ctx context.Context, p *africastalking.WalletTransferPayload) (res *africastalking.WalletTransferResponse, err error) {
	res = &africastalking.WalletTransferResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.WalletTransfer"))
	return
}

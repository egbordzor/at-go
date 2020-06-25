package payments

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Move money from a Payment Product to an application stash.
func (c *Client) TopupStash(ctx context.Context, p *africastalking.TopupStashPayload) (res *africastalking.TopupStashResponse, err error) {
	res = &africastalking.TopupStashResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.TopupStash"))
	return
}

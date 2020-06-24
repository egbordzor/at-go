package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/pkg/gen/africastalking"
)

// Generates a valid auth token
func (s *africastalkingsrvc) Generate(ctx context.Context, p *africastalking.GeneratePayload) (res *africastalking.AccessTokenResponse, err error) {
	res = &africastalking.AccessTokenResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.Generate"))
	return
}

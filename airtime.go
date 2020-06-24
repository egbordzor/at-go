package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/gen/africastalking"
)

// Send Airtime.
func (s *africastalkingsrvc) SendAirtime(ctx context.Context, p *africastalking.AirtimePayload) (res *africastalking.AirtimeResponse, err error) {
	res = &africastalking.AirtimeResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.SendAirtime"))
	return
}



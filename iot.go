package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/gen/africastalking"
)

// Publishes messages to remote devices.
func (s *africastalkingsrvc) PublishIoT(ctx context.Context, p *africastalking.IoTPayload) (res *africastalking.IoTResponse, err error) {
	res = &africastalking.IoTResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.PublishIoT"))
	return
}



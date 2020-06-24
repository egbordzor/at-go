package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/gen/africastalking"
)

// Initiate an application data request.
func (s *africastalkingsrvc) InitiateAppData(ctx context.Context, p string) (res *africastalking.UserResponse, err error) {
	res = &africastalking.UserResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.InitiateAppData"))
	return
}


package atgo

import (
	"github.com/go-kit/kit/log"
	africastalking "github.com/wondenge/at-go/gen/africastalking"
)

// africastalking service example implementation.
// The example methods log the requests and return zero values.
type africastalkingsrvc struct {
	logger log.Logger
}

// NewAfricastalking returns the africastalking service implementation.
func NewAfricastalking(logger log.Logger) africastalking.Service {
	return &africastalkingsrvc{logger}
}

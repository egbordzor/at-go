package edge

import (
	"context"
)

type (
	IoT interface {

		// Publishes messages to remote devices.
		PublishIoT(ctx context.Context, p *IoTPayload) (res *IoTResponse, err error)
	}
)

// IoTPayload is the payload type of the africastalking service PublishIoT
// method.
type IoTPayload struct {
	Username string `form:"username" json:"username" xml:"username"`
	// Device group to which the message is to be sent
	DeviceGroup string `form:"deviceGroup" json:"deviceGroup" xml:"deviceGroup"`
	// Messaging channel to which the message is to be sent.
	Topic string `form:"topic" json:"topic" xml:"topic"`
	// Message packet to be sent to the subscribed devices
	Payload string `form:"payload" json:"payload" xml:"payload"`
}

// IoTResponse is the result type of the africastalking service PublishIoT
// method.
type IoTResponse struct {
	// Response status of the API request.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Verbose response message detailing the status of the HTTP response
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Makes outbound calls.
func (c *Client) makeCall(ctx context.Context, p *africastalking.MakeCallPayload) (res *africastalking.MakeCallResponse, err error) {
	res = &africastalking.MakeCallResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MakeCall"))
	return
}

// Transfers call to another number.
func (c *Client) transferCall(ctx context.Context, p *africastalking.CallTransferPayload) (res *africastalking.Calltransferresponse, err error) {
	res = &africastalking.Calltransferresponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.TransferCall"))
	return
}

// Used when you have more calls than you can handle at one time
func (c *Client) queueCall(ctx context.Context, p *africastalking.QueuedCallsPayload) (res *africastalking.Queuedstatusresult, err error) {
	res = &africastalking.Queuedstatusresult{}
	c.logger.Log("info", fmt.Sprintf("africastalking.Queue"))
	return
}

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
func (c *Client) uploadMedia(ctx context.Context, p *africastalking.UploadMediaFile) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.UploadMedia"))
	return
}

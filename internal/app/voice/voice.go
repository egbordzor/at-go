package voice

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
	"net/http"
	"time"
)

const (

	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = ""

	// APIBaseLive points to the live version of the API
	APIBaseLive = ""
)

type (
	Client struct {
		Client         *http.Client
		APIKey         string
		Username       string
		Sandbox        string
		Live           string
		Token          *africastalking.AccessTokenResponse
		tokenExpiresAt time.Time
		logger         log.Logger // log the requests.
	}
)

// Makes outbound calls.
func (c *Client) MakeCall(ctx context.Context, p *africastalking.MakeCallPayload) (res *africastalking.MakeCallResponse, err error) {
	res = &africastalking.MakeCallResponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.MakeCall"))
	return
}

// Transfers call to another number.
func (c *Client) TransferCall(ctx context.Context, p *africastalking.CallTransferPayload) (res *africastalking.Calltransferresponse, err error) {
	res = &africastalking.Calltransferresponse{}
	c.logger.Log("info", fmt.Sprintf("africastalking.TransferCall"))
	return
}

// Set a text to be read out to the caller.
func (c *Client) Say(ctx context.Context, p *africastalking.Say1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Say"))
	return
}

// Play back an audio file located anywhere on the web.
func (c *Client) Play(ctx context.Context, p *africastalking.Play1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Play"))
	return
}

// Get digits a user enters on their phone in response to a prompt from
// application
func (c *Client) GetDigits(ctx context.Context, p *africastalking.GetDigits1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.GetDigits"))
	return
}

// Connect the user who called your phone number to an external phone number.
func (c *Client) Dial(ctx context.Context, p *africastalking.Dial1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Dial"))
	return
}

// Record a call session into an mp3 file.
func (c *Client) Record(ctx context.Context, p *africastalking.Record1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Record"))
	return
}

// Pass an incoming call to a queue to be handled later.
func (c *Client) Enqueue(ctx context.Context, p *africastalking.Enqueue1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Enqueue"))
	return
}

// Pass the calls enqueued to a separate number to be handled.
func (c *Client) Dequeue(ctx context.Context, p *africastalking.Dequeue1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Dequeue"))
	return
}

// Transfer control of the call to the script whose URL is passed in.
func (c *Client) Redirect(ctx context.Context, p *africastalking.Redirect1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Redirect"))
	return
}

// Reject an incoming call without incurring any usage charges.
func (c *Client) Reject(ctx context.Context, p *africastalking.Reject1) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.Reject"))
	return
}

// Used when you have more calls than you can handle at one time
func (c *Client) Queue(ctx context.Context, p *africastalking.QueuedCallsPayload) (res *africastalking.Queuedstatusresult, err error) {
	res = &africastalking.Queuedstatusresult{}
	c.logger.Log("info", fmt.Sprintf("africastalking.Queue"))
	return
}

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
func (c *Client) UploadMedia(ctx context.Context, p *africastalking.UploadMediaFile) (res string, err error) {
	c.logger.Log("info", fmt.Sprintf("africastalking.UploadMedia"))
	return
}

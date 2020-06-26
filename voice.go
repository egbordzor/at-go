package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	MakeCallLiveURL    = "https://voice.africastalking.com/call"
	MakeCallSandboxURL = "https://voice.sandbox.africastalking.com/call"

	CallTransferLiveURL    = "https://voice.africastalking.com/callTransfer"
	CallTransferSandboxURL = "https://voice.sandbox.africastalking.com/callTransfer"

	QueuedCallsLiveURL    = "https://voice.africastalking.com/queueStatus"
	QueuedCallsSandboxURL = "https://voice.sandbox.africastalking.com/queueStatus"

	MediaUploadLiveURL    = "https://voice.africastalking.com/mediaUpload"
	MediaUploadSandboxURL = "https://voice.sandbox.africastalking.com/mediaUpload"
)

// Makes outbound calls.
func (c *Client) makeCall(ctx context.Context, p *africastalking.MakeCallPayload) (res *africastalking.MakeCallResponse, err error) {

	req, err := c.NewRequest("POST", "https://voice.sandbox.africastalking.com/call", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.MakeCallResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Transfers call to another number.
func (c *Client) transferCall(ctx context.Context, p *africastalking.CallTransferPayload) (res *africastalking.Calltransferresponse, err error) {

	req, err := c.NewRequest("POST", "", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.Calltransferresponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Used when you have more calls than you can handle at one time
func (c *Client) queuedCall(ctx context.Context, p *africastalking.QueuedCallsPayload) (res *africastalking.Queuedstatusresult, err error) {

	req, err := c.NewRequest("POST", "https://voice.sandbox.africastalking.com/queueStatus", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.Queuedstatusresult{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
func (c *Client) uploadMedia(ctx context.Context, p *africastalking.UploadMediaFile) (res string, err error) {

	req, err := c.NewRequest("POST", "https://voice.sandbox.africastalking.com/mediaUpload", p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppApiKey")

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

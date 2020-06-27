package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Sandbox Endpoints
const MakeCallSandboxURL = "https://voice.sandbox.africastalking.com/call"
const CallTransferSandboxURL = "https://voice.sandbox.africastalking.com/callTransfer"
const QueuedCallsSandboxURL = "https://voice.sandbox.africastalking.com/queueStatus"
const MediaUploadSandboxURL = "https://voice.sandbox.africastalking.com/mediaUpload"

// Production Endpoints
const MakeCallLiveURL = "https://voice.africastalking.com/call"
const CallTransferLiveURL = "https://voice.africastalking.com/callTransfer"
const QueuedCallsLiveURL = "https://voice.africastalking.com/queueStatus"
const MediaUploadLiveURL = "https://voice.africastalking.com/mediaUpload"

// Makes outbound calls.
func (c *Client) makeCall(ctx context.Context, p *africastalking.MakeCallPayload) (res *africastalking.MakeCallResponse, err error) {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.voiceEndpoint, "/call"), p)
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
func (c *Client) transferCall(ctx context.Context, p *africastalking.CallTransferPayload) (res *africastalking.CallTransferResponse, err error) {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.voiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.CallTransferResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Used when you have more calls than you can handle at one time
func (c *Client) queuedCall(ctx context.Context, p *africastalking.QueuedCallsPayload) (res *africastalking.QueuedStatusResult, err error) {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.voiceEndpoint, "/queueStatus"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.QueuedStatusResult{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
func (c *Client) uploadMedia(ctx context.Context, p *africastalking.UploadMediaFile) (res string, err error) {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.voiceEndpoint, "/mediaUpload"), p)
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

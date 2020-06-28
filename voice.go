package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	Voice interface {

		// Makes outbound calls.
		makeCall(ctx context.Context, p *at.MakeCallPayload) (res *at.MakeCallResponse, err error)

		// Transfers call to another number.
		transferCall(ctx context.Context, p *at.CallTransferPayload) (res *at.CallTransferResponse, err error)

		// Used when you have more calls than you can handle at one time
		queuedCall(ctx context.Context, p *at.QueuedCallsPayload) (res *at.QueuedStatusResult, err error)

		// Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav
		uploadMedia(ctx context.Context, p *at.UploadMediaFile) (res string, err error)
	}

	Actions interface {

		// Set a text to be read out to the caller.
		say(ctx context.Context, p *at.SayPayload) (res string, err error)

		// Play back an audio file located anywhere on the web.
		play(ctx context.Context, p *at.PlayPayload) (res string, err error)

		// Get digits a user enters on their phone in response to a prompt from
		// application
		getDigits(ctx context.Context, p *at.GetDigitsPayload) (res string, err error)

		// Connect the user who called your phone number to an external phone number.
		dial(ctx context.Context, p *at.DialPayload) (res string, err error)

		// Record a call session into an mp3 file.
		record(ctx context.Context, p *at.RecordPayload) (res string, err error)

		// Pass an incoming call to a queue to be handled later.
		enqueue(ctx context.Context, p *at.EnqueuePayload) (res string, err error)

		// Pass the calls enqueued to a separate number to be handled.
		dequeue(ctx context.Context, p *at.DequeuePayload) (res string, err error)

		// Transfer control of the call to the script whose URL is passed in.
		redirect(ctx context.Context, p *at.RedirectPayload) (res string, err error)

		// Reject an incoming call without incurring any usage charges.
		reject(ctx context.Context, p *at.RejectPayload) (res string, err error)
	}
)

// Makes outbound calls.
func (c *ATClient) makeCall(ctx context.Context, p *at.MakeCallPayload) (res *at.MakeCallResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/call"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.MakeCallResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Transfers call to another number.
func (c *ATClient) transferCall(ctx context.Context, p *at.CallTransferPayload) (res *at.CallTransferResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.CallTransferResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Used when you have more calls than you can handle at one time
func (c *ATClient) queuedCall(ctx context.Context, p *at.QueuedCallsPayload) (res *at.QueuedStatusResult, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/queueStatus"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.QueuedStatusResult{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
func (c *ATClient) uploadMedia(ctx context.Context, p *at.UploadMediaFile) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/mediaUpload"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Set a text to be read out to the caller.
func (c *ATClient) say(ctx context.Context, p *at.SayPayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Play back an audio file located anywhere on the web.
func (c *ATClient) play(ctx context.Context, p *at.PlayPayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Get digits a user enters on their phone in response to a prompt from
// application
func (c *ATClient) getDigits(ctx context.Context, p *at.GetDigitsPayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Connect the user who called your phone number to an external phone number.
func (c *ATClient) dial(ctx context.Context, p *at.DialPayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Record a call session into an mp3 file.
func (c *ATClient) record(ctx context.Context, p *at.RecordPayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Pass an incoming call to a queue to be handled later.
func (c *ATClient) enqueue(ctx context.Context, p *at.EnqueuePayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Pass the calls enqueued to a separate number to be handled.
func (c *ATClient) dequeue(ctx context.Context, p *at.DequeuePayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Transfer control of the call to the script whose URL is passed in.
func (c *ATClient) redirect(ctx context.Context, p *at.RedirectPayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

// Reject an incoming call without incurring any usage charges.
func (c *ATClient) reject(ctx context.Context, p *at.RejectPayload) (res string, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), p)
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return "", err
	}
	return res, nil
}

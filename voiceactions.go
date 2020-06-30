package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	"go.uber.org/zap"
)

// Set a text to be read out to the caller.
// Payload attributes `p *SayPayload` passed as key value args.
func (c *Client) Say(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Play back an audio file located anywhere on the web.
// Payload attributes `p *PlayPayload` passed as key value args.
func (c *Client) Play(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Get digits a user enters on their phone in response to a prompt from
// application
// Payload attributes `p *GetDigitsPayload` passed as key value args.
func (c *Client) GetDigits(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Connect the user who called your phone number to an external phone number.
// Payload attributes `p *DialPayload` passed as key value args.
func (c *Client) Dial(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Record a calls session into an mp3 file.
// Payload attributes `p *RecordPayload` passed as key value args.
func (c *Client) Record(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Pass an incoming calls to a queue to be handled later.
// Payload attributes `p *EnqueuePayload` passed as key value args.
func (c *Client) Enqueue(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Pass the calls enqueued to a separate number to be handled.
// Payload attributes `p *DequeuePayload` passed as key value args.
func (c *Client) Dequeue(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Transfer control of the calls to the script whose URL is passed in.
// Payload attributes `p *RedirectPayload` passed as key value args.
func (c *Client) Redirect(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Reject an incoming calls without incurring any usage charges.
// Payload attributes `p *RejectPayload` passed as key value args.
func (c *Client) Reject(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

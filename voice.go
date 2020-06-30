package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/wondenge/at-go/voice"
	"go.uber.org/zap"
)

// Makes outbound calls.
// Payload attributes `p *MakeCallPayload` passed as key value args.
func (c *Client) MakeCall(ctx context.Context, args map[string]string) (res *voice.MakeCallResponse, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/calls"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Transfers calls to another number.
// Payload attributes `p *CallTransferPayload` passed as key value args.
func (c *Client) TransferCall(ctx context.Context, args map[string]string) (res *voice.CallTransferResponse, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Used when you have more calls than you can handle africastalking one time.
// Payload attributes `p *QueuedCallsPayload` passed as key value args.
func (c *Client) QueuedCall(ctx context.Context, args map[string]string) (res *voice.QueuedStatusResult, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/queueStatus"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
// Payload attributes `p *UploadMediaFile` passed as key value args.
func (c *Client) UploadMedia(ctx context.Context, args map[string]string) (res string, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/mediaUpload"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

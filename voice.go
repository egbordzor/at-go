package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/gen/africastalking"
)

// Makes outbound calls.
func (s *africastalkingsrvc) MakeCall(ctx context.Context, p *africastalking.MakeCallPayload) (res *africastalking.MakeCallResponse, err error) {
	res = &africastalking.MakeCallResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.MakeCall"))
	return
}

// Transfers call to another number.
func (s *africastalkingsrvc) TransferCall(ctx context.Context, p *africastalking.CallTransferPayload) (res *africastalking.Calltransferresponse, err error) {
	res = &africastalking.Calltransferresponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.TransferCall"))
	return
}

// Set a text to be read out to the caller.
func (s *africastalkingsrvc) Say(ctx context.Context, p *africastalking.Say1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Say"))
	return
}

// Play back an audio file located anywhere on the web.
func (s *africastalkingsrvc) Play(ctx context.Context, p *africastalking.Play1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Play"))
	return
}

// Get digits a user enters on their phone in response to a prompt from
// application
func (s *africastalkingsrvc) GetDigits(ctx context.Context, p *africastalking.GetDigits1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.GetDigits"))
	return
}

// Connect the user who called your phone number to an external phone number.
func (s *africastalkingsrvc) Dial(ctx context.Context, p *africastalking.Dial1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Dial"))
	return
}

// Record a call session into an mp3 file.
func (s *africastalkingsrvc) Record(ctx context.Context, p *africastalking.Record1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Record"))
	return
}

// Pass an incoming call to a queue to be handled later.
func (s *africastalkingsrvc) Enqueue(ctx context.Context, p *africastalking.Enqueue1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Enqueue"))
	return
}

// Pass the calls enqueued to a separate number to be handled.
func (s *africastalkingsrvc) Dequeue(ctx context.Context, p *africastalking.Dequeue1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Dequeue"))
	return
}

// Transfer control of the call to the script whose URL is passed in.
func (s *africastalkingsrvc) Redirect(ctx context.Context, p *africastalking.Redirect1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Redirect"))
	return
}

// Reject an incoming call without incurring any usage charges.
func (s *africastalkingsrvc) Reject(ctx context.Context, p *africastalking.Reject1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Reject"))
	return
}

// Used when you have more calls than you can handle at one time
func (s *africastalkingsrvc) Queue(ctx context.Context, p *africastalking.QueuedCallsPayload) (res *africastalking.Queuedstatusresult, err error) {
	res = &africastalking.Queuedstatusresult{}
	s.logger.Log("info", fmt.Sprintf("africastalking.Queue"))
	return
}

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
func (s *africastalkingsrvc) UploadMedia(ctx context.Context, p *africastalking.UploadMediaFile) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.UploadMedia"))
	return
}






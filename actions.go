package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Set a text to be read out to the caller.
func (s *africastalkingsrvc) Say(ctx context.Context, p *africastalking.SayPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Say"))
	return
}

// Play back an audio file located anywhere on the web.
func (s *africastalkingsrvc) Play(ctx context.Context, p *africastalking.PlayPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Play"))
	return
}

// Get digits a user enters on their phone in response to a prompt from
// application
func (s *africastalkingsrvc) GetDigits(ctx context.Context, p *africastalking.GetDigitsPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.GetDigits"))
	return
}

// Connect the user who called your phone number to an external phone number.
func (s *africastalkingsrvc) Dial(ctx context.Context, p *africastalking.DialPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Dial"))
	return
}

// Record a call session into an mp3 file.
func (s *africastalkingsrvc) Record(ctx context.Context, p *africastalking.RecordPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Record"))
	return
}

// Pass an incoming call to a queue to be handled later.
func (s *africastalkingsrvc) Enqueue(ctx context.Context, p *africastalking.EnqueuePayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Enqueue"))
	return
}

// Pass the calls enqueued to a separate number to be handled.
func (s *africastalkingsrvc) Dequeue(ctx context.Context, p *africastalking.DequeuePayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Dequeue"))
	return
}

// Transfer control of the call to the script whose URL is passed in.
func (s *africastalkingsrvc) Redirect(ctx context.Context, p *africastalking.RedirectPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Redirect"))
	return
}

// Reject an incoming call without incurring any usage charges.
func (s *africastalkingsrvc) Reject(ctx context.Context, p *africastalking.RejectPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.Reject"))
	return
}
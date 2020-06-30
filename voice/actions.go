package voice

// SayPayload is the payload type of the  Say method.
type SayPayload struct {

	// This parameter specifies the calls to use
	Voice string `form:"calls,omitempty" json:"calls,omitempty" xml:"calls,omitempty"`

	// Instructs AT to play a beep after reading the text contained in the request
	PlayBeep bool `form:"playBeep,omitempty" json:"playBeep,omitempty" xml:"playBeep,omitempty"`
}

// PlayPayload is the payload type of the Play method.
type PlayPayload struct {

	// A valid URL that contains a link to the file to be played.
	URL string `form:"url" json:"url" xml:"url"`
}

// GetDigitsPayload is the payload type of the GetDigits method.
type GetDigitsPayload struct {

	// Instructs AT to forward results of the GetDigits action to the URL value passed in.
	CallbackURL string `form:"callbackUrl,omitempty" json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`

	// This shows the number of digits you would like to grab from the user input.
	NumDigits string `form:"numDigits,omitempty" json:"numDigits,omitempty" xml:"numDigits,omitempty"`

	// Timeout (in seconds) for getting the digits, after which the system moves on to the next element.
	Timeout string `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`

	// The key which will terminate the action of getting digits.
	FinishOnKey string `form:"finishOnKey,omitempty" json:"finishOnKey,omitempty" xml:"finishOnKey,omitempty"`
}

// DialPayload is the payload type of the Dial method.
type DialPayload struct {
	PhoneNumbers string `form:"phoneNumbers" json:"phoneNumbers" xml:"phoneNumbers"`
	Record       string `form:"record,omitempty" json:"record,omitempty" xml:"record,omitempty"`
	Sequential   string `form:"sequential,omitempty" json:"sequential,omitempty" xml:"sequential,omitempty"`
	CallerID     string `form:"callerId,omitempty" json:"callerId,omitempty" xml:"callerId,omitempty"`
	RingBackTone string `form:"ringBackTone,omitempty" json:"ringBackTone,omitempty" xml:"ringBackTone,omitempty"`
	// This contains the maximum amount of time in seconds a calls should take.
	MaxDuration string `form:"maxDuration,omitempty" json:"maxDuration,omitempty" xml:"maxDuration,omitempty"`
}

// RecordPayload is the payload type of the Record method.
type RecordPayload struct {
	FinishOnKey string `form:"finishOnKey,omitempty" json:"finishOnKey,omitempty" xml:"finishOnKey,omitempty"`
	MaxLength   string `form:"maxLength,omitempty" json:"maxLength,omitempty" xml:"maxLength,omitempty"`
	Timeout     string `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	TrimSilence string `form:"trimSilence,omitempty" json:"trimSilence,omitempty" xml:"trimSilence,omitempty"`
	PlayBeep    string `form:"playBeep,omitempty" json:"playBeep,omitempty" xml:"playBeep,omitempty"`
	CallbackURL string `form:"callbackUrl,omitempty" json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
}

// EnqueuePayload is the payload type of the Enqueue method.
type EnqueuePayload struct {
	HoldMusic string `form:"holdMusic,omitempty" json:"holdMusic,omitempty" xml:"holdMusic,omitempty"`
	Name      string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// DequeuePayload is the payload type of the Dequeue method.
type DequeuePayload struct {
	PhoneNumber string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	Name        string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// RedirectPayload is the payload type of the Redirect method.
type RedirectPayload struct {

	// Reject
	Reject string `form:"Reject,omitempty" json:"Reject,omitempty" xml:"Reject,omitempty"`
}

// RejectPayload is the payload type of the Reject method.
type RejectPayload struct {

	// Redirect
	Redirect string `form:"Redirect,omitempty" json:"Redirect,omitempty" xml:"Redirect,omitempty"`
}

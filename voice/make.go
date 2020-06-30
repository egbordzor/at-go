package voice

// MakeCallPayload is the payload type of the MakeCall method.
type MakeCallPayload struct {

	// Africa’s Talking Username
	Username string `form:"username" json:"username" xml:"username"`

	// Africa’s Talking phone number
	From string `form:"from" json:"from" xml:"from"`

	// A comma separated string of recipients’ phone numbers.
	To string `form:"to" json:"to" xml:"to"`

	// Variable sent to Events Callback URL used to tag the calls
	ClientRequestID string `form:"clientRequestId,omitempty" json:"clientRequestId,omitempty" xml:"clientRequestId,omitempty"`
}

// MakeCallResponse is the result type of the MakeCall method.
type MakeCallResponse struct {
	Entries []CallEntry `form:"entries,omitempty" json:"entries,omitempty" xml:"entries,omitempty"`

	// Error message if ENTIRE request was rejected by the API.
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

type CallEntry struct {

	// The phone number queued.
	PhoneNumber string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	Status      string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`

	// A unique id for the request associated to this phone number
	SessionID string `form:"sessionId,omitempty" json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

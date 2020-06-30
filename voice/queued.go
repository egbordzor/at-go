package voice

// QueuedCallsPayload is the payload type of the Queue method.
type QueuedCallsPayload struct {

	// Your Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`

	// List of one or more numbers mapped to your Africa’s Talking account.
	PhoneNumbers string `form:"phoneNumbers" json:"phoneNumbers" xml:"phoneNumbers"`
}

// QueuedStatusResult is the result type of the Queue method.
type QueuedStatusResult struct {
	Entries []QueuedStatusEntry `form:"Entries,omitempty" json:"Entries,omitempty" xml:"Entries,omitempty"`

	// Error Message
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

type QueuedStatusEntry struct {
	PhoneNumber string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	QueueName   string `form:"queueName,omitempty" json:"queueName,omitempty" xml:"queueName,omitempty"`
	NumCalls    string `form:"numCalls,omitempty" json:"numCalls,omitempty" xml:"numCalls,omitempty"`
}

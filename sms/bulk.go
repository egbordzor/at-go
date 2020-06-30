package sms

// BulkPayload is the payload type of the SendBulkSMS method.
type BulkPayload struct {

	// Africa’s Talking application username
	Username string `json:"username"`

	// Recipients’ phone numbers
	To string `json:"to"`

	// Message to be sent
	Message string `json:"message"`

	// Registered Short Code or Alphanumeric
	From string `json:"from,omitempty"`

	// Used by MSP to determine who gets billed for a message
	BulkSMSMode int `json:"bulkSMSMode,omitempty"`

	// Used for Bulk SMS clients
	Enqueue int `json:"enqueue,omitempty"`

	// The keyword to be used for a premium service
	Keyword string `json:"keyword,omitempty"`

	// Used for premium services to send OnDemand messages
	LinkID string `json:"linkId,omitempty"`

	// No. of hours subscription message should be retried
	RetryDurationInHours string `json:"retryDurationInHours,omitempty"`
}

// BulkResponse is the result type of the SendBulkSMS method.
type BulkResponse struct {
	SMSMessageData BulkSMSMessageData `json:"SMSMessageData,omitempty"`
}

// A Map detailing the eventual result of the sms request
type BulkSMSMessageData struct {

	// A summary of the total number of recipients the sms was sent to and the total cost incurred.
	Messages string `json:"Messages,omitempty"`

	// A list of recipients included in the original request.
	Recipients []Recipients `json:"Recipients,omitempty"`
}

// Recipient Attributes
type Recipients struct {

	// This corresponds to the status of the request
	StatusCode int `json:"statusCode,omitempty"`

	// The recipient’s phone number
	Number string `json:"number,omitempty"`

	// A string indicating whether the sms was sent to this recipient or not.
	Status string `json:"status,omitempty"`

	// Amount incurred to send this sms.
	Cost string `json:"cost,omitempty"`

	// The messageId received when the sms was sent.
	MessageID string `json:"messageId,omitempty"`
}



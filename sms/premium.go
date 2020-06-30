package sms

// PremiumPayload is the payload type of the SendPremiumSMS method.
type PremiumPayload struct {

	// Africa’s Talking application username
	Username string `form:"username" json:"username" xml:"username"`

	// Recipients’ phone numbers
	To string `form:"to" json:"to" xml:"to"`

	// Message to be sent
	Message string `form:"message" json:"message" xml:"message"`

	// Registered Short Code or Alphanumeric
	From string `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`

	// Used by MSP to determine who gets billed for a message
	BulkSMSMode int `form:"bulkSMSMode,omitempty" json:"bulkSMSMode,omitempty" xml:"bulkSMSMode,omitempty"`

	// Used for Bulk SMS clients
	Enqueue int `form:"enqueue,omitempty" json:"enqueue,omitempty" xml:"enqueue,omitempty"`

	// The keyword to be used for a premium service
	Keyword string `form:"keyword,omitempty" json:"keyword,omitempty" xml:"keyword,omitempty"`

	// Used for premium services to send OnDemand messages
	LinkID string `form:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`

	// No. of hours subscription message should be retried
	RetryDurationInHours string `form:"retryDurationInHours,omitempty" json:"retryDurationInHours,omitempty" xml:"retryDurationInHours,omitempty"`
}

// PremiumSMSResponse is the result type of the SendPremiumSMS method.
type PremiumSMSResponse struct {
	SMSMessageData PremiumSMSMessageData `form:"SMSMessageData,omitempty" json:"SMSMessageData,omitempty" xml:"SMSMessageData,omitempty"`
}

// A Map detailing the eventual result of the sms request
type PremiumSMSMessageData struct {

	// A summary of the total number of recipients the sms was sent to and the total cost incurred.
	Messages string `form:"Messages,omitempty" json:"Messages,omitempty" xml:"Messages,omitempty"`

	// A list of recipients included in the original request.
	Recipients []Recipients `form:"Recipients,omitempty" json:"Recipients,omitempty" xml:"Recipients,omitempty"`
}

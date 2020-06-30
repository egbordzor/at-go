package sms

// FetchMsgPayload is the payload type of the  FetchSMS method.
type FetchMsgPayload struct {

	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`

	// This is the id of the message that you last processed.
	LastReceivedID string `form:"lastReceivedId,omitempty" json:"lastReceivedId,omitempty" xml:"lastReceivedId,omitempty"`
}

// FetchMsgResponse is the result type of the FetchSMS method.
type FetchMsgResponse struct {
	SMSMessageData FetchSMSMessageData `form:"SMSMessageData,omitempty" json:"SMSMessageData,omitempty" xml:"SMSMessageData,omitempty"`
}

// FetchSubPayload is the payload type of the FetchPremiumSubscription method.
type FetchSubPayload struct {

	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`

	// Premium short code mapped to your account
	ShortCode string `form:"shortCode" json:"shortCode" xml:"shortCode"`

	// Premium keyword under short code mapped to your account.
	Keyword string `form:"keyword" json:"keyword" xml:"keyword"`

	// ID of the subscription you believe to be your last.
	LastReceivedID string `form:"lastReceivedId,omitempty" json:"lastReceivedId,omitempty" xml:"lastReceivedId,omitempty"`
}

// FetchSubResponse is the result type of the FetchPremiumSubscription method.
type FetchSubResponse struct {

	// A list of subscriptions made to the product.
	Subscriptions []Subscriptions `form:"Subscriptions,omitempty" json:"Subscriptions,omitempty" xml:"Subscriptions,omitempty"`
}

type FetchSMSMessageData struct {
	Messages []Messages `form:"Messages,omitempty" json:"Messages,omitempty" xml:"Messages,omitempty"`
}

// A list of messages from your inbox.
type Messages struct {

	// A unique identifier attached to each incoming message.
	LinkID string `form:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`

	// The content of the message received.
	Text string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`

	// Your registered short code that the sms was sent out to.
	To string `form:"to,omitempty" json:"to,omitempty" xml:"to,omitempty"`

	// The id of the message.
	ID int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`

	// The date when the sms was sent.
	Date string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`

	// The sender’s phone number.
	From string `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
}

type Subscriptions struct {

	// The id of the subscription
	ID int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`

	// The phone number subscribed to the product.
	Number string `form:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`

	// Timestamp when the subscription was made.
	Date string `form:"Date,omitempty" json:"Date,omitempty" xml:"Date,omitempty"`
}

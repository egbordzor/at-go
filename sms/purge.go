package sms

// PurgeSubPayload is the payload type of the PurgePremiumSubscription method.
type PurgeSubPayload struct {

	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`

	// Premium short code mapped to your account
	ShortCode string `form:"shortCode" json:"shortCode" xml:"shortCode"`

	// Premium keyword under short code mapped to your account.
	Keyword string `form:"keyword" json:"keyword" xml:"keyword"`

	// The phoneNumber to be unsubscribed.
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
}

// PurgeSubResponse is the result type of the PurgePremiumSubscription method.
type PurgeSubResponse struct {

	// Indicates whether the phone number was successfully unsubscribed or not.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`

	// Describes status of the delete subscription request.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

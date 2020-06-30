package sms

// NewSubPayload is the payload type of the NewPremiumSubscription method.
type NewSubPayload struct {

	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`

	// Premium short code mapped to your account
	ShortCode string `form:"shortCode" json:"shortCode" xml:"shortCode"`

	// Premium keyword under short code mapped to your account.
	Keyword string `form:"keyword" json:"keyword" xml:"keyword"`

	// The phoneNumber to be subscribed
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`

	// Token used to validate the subscription request
	CheckoutToken string `form:"checkoutToken" json:"checkoutToken" xml:"checkoutToken"`
}

// NewSubResponse is the result type of the NewPremiumSubscription method.
type NewSubResponse struct {

	// Indicates whether the prompt to subscribe to shortcode was successfully raised or not.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`

	// Describes status of the create subscription request.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}


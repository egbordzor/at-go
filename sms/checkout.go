package sms

// CheckoutTokenPayload is the payload type of the NewCheckoutToken method.
type CheckoutTokenPayload struct {

	// Mobile phone number you want to create a subscription for.
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
}

// CheckoutTokenResponse is the result type of the NewCheckoutToken method.
type CheckoutTokenResponse struct {

	// A description of the status of the request.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`

	// The checkout token to be used
	Token string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

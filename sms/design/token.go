package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Create Checkout token generates a checkout token
// required in order to subscribe a phone number.
// This is an open endpoint and does not need your
// authentication credentials.
// Live: https://api.africastalking.com/checkout/token/create
// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
var CreateCheckoutTokenPayload = Type("CreateCheckoutTokenPayload", func() {
	Description("Authorizes a premium SMS subscription.")

	Attribute("Content-Type", String, func() {
		Description("The requests content type")
		Default("application/x-www-form-urlencoded")
	})
	Attribute("phoneNumber", String, func() {
		Description("Mobile phone number you want to create a subscription for.")
	})
	Required("Content-Type", "phoneNumber")
})

var CreateCheckoutTokenResponse = ResultType("CreateCheckoutTokenResponse", func() {
	Description("Generated checkout token response")
	TypeName("CreateCheckoutTokenResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("description", String, func() {
			Description("A description of the status of the request.")
			Enum("Success", "Failed")
			Example("Success")
		})

		// This token shall be expected later when
		// initiating the actual create subscription request.
		// The tokens expire after 10 minutes and are limited to 2
		// tokens for a 5 minute window for each source IP Address.
		Attribute("token", String, func() {
			Description("The checkout token to be used")
			Example("CkTkn_SampleCkTknId123")
		})
	})
	View("default", func() {
		Attribute("description")
		Attribute("token")
	})
})

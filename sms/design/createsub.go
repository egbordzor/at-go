package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Subscribe a phone number
// Live: https://api.africastalking.com/version1/subscription/create
// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create
var CreateSubPayload = Type("CreateSubPayload", func() {
	Description("subscribe a phone number payload")

	Attribute("apiKey", String, func() {
		Description("Africa’s Talking application apiKey")
		Example("")
	})
	Attribute("Content-Type", String, func() {
		Description("The requests content type")
		Default("application/x-www-form-urlencoded")
	})
	Attribute("Accept", String, func() {
		Description("The requests response type.")
		Default("application/json")
	})
	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
		Default("sandbox")
	})
	Attribute("shortCode", String, func() {
		Description("Premium short code mapped to your account")
		Example("")
	})
	Attribute("keyword", String, func() {
		Description("Premium keyword under short code mapped to your account.")
		Example("")
	})
	Attribute("phoneNumber", String, func() {
		Description("The phoneNumber to be subscribed")
		Example("")
	})
	Attribute("checkoutToken", String, func() {
		Description("Token used to validate the subscription request")
		Example("")
	})
	Required("apiKey", "Content-Type", "username", "shortCode", "keyword", "phoneNumber", "checkoutToken")
})

var CreateSubResponse = ResultType("CreateSubResponse", func() {
	Description("subscribe a phone number response")
	TypeName("CreateSubResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Indicates whether the prompt to subscribe to shortcode was successfully raised or not.")
			Enum("Success", "Failed")
			Example("Success")
		})
		Attribute("description", String, func() {
			Description("Describes status of the create subscription request.")
			Example("Waiting for user input")
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
	})
})

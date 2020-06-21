package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Fetch Subscriptions
// Live: https://api.africastalking.com/version1/subscription
// Sandbox: https://api.sandbox.africastalking.com/verson1/subscription
var FetchSubPayload = Type("FetchSubPayload", func() {
	Description("Incrementally fetch Premium SMS Subscriptions.")

	Attribute("apiKey", String, func() {
		Description("Africa’s Talking application apiKey")
		Example("")
	})
	Attribute("Content-Type", String, func() {
		Description("The requests content type")
		Default("application/json")
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
	Attribute("lastReceivedId", String, func() {
		Description("ID of the subscription you believe to be your last.")
		Default("0") // Set it to 0 to for the first time.
	})
	Required(
		"apiKey", "Content-Type", "username", "shortCode", "keyword")
})

var FetchSubResponse = ResultType("FetchSubResponse", func() {
	Description("Fetch Subscriptions response")
	TypeName("FetchSubResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("Subscriptions", ArrayOf(Subscriptions), func() {
			Description("A list of subscriptions made to the product.")
			MinLength(1)
		})
	})

	View("default", func() {
		Attribute("Subscriptions")
	})
})

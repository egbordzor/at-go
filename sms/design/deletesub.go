package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Delete Subscription
// Delete a premium sms subscription
// Live: https://api.africastalking.com/version1/subscription/delete
// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete
var PurgeSubPayload = Type("PurgeSubPayload", func() {
	Description("Delete a premium sms subscription.")

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
		Description("The phoneNumber to be unsubscribed.")
		Example("")
	})
	Required("apiKey", "Content-Type", "username", "shortCode", "keyword", "phoneNumber")
})

var PurgeSubResponse = ResultType("PurgeSubResponse", func() {
	Description("Delete Subscription Response")
	TypeName("PurgeSubResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Indicates whether the phone number was successfully unsubscribed or not.")
			Enum("Success", "Failed")
			Example("Success")
		})
		Attribute("description", String, func() {
			Description("Describes status of the delete subscription request.")
			Example("Succeeded")
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
	})
})

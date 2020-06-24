package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Fetch Messages
// You can incrementally fetch your application inbox
// Live: https://api.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
var FetchMsgPayload = Type("FetchMsgPayload", func() {
	Description("Incrementally fetches our application inbox")

	Attribute("apiKey", String, func() {
		Description("Africa’s Talking application apiKey")
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
	Attribute("lastReceivedId", String, func() {
		Description("This is the id of the message that you last processed.")
		Default("0") // The default is 0.
	})
	Required("apiKey", "Content-Type", "username")
})

var FetchMsgResponse = ResultType("FetchMsgResponse", func() {
	Description("Fetch Messages Response")
	TypeName("FetchMsgResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("SMSMessageData", func() {
			Description("A Map detailing the eventual result of the sms request")
			Attribute("Messages", ArrayOf(Messages), func() {
				Description("A list of recipients included in the original request.")
				MinLength(1)
			})
		})
	})

	View("default", func() {
		Attribute("SMSMessageData")
	})
})

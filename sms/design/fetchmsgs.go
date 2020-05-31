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

var Messages = Type("Messages", func() {
	Description("A list of messages from your inbox.")

	Attribute("linkId", String, func() {
		Description("A unique identifier attached to each incoming message.")
		Example("SampleLinkId123")
	})
	Attribute("text", String, func() {
		Description("The content of the message received.")
		Example("Hello")
	})
	Attribute("to", String, func() {
		Description("Your registered short code that the sms was sent out to.")
		Example("28901")
	})
	Attribute("id", Int, func() {
		Description("The id of the message.")
		Example(15071)
	})
	Attribute("date", String, func() {
		Description("The date when the sms was sent.")
		Example("2018-03-19T08:34:18.445Z")
	})
	Attribute("from", String, func() {
		Description("The sender’s phone number.")
		Example("+254711XXXYYY")
	})
})

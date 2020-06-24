package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var BulkSMSMessageData = Type("BulkSMSMessageData", func() {
	Description("A Map detailing the eventual result of the sms request")

	Attribute("Messages", String, func() {
		Description("A summary of the total number of recipients the sms was sent to and the total cost incurred.")
		Pattern("[a-zA-Z]+")
		Example("Sent to 1/1 Total Cost: KES 0.8000")
	})
	Attribute("Recipients", ArrayOf(Recipients), func() {
		Description("A list of recipients included in the original request.")
		MinLength(1)
	})
})

var PremiumSMSMessageData = Type("PremiumSMSMessageData", func() {
	Description("A Map detailing the eventual result of the sms request")

	Attribute("Messages", String, func() {
		Description("A summary of the total number of recipients the sms was sent to and the total cost incurred.")
		Pattern("[a-zA-Z]+")
		Example("Sent to 1/1 Total Cost: KES 0.8000")
	})
	Attribute("Recipients", ArrayOf(Recipients), func() {
		Description("A list of recipients included in the original request.")
		MinLength(1)
	})
})

var FetchSMSMessageData = Type("FetchSMSMessageData", func() {
	Attribute("Messages", ArrayOf(Messages), func() {
		Description("")
		MinLength(1)
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

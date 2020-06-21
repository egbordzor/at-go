package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var Recipient = Type("Recipient", func() {
	Description("Recipient Attributes")
	Attribute("statusCode", Int, func() {
		Description("This corresponds to the status of the request")
		Enum(
			100, // 100: Processed
			101, // 101: Sent
			102, //102: Queued
			401, // 401: RiskHold
			402, // 402: InvalidSenderId
			403, // 403: InvalidPhoneNumber
			404, // 404: UnsupportedNumberType
			405, // 405: InsufficientBalance
			406, // 406: UserInBlacklist
			407, // 407: CouldNotRoute
			408, // 500: InternalServerError
			501, // 501: GatewayError
			502, // 502: RejectedByGateway
		)
		Example(101)
	})
	Attribute("number", String, func() {
		Description("The recipient’s phone number")
		Example("+254711XXXYYY")
	})

	// This does not indicate the delivery status
	// of the sms to this recipient.
	Attribute("status", String, func() {
		Description("A string indicating whether the sms was sent to this recipient or not.")
		Example("Success")
	})

	// The format of this string is:
	// (3-digit Currency Code)(space)(Decimal Value)
	Attribute("cost", String, func() {
		Description("Amount incurred to send this sms.")
		Example("KES 0.8000")
	})
	Attribute("messageId", String, func() {
		Description("The messageId received when the sms was sent.")
		Example("ATPid_SampleTxnId123")
	})
})

var Subscriptions = Type("Subscriptions", func() {
	Attribute("id", Int, func() {
		Description("The id of the subscription")
		Example(100)
	})
	Attribute("number", String, func() {
		Description("The phone number subscribed to the product.")
		Example("+254711XXXYYY")
	})
	Attribute("Date", String, func() {
		Description("Timestamp when the subscription was made.")
		Format(FormatDateTime)
		Example("Timestamp")
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

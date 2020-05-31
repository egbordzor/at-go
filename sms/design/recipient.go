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
	// e.g KES 1.00
	Attribute("cost", String, func() {
		Description("Amount incurred to send this sms.")
		Example("KES 0.8000")
	})
	Attribute("messageId", String, func() {
		Description("The messageId received when the sms was sent.")
		Example("ATPid_SampleTxnId123")
	})
})

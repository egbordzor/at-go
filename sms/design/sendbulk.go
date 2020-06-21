package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var BulkPayload = Type("BulkPayload", func() {
	Description("Send Bulk SMS through your application")

	Attribute("apiKey", String, func() {
		Description("Africa’s Talking application apiKey")
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
		Description("Africa’s Talking application username")
	})
	Attribute("to", String, func() {
		Description("Recipients’ phone numbers")
	})
	Attribute("message", String, func() {
		Description("Message to be sent")
	})
	Attribute("from", String, func() {
		Description("Registered Short Code or Alphanumeric")
	})

	// Used by MSP to determine who gets billed for a
	// message sent using a Mobile-Terminated ShortCode.
	// This parameter will be ignored for messages sent
	// using alphanumerics or Mobile-Originated shortcodes.
	Attribute("bulkSMSMode", Int, func() {
		Description("Used by MSP to determine who gets billed for a message")

		// The default value is 1
		// This means that the sender - Africa’s Talking account
		// being used - gets charged.
		Default(1)
	})

	// Used for Bulk SMS clients that would like to
	// deliver as many messages to the API before
	// waiting for an acknowledgement from the Telcos.
	// If enabled, the API will store the messages in
	// a queue and send them out asynchronously after
	// responding to the request.
	// The default value is 1
	Attribute("enqueue", Int, func() {
		Description("Used for Bulk SMS clients")

		// 1 to enable and 0 to disable.
		Enum(0, 1)
		Default(1)
	})
	Attribute("keyword", String, func() {
		Description("The keyword to be used for a premium service")
	})
	Attribute("linkId", String, func() {
		// AT forwards the linkId to our application
		// when the user sends a message to our service.
		Description("Used for premium services to send OnDemand messages")
	})
	Attribute("retryDurationInHours", String, func() {
		// In case it’s not delivered to the subscriber
		Description("No. of hours subscription message should be retried")
	})
	Required("apiKey", "Content-Type", "username", "to", "message")
})

var BulkResponse = ResultType("BulkResponse", func() {
	Description("Send Bulk SMS Response")
	TypeName("BulkResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("SMSMessageData", func() {
			Description("A Map detailing the eventual result of the sms request")
			Attribute("Message", String, func() {
				Description("Summary of recipients sms was sent to and the total cost incurred.")
			})
			Attribute("Recipients", ArrayOf(Recipient))
		})
	})
	View("default", func() {
		Attribute("SMSMessageData")
	})
})

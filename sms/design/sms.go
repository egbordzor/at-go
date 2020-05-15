package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Live: https://api.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
// Send SMS through your application by making a HTTP POST request to the above endpoints.
var SendSMS = Type("SendSMS", func() {
	Attribute("username", String, "Africa’s Talking application username.")
	Attribute("to", String, "Recipients’ phone numbers.")
	Attribute("message", String, "Message to be sent.")
	Attribute("from", String, "Registered Short Code or Alphanumeric.")

	// This is used by the Mobile Service Provider to determine who
	// gets billed for a message sent using a Mobile-Terminated ShortCode.
	// This parameter will be ignored for messages sent using alphanumerics
	// or Mobile-Originated shortcodes.
	Attribute("bulkSMSMode", Int, func() {
		Description("Used by MSP to determine who gets billed for a message")

		// The value must be set to 1 for bulk messages.
		Enum(1)

		// The default value is 1(which means that the sender
		// - Africa’s Talking account being used - gets charged).
		Default(1)
	})

	// This is used for Bulk SMS clients that would like to deliver
	// as many messages to the API before waiting for an
	// acknowledgement from the Telcos. If enabled, the API will store
	// the messages in a queue and send them out asynchronously after responding
	// to the request. The default value is 1
	Attribute("enqueue", Int, func() {
		Description("Used for Bulk SMS clients")

		// Possible values are 1 to enable and 0 to disable.
		Enum(0, 1)
		Default(1)
	})
	Attribute("keyword", String, "The keyword to be used for a premium service.")
	// This is used for premium services to send OnDemand messages. AT forwards
	// the linkId to our application when the user sends a message to our service.
	Attribute("linkId")
	Attribute("retryDurationInHours", String, func() {
		Description("No. of hours subscription message should be retried in case it’s not delivered to the subscriber.")
	})
	Required("username", "to", "message")
})

// Live: https://api.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
// To do so, you make a HTTP GET request to the above endpoints.
// Incrementally fetches our application inbox.
var FetchMessages = Type("FetchMessages", func() {
	Attribute("username", String, "Africa’s Talking application username.")
	Attribute("lastReceivedId", String, func() {
		Description("This is the id of the message that you last processed.")

		// The default is 0.
		Default("0")
	})
})

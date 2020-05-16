package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("messaging", func() {
	Title("Messaging Service.")

	HTTP(func() {
		Path("/version1/messaging")
	})

	// Method describes a service method (endpoint).
	Method("send", func() {
		Description("Send SMS through your application")

		// Payload describes the method payload.
		Payload(SendSMSPayload)

		// Result describes the method result.
		Result(SendSMSMedia)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://api.africastalking.com/version1/messaging
		// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
		POST("/")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("apiKey", String, "Africa’s Talking application apiKey.")
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/x-www-form-urlencoded")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("Content-Type")
			})

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("fetch", func() {
		Description("incrementally Fetch your application inbox.")

		// Payload describes the method payload.
		Payload(FetchMessagesPayload)

		// Result describes the method result.
		Result(FetchMessagesMedia)

		Error("not_found")
		Error("bad_request")
		Error("internal_error")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {

			// Requests to the service consist of HTTP GET requests.
			// Live: https://api.africastalking.com/version1/messaging
			// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
			GET("/")
			Headers(func() {

				// Attribute describes an object field
				Attribute("apiKey", String, "Africa’s Talking application apiKey.")
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/json")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("Content-Type")
			})

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)

			Response("not_found", StatusNotFound, func() {
				Tag("code", "not_found")
			})
			Response("bad_request", StatusBadRequest, func() {
				Tag("code", "bad_request")
			})
			Response("internal_error", StatusInternalServerError, func() {
				Tag("code", "internal_error")
			})
		})
	})
})

// Live: https://api.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
var SendSMSPayload = Type("SendSMSPayload", func() {
	Description("Send SMS through your application.")
	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("to", String, func() {
		Description("Recipients’ phone numbers.")
	})
	Attribute("message", String, func() {
		Description("Message to be sent.")
	})
	Attribute("from", String, func() {
		Description("Registered Short Code or Alphanumeric.")
	})

	// This is used by the Mobile Service Provider
	// to determine who gets billed for a message
	// sent using a Mobile-Terminated ShortCode.
	// This parameter will be ignored for messages
	// sent using alphanumerics or Mobile-Originated
	// shortcodes.
	Attribute("bulkSMSMode", Int, func() {
		Description("Used by MSP to determine who gets billed for a message")

		// The value must be set to 1 for bulk messages.
		Enum(1)

		// The default value is 1(which means that
		// the sender - Africa’s Talking account being
		// used - gets charged).
		Default(1)
	})

	// This is used for Bulk SMS clients that would
	// like to deliver as many messages to the API
	// before waiting for an acknowledgement from
	// the Telcos. If enabled, the API will store
	// the messages in a queue and send them out
	// asynchronously after responding to the request.
	// The default value is 1
	Attribute("enqueue", Int, func() {
		Description("Used for Bulk SMS clients")
		Enum(0, 1) // 1 to enable and 0 to disable.
		Default(1)
	})
	Attribute("keyword", String, func() {
		Description("The keyword to be used for a premium service.")
	})
	// AT forwards the linkId to our application when
	// the user sends a message to our service.
	Attribute("linkId", String, func() {
		Description("Used for premium services to send OnDemand messages")
	})
	Attribute("retryDurationInHours", String, func() {
		// In case it’s not delivered to the subscriber
		Description("No. of hours subscription message should be retried.")
	})
	Required("username", "to", "message")
})

var SendSMSMedia = ResultType("SendSMSMedia", func() {
	Description("Send SMS Response ")
	TypeName("SubscribeResult")
	ContentType("application/json")

	Attribute("SMSMessageData", ArrayOf(SendSMSData), func() {
		Description("A Map detailing the eventual result of the sms request.")
	})

	View("default", func() {
		Attribute("SMSMessageData")
	})
})

// A Map detailing the eventual result of the sms request.
// // SMSMessageData
var SendSMSData = Type("SendSMSData ", func() {
	Attribute("Message", String, func() {
		Description("A summary of the total number of recipients the sms was sent to and the total cost incurred.")
		Example("Sent to 1/1 Total Cost: KES 0.8000")
	})

	// Each recipient is a Map
	Attribute("Recipients", ArrayOf(Recipients), func() {
		Description("A list of recipients included in the original request.")
	})
})

// A list of recipients included in the original request.
// Each recipient is a Map with the following fields:
var Recipients = Type("Recipients", func() {
	Description("A list of recipients included in the original request")
	Attribute("statusCode", Int, func() {
		Description("This corresponds to the status of the request")
		Enum(
			// 100: Processed
			100,

			// 101: Sent
			101,

			//102: Queued
			102,

			// 401: RiskHold
			401,

			// 402: InvalidSenderId
			402,

			// 403: InvalidPhoneNumber
			403,

			// 404: UnsupportedNumberType
			404,

			// 405: InsufficientBalance
			405,

			// 406: UserInBlacklist
			406,

			// 407: CouldNotRoute
			407,

			// 500: InternalServerError
			408,

			// 501: GatewayError
			501,

			// 502: RejectedByGateway
			502,
		)
		Example(101)
	})
	Attribute("number", String, func() {
		Description("The recipient’s phone number")
		Example("+254711XXXYYY")
	})

	// This does not indicate the delivery status of the sms to this recipient.
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

// Live: https://api.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
var FetchMessagesPayload = Type("FetchMessagesPayload", func() {
	Description("Incrementally fetches our application inbox")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("lastReceivedId", String, func() {
		Description("This is the id of the message that you last processed.")

		// The default is 0.
		Default("0")
	})
	Required("username")
})

var FetchMessagesMedia = ResultType("FetchMessagesMedia", func() {
	Description("Fetch Messages Response")
	TypeName("FetchMessagesResult")
	ContentType("application/json")

	Attribute("SMSMessageData", ArrayOf(FetchSMSData), func() {
		Description("A Map containing the messages from your inbox.")
	})

	View("default", func() {
		Attribute("SMSMessageData")
	})

})

// SMSMessageData
var FetchSMSData = Type("SendSMSData", func() {
	Description("A Map containing the messages from your inbox.")

	// Each message is a Map with the following fields
	Attribute("Messages", ArrayOf(Messages), func() {
		Description("A list of messages from your inbox.")
	})
})

var Messages = Type("Messages", func() {
	Attribute("linkId", String, func() {
		Description("A unique identifier attached to each incoming message.")
	})
	Attribute("text", String, func() {
		Description("The content of the message received.")
	})
	Attribute("to", String, func() {
		Description("Your registered short code that the sms was sent out to.")
	})
	Attribute("id", String, func() {
		Description("The id of the message.")
	})
	Attribute("date", String, func() {
		Description("The date when the sms was sent.")
	})
	Attribute("from", String, func() {
		Description("The sender’s phone number.")
	})
})

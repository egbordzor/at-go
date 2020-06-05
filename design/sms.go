package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var BulkSMSPayload = Type("BulkSMSPayload", func() {
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

var BulkSMSResponse = ResultType("BulkSMSResponse", func() {
	Description("Send Bulk SMS Response")
	TypeName("BulkSMSResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("SMSMessageData", func() {
			Description("A Map detailing the eventual result of the sms request")

			Attribute("Message", String, func() {
				Description("A summary of the total number of recipients the sms was sent to and the total cost incurred.")
				Pattern("[a-zA-Z]+")
				Example("Sent to 1/1 Total Cost: KES 0.8000")
			})
			Attribute("Recipients", ArrayOf(Recipient), func() {
				Description("A list of recipients included in the original request.")
				MinLength(1)
			})
		})
	})

	View("default", func() {
		Attribute("SMSMessageData")
	})
})

// Send Premium SMS
// Live: https://content.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
var PremiumSMSPayload = Type("PremiumSMSPayload", func() {
	Description("Send a Premium SMS through your application")

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
		// AT forwards the linkId to our application when the user sends a message to our service.
		Description("Used for premium services to send OnDemand messages")
	})
	Attribute("retryDurationInHours", String, func() {
		// In case it’s not delivered to the subscriber.
		Description("No. of hours subscription message should be retried")
	})
	Required("apiKey", "Content-Type", "username", "to", "message")
})

var PremiumSMSResponse = ResultType("PremiumSMSResponse", func() {
	Description("Send Bulk SMS Response")
	TypeName("PremiumSMSResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("SMSMessageData", func() {
			Description("A Map detailing the eventual result of the sms request")

			Attribute("Message", String, func() {
				Description("A summary of the total number of recipients the sms was sent to and the total cost incurred.")
				Pattern("[a-zA-Z]+")
				Example("Sent to 1/1 Total Cost: KES 0.8000")
			})
			Attribute("Recipients", ArrayOf(Recipient), func() {
				Description("A list of recipients included in the original request.")
				MinLength(1)
			})
		})
	})

	View("default", func() {
		Attribute("SMSMessageData")
	})
})

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

// Create Checkout token generates a checkout token
// required in order to subscribe a phone number.
// This is an open endpoint and does not need your
// authentication credentials.
// Live: https://api.africastalking.com/checkout/token/create
// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
var CreateCheckoutTokenPayload = Type("CreateCheckoutTokenPayload", func() {
	Description("Authorizes a premium SMS subscription.")

	Attribute("Content-Type", String, func() {
		Description("The requests content type")
		Default("application/x-www-form-urlencoded")
	})
	Attribute("phoneNumber", String, func() {
		Description("Mobile phone number you want to create a subscription for.")
	})
	Required("Content-Type", "phoneNumber")
})

var CreateCheckoutTokenResponse = ResultType("CreateCheckoutTokenResponse", func() {
	Description("Generated checkout token response")
	TypeName("CreateCheckoutTokenResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("description", String, func() {
			Description("A description of the status of the request.")
			Enum("Success", "Failed")
			Example("Success")
		})

		// This token shall be expected later when
		// initiating the actual create subscription request.
		// The tokens expire after 10 minutes and are limited to 2
		// tokens for a 5 minute window for each source IP Address.
		Attribute("token", String, func() {
			Description("The checkout token to be used")
			Example("CkTkn_SampleCkTknId123")
		})
	})
	View("default", func() {
		Attribute("description")
		Attribute("token")
	})
})

// Subscribe a phone number
// Live: https://api.africastalking.com/version1/subscription/create
// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create
var CreateSubPayload = Type("CreateSubPayload", func() {
	Description("subscribe a phone number payload")

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
		Description("The phoneNumber to be subscribed")
		Example("")
	})
	Attribute("checkoutToken", String, func() {
		Description("Token used to validate the subscription request")
		Example("")
	})
	Required("apiKey", "Content-Type", "username", "shortCode", "keyword", "phoneNumber", "checkoutToken")
})

var CreateSubResponse = ResultType("CreateSubResponse", func() {
	Description("subscribe a phone number response")
	TypeName("CreateSubResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Indicates whether the prompt to subscribe to shortcode was successfully raised or not.")
			Enum("Success", "Failed")
			Example("Success")
		})
		Attribute("description", String, func() {
			Description("Describes status of the create subscription request.")
			Example("Waiting for user input")
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
	})
})

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

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("sms", func() {
	Title("SMS Callback")
	Description("SMS notifications sent from Africa'sTalking gateway")

	// SMS Delivery Reports
	Method("delivery", func() {
		Description("Sent whenever the MSP confirms or rejects delivery of a message.")
		Payload(DeliveryReport)
		Result(String)
		//Error()
		HTTP(func() {
			Headers(func() {
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
			POST("/deliveryreport")
			Response(StatusOK) // 200 RFC 7231, 6.3.1
		})
	})

	// SMS Incoming messages
	Method("incoming", func() {
		Description("Sent whenever a message is sent to any of your registered shortcodes.")
		Payload(IncomingMessage)
		Result(String)
		//Error()
		HTTP(func() {
			Headers(func() {
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
			POST("/incomingmessage")
			Response(StatusOK) // 200 RFC 7231, 6.3.1
		})
	})

	// Bulk SMS Opt Out
	Method("optout", func() {
		Description("Sent whenever a user opts out of receiving messages from your alphanumeric sender ID")
		Payload(BulkSMSOptOut)
		Result(String)
		//Error()
		HTTP(func() {
			Headers(func() {
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
			POST("/bulksmsoptout")
			Response(StatusOK) // 200 RFC 7231, 6.3.1
		})
	})

	// SMS Subscription Notifications
	Method("subscription", func() {
		Description("Sent whenever someone subscribes or unsubscribes from any of your premium SMS products.")
		Payload(SubscriptionNotification)
		Result(String)
		//Error()
		HTTP(func() {
			Headers(func() {
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
			POST("/subscriptionnotification")
			Response(StatusOK) // 200 RFC 7231, 6.3.1
		})
	})
})

// 1a.SMS POST request to AT gateway
// 1b. JSON Response from AT gateway
// 2. AT gateway sends SMS to user.
// 3. Status update from Telco
// 4. Status update from AT gateway to Callback

// Delivery Report notification contents
var DeliveryReport = Type("DeliveryReport", func() {
	Description("Sent whenever the MSP confirms or rejects delivery of a message")
	Attribute("id", String, func() {
		Description("A unique identifier for each message.") // Same id as the one in the response when a message is sent
	})
	Attribute("status", String, func() {
		Description("The status of the message.")
		Enum(
			// The message has successfully been sent by our network.
			"Sent",

			// The message has successfully been submitted to the MSP (Mobile Service Provider).
			"Submitted",

			// The message has been queued by the MSP.
			"Buffered",

			// The message has been rejected by the MSP.
			// This is a final status.
			"Rejected",

			// The message has successfully been delivered to the receiver’s handset.
			// This is a final status.
			"Success",

			// The message could not be delivered to the receiver’s handset.
			// This is a final status.
			"Failed",
		)
	})
	Attribute("phoneNumber", String, func() {
		Description("This is phone number that the message was sent out to.")
	})
	Attribute("networkCode", String, func() {
		Description("A unique identifier for the telco that handled the message.")
		Enum(
			"62120", // 62120: Airtel Nigeria
			"62130", // 62130: MTN Nigeria
			"62150", // 62150: Glo Nigeria
			"62160", // 62160: Etisalat Nigeria
			"63510", // 63510: MTN Rwanda
			"63513", // 63513: Tigo Rwanda
			"63514", // 63514: Airtel Rwanda
			"63902", // 63902: Safaricom
			"63903", // 63903: Airtel Kenya
			"63907", // 63907: Orange Kenya
			"63999", // 63999: Equitel Kenya
			"64002", // 64002: Tigo Tanzania
			"64003", // 64003: Zantel Tanzania
			"64004", // 64004: Vodacom Tanzania
			"64005", // 64005: Airtel Tanzania
			"64007", // 64007: TTCL Tanzania
			"64009", // 64009: Halotel Tanzania
			"64101", // 64101: Airtel Uganda
			"64110", // 64110: MTN Uganda
			"64111", // 64111: UTL Uganda
			"64114", // 64114: Africell Uganda
			"65001", // 65001: TNM Malawi
			"65010", // 65010: Airtel Malawi
			"99999", // 99999: Athena (sandbox environment).
		)
	})
	Attribute("failureReason", String, func() {
		Description("Only provided if status is Rejected or Failed.")
		Enum(
			// This occurs when the subscriber doesn’t have enough airtime for
			// a premium subscription service/message
			"InsufficientCredit",

			// This occurs when a message is sent with an invalid linkId for an
			// onDemand service
			"InvalidLinkId",

			// This occurs when the subscriber is inactive or the account deactivated
			// by the MSP (Mobile Service Provider).
			"UserIsInactive",

			// This occurs if the user has been blacklisted not to receive messages
			// from a paricular service (shortcode or keyword)
			"UserInBlackList",

			// This occurs when the mobile subscriber has been suspended by the MSP.
			"UserAccountSuspended",

			// This occurs when the message is passed to an  MSP where the subscriber
			// doesn’t belong.
			"NotNetworkSubscriber",

			// This occurs when the message from a subscription product is sent to a
			// phone number that has not subscribed to the product.
			"UserNotSubscribedToProduct",

			// This occurs when the message is sent to a non-existent mobile number.
			"UserDoesNotExist",

			// This occurs when message delivery fails for any reason not listed above
			// or where the MSP didn’t provide a delivery failure reason.
			"DeliveryFailure",
		)
	})
	// Note: This only applies for premium SMS messages.
	Attribute("retryCount", String, func() {
		Description("Number of times the request to send a message to the device was retried before it succeeded or definitely failed.")
	})
})

// Incoming message notification contents
var IncomingMessage = Type("IncomingMessage", func() {
	Description("Sent whenever a message is sent to any of your registered shortcodes.")
	Attribute("date", String, func() {
		Description("The date and time when the message was received.")
		Format(FormatDate)
	})
	Attribute("from", String, "The number that sent the message.")
	Attribute("id", String, "The internal ID that we use to store this message.")
	Attribute("linkId", String, func() {
		Description("Parameter required when responding to an on-demand user request with a premium message.")
	})
	Attribute("text", String, "The message content.")
	Attribute("to", String, "The number to which the message was sent.")
	Attribute("networkCode", String, func() {
		Description("A unique identifier for the telco that handled the message.")
		Enum(
			"62120", // 62120: Airtel Nigeria
			"62130", // 62130: MTN Nigeria
			"62150", // 62150: Glo Nigeria
			"62160", // 62160: Etisalat Nigeria
			"63510", // 63510: MTN Rwanda
			"63513", // 63513: Tigo Rwanda
			"63514", // 63514: Airtel Rwanda
			"63902", // 63902: Safaricom
			"63903", // 63903: Airtel Kenya
			"63907", // 63907: Orange Kenya
			"63999", // 63999: Equitel Kenya
			"64002", // 64002: Tigo Tanzania
			"64003", // 64003: Zantel Tanzania
			"64004", // 64004: Vodacom Tanzania
			"64005", // 64005: Airtel Tanzania
			"64007", // 64007: TTCL Tanzania
			"64009", // 64009: Halotel Tanzania
			"64101", // 64101: Airtel Uganda
			"64110", // 64110: MTN Uganda
			"64111", // 64111: UTL Uganda
			"64114", // 64114: Africell Uganda
			"65001", // 65001: TNM Malawi
			"65010", // 65010: Airtel Malawi
			"99999", // 99999: Athena (This is a custom networkCode that only applies when working in the sandbox environment).
		)
	})
})

// The instructions on how to opt out are automatically appended to the
// first message you send to the mobile subscriber. From then onwards,
// any other message will be sent ‘as is’ to the subscriber.
var BulkSMSOptOut = Type("BulkSMSOptOut", func() {
	Description("Sent whenever a user opts out of receiving messages from your alphanumeric sender ID")
	Attribute("senderId", String, func() {
		Description("This is the shortcode/alphanumeric sender id the user opted out from.")
	})
	Attribute("phoneNumber", String, func() {
		Description("This will contain the phone number of the subscriber who opted out.")
	})
})

var SubscriptionNotification = Type("SubscriptionNotification", func() {
	Description("Sent whenever someone subscribes or unsubscribes from any of your premium SMS products.")

	Attribute("phoneNumber", String, func() {
		Description("Phone number to subscribe or unsubscribe.")
	})
	Attribute("shortCode", String, func() {
		Description("The short code that has this product.")
	})
	Attribute("keyword", String, func() {
		Description("The keyword of the product that the user has subscribed or unsubscribed from.")
	})
	Attribute("updateType", String, func() {
		Description("The type of the update.")
		Enum("addition", "deletion")
	})
})

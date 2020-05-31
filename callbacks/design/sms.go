package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// 1. Delivery Reports
//    To receive delivery reports, you need to set a delivery report callback URL.
//    From the dashboard select SMS -> SMS Callback URLs -> Delivery Reports.
//
//    Delivery Report notification contents.
var DeliveryReport = Type("DeliveryReport", func() {
	Description("Sent whenever the MSP confirms or rejects delivery of a message")

	// Same id as the one in the response when a message is sent
	Attribute("id", String, func() {
		Description("A unique identifier for each message.")
	})
	Attribute("status", String, func() {
		Description("The status of the message.")
		Enum(
			// The message has successfully been sent by our network.
			"Sent",

			// The message has successfully been submitted to the MSP
			// (Mobile Service Provider).
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
		Description("Mobile phone number message was sent out to.")
	})
	Attribute("networkCode", String, func() {
		Description("A unique identifier for the Telco that handled the message.")
		Enum(
			// Airtel Nigeria
			"62120",

			// MTN Nigeria
			"62130",

			// Glo Nigeria
			"62150",

			// Etisalat Nigeria
			"62160",

			// MTN Rwanda
			"63510",

			// Tigo Rwanda
			"63513",

			// Airtel Rwanda
			"63514",

			// Safaricom
			"63902",

			// Airtel Kenya
			"63903",

			// Orange Kenya
			"63907",

			// Equitel Kenya
			"63999",

			// Tigo Tanzania
			"64002",

			// Zantel Tanzania
			"64003",

			// Vodacom Tanzania
			"64004",

			// Airtel Tanzania
			"64005",

			// TTCL Tanzania
			"64007",

			// Halotel Tanzania
			"64009",

			// Airtel Uganda
			"64101",

			// MTN Uganda
			"64110",

			// UTL Uganda
			"64111",

			// Africell Uganda
			"64114",

			// TNM Malawi
			"65001",

			// Airtel Malawi
			"65010",

			// Athena (sandbox environment).
			"99999",
		)
	})
	Attribute("failureReason", String, func() {
		Description("Only provided if status is Rejected or Failed.")
		Enum(
			// This occurs when the subscriber doesn’t have enough
			// airtime for a premium subscription service/message.
			"InsufficientCredit",

			// This occurs when a message is sent with an invalid
			// linkId for an onDemand service.
			"InvalidLinkId",

			// This occurs when the subscriber is inactive or the
			// account deactivated by the MSP (Mobile Service Provider).
			"UserIsInactive",

			// This occurs if the user has been blacklisted not
			// to receive messages from a particular service
			// (shortcode or keyword)
			"UserInBlackList",

			// This occurs when the mobile subscriber has been
			// suspended by the MSP.
			"UserAccountSuspended",

			// This occurs when the message is passed to an MSP
			// where the subscriber doesn’t belong.
			"NotNetworkSubscriber",

			// This occurs when the message from a subscription
			// product is sent to a phone number that has not
			// subscribed to the product.
			"UserNotSubscribedToProduct",

			// This occurs when the message is sent to a non-existent
			// mobile number.
			"UserDoesNotExist",

			// This occurs when message delivery fails for any
			// reason not listed above or where the MSP didn’t
			// provide a delivery failure reason.
			"DeliveryFailure",
		)
	})

	// Note: This only applies for premium SMS messages.
	Attribute("retryCount", String, func() {
		Description("Number of times the request to send a message to the device was retried before it succeeded or definitely failed.")
	})
})

// 2. Incoming Messages.
// 	  To receive incoming messages, you need to set an incoming messages callback URL.
//    From the dashboard select SMS -> SMS Callback URLs -> Incoming Messages.
//
//    Incoming message notification contents
var IncomingMessage = Type("IncomingMessage", func() {
	Description("Sent whenever a message is sent to any of your registered shortcodes.")

	Attribute("date", String, func() {
		Description("The date and time when the message was received.")
		Format(FormatDate)
	})
	Attribute("from", String, func() {
		Description("The number that sent the message.")
	})
	Attribute("id", String, func() {
		Description("The internal ID that we use to store this message.")
	})
	Attribute("linkId", String, func() {
		Description("Parameter required when responding to an on-demand user request with a premium message.")
	})
	Attribute("text", String, func() {
		Description("The message content.")
	})
	Attribute("to", String, func() {
		Description("The number to which the message was sent.")
	})
	Attribute("networkCode", String, func() {
		Description("A unique identifier for the telco that handled the message.")
		Enum(
			// Airtel Nigeria
			"62120",

			// MTN Nigeria
			"62130",

			// Glo Nigeria
			"62150",

			// Etisalat Nigeria
			"62160",

			// MTN Rwanda
			"63510",

			// Tigo Rwanda
			"63513",

			// Airtel Rwanda
			"63514",

			// Safaricom
			"63902",

			// Airtel Kenya
			"63903",

			// Orange Kenya
			"63907",

			// Equitel Kenya
			"63999",

			// Tigo Tanzania
			"64002",

			// Zantel Tanzania
			"64003",

			// Vodacom Tanzania
			"64004",

			// Airtel Tanzania
			"64005",

			// TTCL Tanzania
			"64007",

			// Halotel Tanzania
			"64009",

			// Airtel Uganda
			"64101",

			// MTN Uganda
			"64110",

			// UTL Uganda
			"64111",

			// Africell Uganda
			"64114",

			// TNM Malawi
			"65001",

			// Airtel Malawi
			"65010",

			// Athena (sandbox environment).
			"99999",
		)
	})
})


// 3. Bulk SMS Opt Out
//    To receive bulk sms opt out notifications, you need to set a bulk sms opt out callback URL.
//    From the dashboard select SMS -> SMS Callback URLs -> Bulk SMS Opt Out.
//    The instructions on how to opt out are automatically appended to the first message
//    you send to the mobile subscriber.
//    From then onwards, any other message will be sent ‘as is’ to the subscriber.
//
//    Bulk sms opt out notification contents.
var BulkSMSOptOut = Type("BulkSMSOptOut", func() {
	Description("Triggered whenever a user opts out of receiving messages from Alphanumeric sender ID")

	Attribute("senderId", String, func() {
		Description("Shortcode/Alphanumeric Sender ID the user opted out from.")
	})
	Attribute("phoneNumber", String, func() {
		Description("Mobile phone number of the subscriber who opted out.")
	})
})


// 4. Subscription Notification.
//    To receive premium sms subscription notifications, you need to set a
//    subscription notification callback URL.
//    From the dashboard select SMS -> SMS Callback URLs -> Subscription Notifications.
//
//    Subscription notification contents
var SubscriptionNotification = Type("SubscriptionNotification", func() {
	Description("Triggered whenever someone subscribes or unsubscribes from any premium SMS product.")

	Attribute("phoneNumber", String, func() {
		Description("Mobile phone number to subscribe or unsubscribe.")
	})
	Attribute("shortCode", String, func() {
		Description("The short code that has this product.")
	})
	Attribute("keyword", String, func() {
		Description("The product keyword that the user has subscribed or unsubscribed from.")
	})
	Attribute("updateType", String, func() {
		Description("The type of the update.")
		Enum("addition", "deletion")
	})
})


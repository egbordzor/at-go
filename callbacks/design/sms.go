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

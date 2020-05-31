package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

/*
	 SMS Callback Service.
 	 This service hosts the following notifications.
     1. Delivery Reports.
     2. Incoming Messages.
     3. Bulk SMS Opt Out.
     4. Subscription Notification.
*/
var _ = Service("sms", func() {
	HTTP(func() {
		Path("/africastalking/sms/version1")
	})

	// Delivery Reports
	// To receive delivery reports, you need to set a delivery report
	// callback URL.
	// From the dashboard select SMS -> SMS Callback URLs -> Delivery Reports.
	//
	// Delivery Report notification contents
	Method("deliveryReport", func() {
		Description("Sent whenever an MSP confirms or rejects delivery of a message.")
		Payload(DeliveryReport)
		Result(String)
		HTTP(func() {
			POST("/deliveryreport")
			Response(StatusOK)
		})
	})

	// Incoming Messages
	// To receive incoming messages, you need to set an incoming messages
	// callback URL.
	// From the dashboard select SMS -> SMS Callback URLs -> Incoming Messages.
	//
	// Incoming message notification contents
	Method("incomingMessage", func() {
		Description("Sent whenever a message is sent to any of your registered shortcodes.")
		Payload(IncomingMessage)
		Result(String)
		HTTP(func() {
			POST("/incomingmessage")
			Response(StatusOK)
		})
	})

	// Bulk SMS Opt Out
	// To receive bulk sms opt out notifications, you need to set a bulk
	// sms opt out callback URL.
	// From the dashboard select SMS -> SMS Callback URLs -> Bulk SMS Opt Out.
	// The instructions on how to opt out are automatically appended to the
	// first message you send to the mobile subscriber.
	// From then onwards, any other message will be sent ‘as is’ to the subscriber.
	//
	// Bulk sms opt out notification contents
	Method("bulkSMSOptOut", func() {
		Description("Sent whenever a user opts out of receiving messages from your alphanumeric sender ID")
		Payload(BulkSMSOptOut)
		Result(String)
		HTTP(func() {
			POST("/bulksmsoptout")
			Response(StatusOK)
		})
	})

	// Subscription Notification
	// To receive premium sms subscription notifications, you need to
	// set a subscription notification callback URL.
	// From the dashboard select SMS -> SMS Callback URLs -> Subscription Notifications.
	//
	// Subscription notification contents
	Method("subNotifier", func() {
		Description("Sent whenever someone subscribes or unsubscribes from any of your premium SMS products.")
		Payload(SubscriptionNotification)
		Result(String)
		HTTP(func() {
			POST("/subscription")
			Response(StatusOK)
		})
	})
})

/*
   Voice Callback Service.
   This service hosts the following notifications.
   1. Voice notifications sent from Africa'sTalking gateway.
   2. Call Transfer Event Notifications
*/
var _ = Service("voice", func() {

	HTTP(func() {
		Path("/voice/version1")
	})

	Method("voiceNotifier", func() {
		Description("Voice Notification delivered to our callback URL")
		Payload(VoiceNotification)
		Result(String)
		HTTP(func() {
			POST("/notifications")
			Response(StatusCreated)
		})
	})

	Method("transferEvents", func() {
		Description("Event Notifications sent from AT after call transfer initiated.")
		Payload(CallTransferEvent)
		Result(String)
		HTTP(func() {
			POST("/transferevents")
			Response(StatusOK)
		})
	})
})

/*
     USSD Callback Service.
	 This service hosts the following notifications.
     1. USSD Notifications sent from Africa'sTalking gateway.
*/
var _ = Service("ussd", func() {

	HTTP(func() {
		Path("/ussd/version1")
	})

	Method("ussdNotifier", func() {
		Description("Callback URL that sends request data our App using HTTP POST.")
		Payload(USSDPayload)
		Result(USSDResponse)
		HTTP(func() {
			POST("/notification")
			Response(StatusOK)
		})
	})

})

/*
   Airtime Callback Service
   This service hosts the following notifications.
   1. Airtime Validation Notifications
   2. Airtime Status Notifications
*/
var _ = Service("airtime", func() {

	HTTP(func() {
		Path("airtime/version1")
	})

	Method("validation", func() {
		Description("Airtime Validation Notifications")
		Payload(AirtimeValidationPayload)
		Result(AirtimeValidationResponse)
		HTTP(func() {
			POST("/validation")
			Response(StatusCreated)
		})
	})

	Method("status", func() {
		Description("Airtime Status Notifications")
		Payload(AirtimeStatus)
		Result(String)
		HTTP(func() {
			POST("/status")
			Response(StatusCreated)
		})
	})
})

/*
   Payments Callback Service
   This service hosts the following notifications.
   1. Payment Notifications
   2. C2B Validation Notifications
   3. B2C Validation Notifications.
*/
var _ = Service("payments", func() {

	HTTP(func() {
		Path("/payments/version1")
	})

	Method("paymentNotifier", func() {
		Description("Payment Notifications")
		Payload(PaymentNotification)
		Result(String)
		HTTP(func() {
			POST("/events")
			Response(StatusCreated)
		})
	})

	Method("c2bNotifier", func() {
		Description("C2B Validation Notifications")
		Payload(C2BValidationNotification)
		Result(String)
		HTTP(func() {
			POST("/c2b/validation")
			Response(StatusCreated)
		})
	})

	Method("b2cNotifier", func() {
		Description("B2C Validation Notifications")
		Payload(B2CValidationNotificationPayload)
		Result(B2CValidationNotificationResponse)
		HTTP(func() {
			POST("/b2c/validation")
			Response(StatusCreated)
		})
	})
})

/*
   IoT Callback Service
   This service hosts the following notifications.
   1. IoT Notifications
*/
var _ = Service("iot", func() {

	HTTP(func() {
		Path("/iot/version1")
	})

	Method("iotNotifier", func() {
		Description("IoT Notifications")
		Payload(IoTNotification)
		Result(String)
		HTTP(func() {
			POST("/events")
			Response(StatusCreated)
		})
	})
})

var _ = Service("health", func() {
	HTTP(func() {
		Path("/health")
	})

	Method("show", func() {
		Description("Health check endpoint.")
		Result(String)
		HTTP(func() {
			GET("/")
			Response(func() {
				Code(StatusOK)
				ContentType("text/plain")
			})
		})
	})
})

var _ = Service("swagger", func() {
	Title("Open API Service")
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/swagger")
	})

	// Serve the file with relative path
	// ../../gen/http/openapi.json for requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})

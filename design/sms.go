package design

import (
	. "github.com/wondenge/atgo/sms/design"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("sms", func() {

	// Send Bulk SMS
	Method("SendBulkSMS", func() {
		Description("Send Bulk SMS")
		Payload(BulkPayload)
		Result(BulkResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/messaging
			// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
			POST("version1/messaging")
			Response(StatusCreated)
		})
	})

	// Send Premium SMS
	Method("SendPremiumSMS", func() {
		Description("Send Premium SMS")
		Payload(PremiumPayload)
		Result(PremiumResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/messaging
			// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
			POST("version1/messaging")
			Response(StatusCreated)
		})
	})

	// Fetch Messages
	Method("fetchMessages", func() {
		Description("Incrementally fetch messages from application inbox.")
		Payload(FetchMsgPayload)
		Result(FetchMsgResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/messaging
			// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
			GET("version1/messaging?{username}&{lastReceivedId}")
			Params(func() {
				Param("username:username", String, "Africa’s Talking application username.")
				Param("lastReceivedId:lastReceivedId", String, "ID of the message last processed.")
				Required("username")
			})
			Response(StatusOK)
		})
	})

	// Generate a checkout token
	Method("checkout", func() {
		Description("Generate a checkout token")
		Payload(CreateCheckoutTokenPayload)
		Result(CreateCheckoutTokenResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/checkout/token/create
			// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
			POST("checkout/token/create")
			Response(StatusCreated)
		})
	})

	// Subscribe a phone number
	Method("newSub", func() {
		Description("Subscribe a phone number")
		Payload(CreateSubPayload)
		Result(CreateSubResponse)
		HTTP(func() {

			// Live: https://content.africastalking.com/version1/subscription/create
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create
			POST("version1/subscription/create")
			Response(StatusCreated)
		})
	})

	// Incrementally fetch your premium sms subscriptions.
	Method("fetchSub", func() {
		Description("Incrementally fetch your premium sms subscriptions.")
		Payload(FetchSubPayload)
		Result(FetchSubResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/subscription
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription
			GET("/version1/subscription?username={username}&shortCode={shortCode}&keyword={keyword}&lastReceivedId={lastReceivedId}")
			Params(func() {
				Param("username:username", String, "Africa’s Talking application username.")
				Param("shortCode:shortCode", String, "Premium short code mapped to your account")
				Param("keyword:keyword", String, "Premium keyword under short code mapped to your account.")
				Param("lastReceivedId:lastReceivedId", String, "ID of the message last processed.")
				Required("username", "shortCode", "keyword")
			})
			Response(StatusOK)
		})
	})

	// Delete a Premium SMS Subscription
	Method("removeSub", func() {
		Description("Delete a Premium SMS Subscription")
		Payload(PurgeSubPayload)
		Result(PurgeSubResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/subscription/delete
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete
			POST("version1/subscription/delete")
			Response(StatusCreated)
		})
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("sms", func() {

	HTTP(func() {
		Path("/")
	})

	// Send Bulk SMS
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/version1/messaging
	// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
	Method("SendBulkSMS", func() {
		Description("Send Bulk SMS")
		Payload(BulkSMSPayload)
		Result(BulkSMSResponse)
		HTTP(func() {
			POST("version1/messaging")

			// Authenticating using an API using username and API Key
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Send Premium SMS
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/version1/messaging
	// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
	Method("SendPremiumSMS", func() {
		Description("Send Premium SMS")
		Payload(PremiumSMSPayload)
		Result(PremiumSMSResponse)
		HTTP(func() {
			POST("version1/messaging")

			// Authenticating using an API using username and API Key
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Fetch Messages
	// Requests to the service consist of HTTP GET requests.
	// Live: https://api.africastalking.com/version1/messaging
	// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
	Method("fetchMessages", func() {
		Description("Incrementally fetch messages from application inbox.")
		Payload(FetchMsgPayload)
		Result(FetchMsgResponse)
		HTTP(func() {
			GET("version1/messaging?{username}&{lastReceivedId}")
			Response(StatusOK)
			Params(func() {
				Param("username:username", String, "Africa’s Talking application username.")
				Param("lastReceivedId:lastReceivedId", String, "ID of the message last processed.")
				Required("username")
			})
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
		})
	})

	// Generate a checkout token
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/checkout/token/create
	// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
	Method("checkout", func() {
		Description("Generate a checkout token")
		Payload(CreateCheckoutTokenPayload)
		Result(CreateCheckoutTokenResponse)
		HTTP(func() {
			POST("checkout/token/create")
			Headers(func() {
				Header("Content-Type:Content-Type", String, "Content-Type")
				Required("Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Subscribe a phone number
	// Requests to the service consist of HTTP POST requests.
	// Live: https://content.africastalking.com/version1/subscription/create
	// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create
	Method("newSub", func() {
		Description("Subscribe a phone number")
		Payload(CreateSubPayload)
		Result(CreateSubResponse)
		HTTP(func() {
			POST("version1/subscription/create")
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Incrementally fetch your premium sms subscriptions.
	// Requests to the service consist of HTTP GET requests.
	// Live: https://api.africastalking.com/version1/subscription
	// Sandbox: https://api.sandbox.africastalking.com/version1/subscription
	Method("fetchSub", func() {
		Description("Incrementally fetch your premium sms subscriptions.")
		Payload(FetchSubPayload)
		Result(FetchSubResponse)
		HTTP(func() {
			GET("/version1/subscription?username={username}&shortCode={shortCode}&keyword={keyword}&lastReceivedId={lastReceivedId}")
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
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
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/version1/subscription/delete
	// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete
	Method("removeSub", func() {
		Description("Delete a Premium SMS Subscription")
		Payload(PurgeSubPayload)
		Result(PurgeSubResponse)
		HTTP(func() {
			POST("version1/subscription/delete")
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})
})

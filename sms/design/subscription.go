package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Service describes a Premium Subscription SMS Service.
var _ = Service("subscription", func() {

	Title("Premium Subscription SMS Service.")

	// Method describes a service method (endpoint)
	Method("checkout", func() {
		Description("Generate a checkout token")

		// Payload describes the method payload.
		Payload(CheckoutPayload)

		// Result describes the method result.
		Result(CheckoutResult)

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

			// Requests to the service consist of HTTP POST requests.
			// Live: https://api.africastalking.com/checkout/token/create
			// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
			POST("/checkout/token/create")

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("subscribe", func() {
		Description("Subscribe a phone number")

		// Payload describes the method payload.
		Payload(SubscriptionPayload)

		// Result describes the method result.
		Result(SubscriptionResult)

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

			// Requests to the service consist of HTTP POST requests.
			// Live: https://api.africastalking.com/checkout/token/create
			// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
			POST("/checkout/token/create")

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("fetch", func() {
		Description("Incrementally fetch your premium sms subscriptions.")

		// Payload describes the method payload.
		Payload(FetchSubscriptionPayload)

		// Result describes the method result.
		Result(FetchSubscriptionResult)

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

			// Requests to the service consist of HTTP GET requests.
			// Live: https://api.africastalking.com/version1/subscription
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription
			GET("/version1/subscription")

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("purge", func() {
		Description("Delete a Premium SMS Subscription")

		// Payload describes the method payload.
		Payload(PurgeSubscriptionPayload)

		// Result describes the method result.
		Result(PurgeSubscriptionResult)

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

			// Requests to the service consist of HTTP POST requests.
			// Live: https://api.africastalking.com/version1/subscription/delete
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete
			POST("/version1/subscription/delete")

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})
})

// Live: https://api.africastalking.com/checkout/token/create
// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
var CheckoutPayload = Type("CheckoutPayload", func() {

	// It is required in order to subscribe a phone number.
	Description("Authorizes a premium SMS subscription.")
	Attribute("phoneNumber", String, func() {
		Description("Mobile phone number you want to create a subscription for.")
	})
	Required("phoneNumber")

})

var CheckoutResult = ResultType("CheckoutResult", func() {
	Description("Generate checkout token response")
	TypeName("CheckoutResult")
	ContentType("application/json")

	Attribute("description", String, func() {
		Description("A description of the status of the request.")
		Enum("Success", "Failed")
		Example("Success")
	})

	// This token shall be expected later when initiating
	// the actual create subscription request.
	// The tokens expire after 10 minutes and are limited
	// to 2 tokens for a 5 minute window for each source IP Address.
	Attribute("token", String, func() {
		Description("The checkout token to be used")
		Example("CkTkn_SampleCkTknId123")
	})

	View("default", func() {
		Attribute("description")
		Attribute("token")
	})
})

// Live: https://api.africastalking.com/version1/subscription/create
// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create
var SubscriptionPayload = Type("SubscriptionPayload", func() {
	Description("Subscribe a phone number")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("shortCode", String, func() {
		Description("The premium short code mapped to your account.")
	})
	Attribute("keyword", String, func() {
		Description("The premium keyword under the above short code mapped to your account.")
	})
	Attribute("phoneNumber", String, func() {
		Description("The phoneNumber to be subscribed")
	})
	Attribute("checkoutToken", String, func() {
		Description("This is a token used to validate the subscription request")
	})
	Required("username", "shortCode", "keyword", "phoneNumber", "checkoutToken")

})

var SubscriptionResult = ResultType("SubscriptionResult", func() {
	Description("Premium SMS Create Subscription response")
	TypeName("SubscribeResult")
	ContentType("application/json")

	Attribute("status", String, func() {
		Description("Indicates whether the prompt to subscribe to this shortcode was successfully raised or not.")
		Enum("Success", "Failed")
		Example("Success")
	})
	Attribute("description", String, func() {
		Description("Describes the status of the create subscription request.")
		Example("Waiting for user input")
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
	})
})

// Live: https://api.africastalking.com/version1/subscription
// Sandbox: https://api.sandbox.africastalking.com/verson1/subscription
var FetchSubscriptionPayload = Type("FetchSubscriptionPayload", func() {
	Description("Incrementally fetch Premium SMS Subscriptions.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("shortCode", String, func() {
		Description("The premium short code mapped to your account.")
	})
	Attribute("keyword", String, func() {
		Description("The premium keyword under the above short code mapped to your account.")
	})
	// Set it to 0 to for the first time.
	Attribute("lastReceivedId", String, func() {
		Description("ID of the subscription you believe to be your last.")
	})
	Required("username", "shortCode", "keyword")
})

var FetchSubscriptionResult = ResultType("FetchSubscriptionResult", func() {
	Description("Fetch Subscriptions response")
	TypeName("FetchResult")
	ContentType("application/json")

	Attribute("Subscriptions", ArrayOf(FetchSubscriptionResponses))

	View("default", func() {
		Attribute("Subscriptions")
	})
})

var FetchSubscriptionResponses = Type("FetchSubscriptionResponses", func() {
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

// Live: https://api.africastalking.com/version1/subscription/delete
// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete
var PurgeSubscriptionPayload = Type("PurgeSubscriptionPayload", func() {
	Description("Delete a premium sms subscription.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("shortCode", String, func() {
		Description("The premium short code mapped to your account")
	})
	Attribute("keyword", String, func() {
		Description("The premium keyword under the above short code mapped to your account.")
	})
	Attribute("phoneNumber", String, func() {
		Description("The phoneNumber to be unsubscribed.")
	})
	Required("username", "shortCode", "keyword", "phoneNumber")
})

var PurgeSubscriptionResult = ResultType("PurgeSubscriptionResult", func() {
	Description("Delete Subscription Response")
	TypeName("PurgeSubscriptionResult ")
	ContentType("application/json")

	Attribute("status", String, func() {
		Description("Indicates whether the phone number was successfully unsubscribed or not.")
		Enum("Success", "Failed")
		Example("Success")
	})
	Attribute("description", String, func() {
		Description("Describes the status of the delete subscription request.")
		Example("Succeeded")
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
	})
})

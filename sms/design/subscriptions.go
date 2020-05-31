package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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

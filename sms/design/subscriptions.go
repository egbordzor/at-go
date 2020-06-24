package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var CheckoutTokenPayload = Type("CheckoutTokenPayload", func() {
	Description("Authorizes a premium SMS subscription.")

	Attribute("phoneNumber", String, func() {
		Description("Mobile phone number you want to create a subscription for.")
	})
	Required("phoneNumber")
})

var CheckoutTokenResponse = ResultType("CheckoutTokenResponse", func() {
	Description("Generated checkout token response")
	TypeName("CheckoutTokenResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("description", String, func() {
			Description("A description of the status of the request.")
			Enum("Success", "Failed")
			Example("Success")
		})

		// This token shall be expected later when initiating the actual create
		// subscription request.
		// The tokens expire after 10 minutes and are limited to 2 tokens for a
		// 5 minute window for each source IP Address.
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

var NewSubPayload = Type("NewSubPayload", func() {
	Description("subscribe a phone number payload")

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
	Required("username", "shortCode", "keyword", "phoneNumber", "checkoutToken")
})

var NewSubResponse = ResultType("NewSubResponse", func() {
	Description("subscribe a phone number response")
	TypeName("NewSubResponse")
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

var FetchSubPayload = Type("FetchSubPayload", func() {
	Description("Incrementally fetch Premium SMS Subscriptions.")

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
	Required("username", "shortCode", "keyword")
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

var PurgeSubPayload = Type("PurgeSubPayload", func() {
	Description("Delete a premium sms subscription.")

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
	Required("username", "shortCode", "keyword", "phoneNumber")
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

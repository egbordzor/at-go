package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var TopupStashPayload = Type("TopupStashPayload", func() {
	Description("Topup Stash HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
		Example("MyUserName")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment product initiating transaction.")
		Example("myProductName")
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code.")
		Example("KES")
	})
	Attribute("amount", Float64, func() {
		Description("Amount application will be topped up with.")
		Example(2000)
	})
	Attribute("metadata", MapOf(String, String), func() {
		Description("Metadata associated with the request.")
		Key(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
			Example("Foo")
		})
		Elem(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
			Example("Bar")
		})
	})

	Required("username",
		"productName",
		"currencyCode",
		"amount",
		"metadata",
	)
})

var TopupStashResponse = ResultType("TopupStashResponse", func() {
	Description("Topup Stash HTTP response.")
	TypeName("TopupStashResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, "Corresponds to the status of the request", func() {
			Enum(

				// The request has been accepted
				// and your application stash
				// has been topped up.
				"Success",

				// The request failed for some
				// other reason.
				// The description field will
				// contain more information.
				"Failed",
			)
			Example("Success")
		})
		Attribute("description", String, func() {
			Description("A detailed description of the request status.")
			Example("Topped up user stash. New Stash Balance: KES 1500.00")
		})
		Attribute("transactionId", String, func() {
			Description("Unique ID for successful requests.")
			Example("ATPid_SampleTxnId123")
		})

		Required("status", "description")
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
		Attribute("transactionId")
	})
})

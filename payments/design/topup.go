package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Topup Stash
// Topup stash APIs allow you to move money from a Payment Product to an Africa’s Talking application stash.
// An application stash is the wallet that funds your service usage expenses.
var TopupStashPayload = Type("TopupStashPayload", func() {
	Description("Topup Stash HTTP request.")

	Attribute("username", String, "Africa’s Talking application username.")
	Attribute("productName", String, "Africa’s Talking Payment product initiating transaction.")
	Attribute("currencyCode", String, "3-digit ISO format currency code.")
	Attribute("amount", Float64, "Amount application will be topped up with.")
	Attribute("metadata", MapOf(String, String), "Metadata associated with the request.")
	Required("username", "productName", "currencyCode", "amount", "metadata")
})

var TopupStashResponse = ResultType("TopupStashResponse", func() {
	Description("Topup Stash HTTP response.")
	TypeName("TopupStashResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, "Corresponds to the status of the request", func() {
			Enum(

				// The request has been accepted and your application stash has been topped up.
				"Success",

				// The request failed for some other reason.
				// The description field will contain more information.
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
	})
})

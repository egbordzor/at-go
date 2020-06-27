package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var WalletTransferPayload = Type("WalletTransferPayload", func() {
	Description("Wallet Transfer HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
		Example("MyUserName")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment product to initiate this transaction.")
		Example("myProductName")
	})
	Attribute("targetProductCode", Int, func() {
		Description("Unique product code to transfer the funds to.")
		Example(2373)
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code")
		Example("KES")
	})
	Attribute("amount", Float64, func() {
		Description("Amount application will be topped up with.")
		Example(2000)
	})
	Attribute("metadata", MapOf(String, String), func() {
		Description("Metadata associated with the request. ")
		Key(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
			Example("Foo")
		})
		Elem(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
			Example("Bar")
		})
	})

	Required(
		"username",
		"productName",
		"targetProductCode",
		"currencyCode",
		"amount",
		"metadata",
	)
})

var WalletTransferResponse = ResultType("WalletTransferResponse", func() {
	Description("Wallet Transfer HTTP response.")
	TypeName("WalletTransferResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Corresponds to the status of the request.")
			Enum(

				// The request has been
				// accepted and the funds
				// have been transferred
				// to the target payment product.
				"Success",

				// The request failed for
				// some other reason.
				// The description filed will
				// contain more information.
				"Failed",
			)
			Example("Success")
		})
		Attribute("description", String, func() {
			Description("A detailed description of the request status.")
			Example("Transferred funds to sandbox [TestProduct]")
		})
		Attribute("transactionId", String, func() {
			Description("A unique id that our API generates for successful requests.")
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

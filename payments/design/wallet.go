package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Wallet Transfer
// Wallet transfer APIs allow you to transfer money from one Payment Product
// to another Payment Product hosted on Africa’s Talking.
var WalletTransferPayload = Type("WalletTransferPayload", func() {
	Description("Wallet Transfer HTTP request.")

	Attribute("username", String, "Africa’s Talking application username.")
	Attribute("productName", String, "Africa’s Talking Payment product to initiate this transaction.")
	Attribute("targetProductCode", String, "Unique product code to transfer the funds to.")
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code")
	})
	Attribute("amount", Float64, "Amount application will be topped up with.")
	Attribute("metadata", MapOf(String, String), "Metadata associated with the request. ")
	Required("username", "productName", "targetProductCode", "currencyCode", "amount", "metadata")
})

var WalletTransferResponse = ResultType("WalletTransferResponse", func() {
	Description("Wallet Transfer HTTP response.")
	TypeName("WalletTransferResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Corresponds to the status of the request.")
			Enum(
				// The request has been accepted and the funds have
				// been transferred to the target payment product.
				"Success",

				// The request failed for some other reason.
				// The description filed will contain more information.
				"Failed")
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
	})
})

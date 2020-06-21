package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var MobileCheckoutPayload = Type("MobileCheckoutPayload", func() {
	Description("Mobile Checkout HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment product to initiate this transaction.")
	})
	Attribute("providerChannel", String, func() {
		Description("Provider channel the payment will be initiated from.")
	})
	Attribute("phoneNumber", String, func() {
		Description("Phone number of the client that will complete this transaction.")
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code.")
		Example("KES")
	})
	Attribute("amount", String, func() {
		Description("Amount client is expected to confirm.")
	})
	Attribute("metadata", MapOf(String, String), func() {
		Description("Map of any metadata associates with the request")
	})
	Required("username", "productName", "phoneNumber", "currencyCode", "amount")
})

var MobileCheckoutResponse = Type("MobileCheckoutResponse", func() {
	Description("Mobile Checkout HTTP response.")
	TypeName("MobileCheckoutResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Status of the request")
			Enum(

				// The request has been accepted and we are waiting
				// for the subscriber to confirm the payment.
				"PendingConfirmation",

				// The request could not be accepted as one of the
				// fields was invalid.
				// The description field will contain more information.
				"InvalidRequest",

				// Checkout to the provided phone number is not supported.
				"NotSupported",

				// The request failed for some other reason.
				// The description filed will contain more information.
				"Failed")
			Example("PendingConfirmation")
		})
		Attribute("description", String, func() {
			Description("Detailed description of the request status.")
			Example("Waiting for user input")
		})
		Attribute("transactionId", String, func() {

			// This transactionId will be sent along with the payment notification.
			Description("Unique ID that our API generates for successful requests.")
			Example("ATPid_SampleTxnId123")
		})
		Attribute("providerChannel", String, func() {
			Description("Provider channel used to initiate this transaction.")
			Example("345678")
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
		Attribute("transactionId")
		Attribute("providerChannel")
	})
})

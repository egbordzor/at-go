package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var BankCheckoutValidatePayload = Type("BankCheckoutValidatePayload", func() {
	Description("Bank Checkout Validation HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application Username.")
	})
	Attribute("transactionId", String, func() {
		Description("The ID of the transaction to be validated.")
	})
	Attribute("otp", String, func() {
		Description("One Time Password bank sent to the client.")
	})
	Required("username", "transactionId", "otp")
})

var BankCheckoutValidateResponse = ResultType("BankCheckoutValidateResponse", func() {
	Description("Bank Checkout Validation HTTP response.")
	TypeName("BankCheckoutValidateResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("The final status of this request.")
			Enum(
				// The request has been accepted and the funds have
				// been credited to your applications payment wallet.
				"Success",

				// The request failed for some other reason.
				// The description field will contain more information.
				"Failed")
			Example("Success")
		})
		Attribute("description", String, func() {
			Description("A detailed description of the request status.")
			Example("Payment completed successfully")
		})
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Card Checkout Validate
// Card checkout validation APIs allow your application to validate card checkout charge requests.
var CardCheckoutValidatePayload = Type("CardCheckoutValidatePayload", func() {
	Description("Card Checkout Validation HTTP request.")

	Attribute("username", String, "Africa’s Talking application username.")
	Attribute("transactionId", String, "ID of the transaction application wants to validate.")
	Attribute("otp", String, "One Time Password card provider sent to the client.")
	Required("username", "transactionId", "otp")
})

var CardCheckoutValidateResponse = ResultType("CardCheckoutValidateResponse", func() {
	Description("Card Checkout Validation HTTP response.")
	TypeName("CardCheckoutValidateResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Corresponds to the final status of this request.")
			Enum(

				// The request has been accepted and the funds have
				// been credited to your applications payment wallet.
				"Success",

				// The request failed for some other reason.
				// The description filed will contain more information.
				"Failed",
			)
			Example("Success")
		})
		Attribute("description", String, func() {
			Description("A detailed description of the request status.")
			Example("Payment completed successfully")
		})
		Attribute("checkoutToken", String, func() {

			//  Without having to go through the card checkout validation process.
			Description("Token application can use to initiate subsequent charges.")
			Example("ATCdTkn_SampleCdTknId123")
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
		Attribute("checkoutToken")
	})
})

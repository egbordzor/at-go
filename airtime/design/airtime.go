package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var AirtimePayload = Type("AirtimePayload", func() {
	Description("Airtime request payload")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("recipients", func() {
		Description("A url encoded json list of Recipients")

		Attribute("phoneNumber", String, func() {
			Description("Phone number that will be topped up.")
			Example("+234811222333")
		})
		Attribute("amount ", String, func() {
			Description("Value of airtime to send together with the currency code.")
			Example("KES 100.50")
		})
	})
	Required("username", "recipients")
})

var AirtimeResponse = ResultType("AirtimeResponse", func() {
	Description("An Airtime HTTP response")
	TypeName("AirtimeResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("numSent", Int, func() {
			Description("Number of requests sent to the provider")
		})

		// The format of this string is:
		// (3-digit Currency Code)(space)(Decimal Value)
		Attribute("totalAmount", String, func() {
			Description("Total value of airtime sent to the provider.")
			Example("KES 1000.0000")
		})

		// The format of this string is:
		// (3-digit Currency Code)(space)(Decimal Value)
		Attribute("totalDiscount", String, func() {
			Description("Total discount applied on the airtime.")
			Example("KES 40.0000")
		})
		Attribute("responses", ArrayOf(AirtimeEntry))
		Attribute("errorMessage", String, func() {
			Description("Error message if the ENTIRE request was rejected by the API.")
			Example("None")
		})
	})

	View("default", func() {
		Attribute("numSent")
		Attribute("totalAmount")
		Attribute("totalDiscount")
		Attribute("responses")
		Attribute("errorMessage")
	})
})

var AirtimeEntry = Type("AirtimeEntry", func() {

	Attribute("phoneNumber", String, func() {
		Description("Phone number for this transaction.")
		Example("+254711XXXYYY")
	})

	// The format of this string is:
	// (3-digit Currency Code)(space)(Decimal Value)
	Attribute("amount", String, func() {
		Description("Value of airtime requested.")
		Example("KES 1000.0000")
	})

	// The format of this string is:
	// (3-digit Currency Code)(space)(Decimal Value)
	Attribute("discount", String, func() {
		Description("Discount applied to the requested airtime amount.")
		Example("KES 40.0000")
	})

	Attribute("status", func() {
		Description("Status of the request associated to this phone number")
		Enum("Sent", "Failed")
		Example("Sent")
	})

	// This is only generated if the status of the request is Sent.
	// If the status is Failed this will have the special value None
	Attribute("requestId", String, func() {
		Description("Unique ID for the request associated to this phone number")
		Example("ATQid_1be914ac47845eef1a1dab5d89ec50ff")
	})

	// This is only generated if the status of the request is Failed.
	// If the status is Sent this will have the special value None
	Attribute("errorMessage", String, func() {
		Description("Error message for the request associated to this phone number. ")
		Example("None")
	})
})

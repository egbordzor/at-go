package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var MobileB2BPayload = Type("MobileB2BPayload", func() {
	Description("Initiate Mobile B2B HTTP request.")
	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment Product initiating transaction.")
	})
	Attribute("provider", String, func() {
		Description("Provider used to process the B2C request.")
		Enum(
			// Payments facilitated by Safaricom’s M-PESA APIs
			"Mpesa",

			// Payments facilitated by Tigo’s TigoPesa APIs
			"TigoTanzania",

			// Payments facilitated by our Developer Sandbox.
			// Please note: This is not available on our production systems
			"Athena")
	})
	Attribute("transferType", String, func() {
		Description("Transfer type of the payment.")
		Enum("BusinessBuyGoods", "BusinessPayBill", "DisburseFundsToBusiness", "BusinessToBusinessTransfer")
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code")
		Example("KES")
	})
	Attribute("amount", String, func() {
		Description("Amount client is expected to confirm.")
	})
	Attribute("destinationChannel", String, func() {
		Description("Name or number of the channel receiving payment by the provider.")
	})
	Attribute("destinationAccount", String, func() {
		Description("Account name used by the business to receive money on the provided destinationChannel.")
	})
	Attribute("metadata", String, func() {
		Description("A map of any metadata associated with the request.")
	})
})

var MobileB2BResponse = ResultType("MobileB2BResponse", func() {
	Description("Mobile B2B HTTP response.")
	TypeName("MobileB2BResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("The status of the B2B transaction.")
			Enum(
				// The transaction has been accepted and
				// queued for processing by the payment provider.
				"Queued",

				// We could not accept the request because
				// one of the fields was invalid.
				// The errorMessage field will contain a
				// detailed description of the requests status.
				"InvalidRequest",

				// B2B requests to the provided phone number
				// is not supported.
				"NotSupported",

				// The B2B request failed for some other reason.
				// The errorMessage field will contain a detailed
				// description of the requests status.
				"Failed")
			Example("Queued")
		})
		Attribute("transactionId", String, func() {
			Description("A unique id that our API generates for successful requests.")
			Example("ATPid_SampleTxnId123")
		})
		Attribute("transactionFee", String, func() {
			Description("Transaction fee charged by Africa’s Talking for this transaction.")
			Example("KES 1.00")
		})
		Attribute("providerChannel", String, func() {
			Description("Provider channel which facilitated the payment.")
			Example("12345")
		})
		Attribute("errorMessage", String, func() {
			Description("A more descriptive error message for the status of this transaction.")
		})
		Required("status")
	})
})

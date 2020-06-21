package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var MobileB2CPayload = Type("MobileB2CPayload", func() {
	Description("Initiate Mobile B2C HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment Product to initiate this transaction.")
	})

	// Please note: You can currently pass a maximum of 10 recipient elements.
	// Recipient is a Map of a mobile subscribers account that will receive the B2C transaction.
	Attribute("recipients", func() {
		Description("List of Recipient elements of a B2C transaction request.")

		Attribute("name", String, func() {
			Description("Name of the B2C transaction recipient.")
		})
		Attribute("phoneNumber", String, func() {
			Description("Phone number of the B2C transaction recipient.")
		})
		Attribute("currencyCode", String, func() {
			Description("3-digit ISO format currency code.")
		})
		Attribute("amount", String, func() {
			Description("Amount that the client is expected to confirm.")
		})

		// The payment channel must be mapped to your Africa’s Talking payment product.
		// If not specified, a default provider channel will be used.
		Attribute("providerChannel", String, func() {
			Description("Channel payment will be made from.")
		})
		Attribute("reason", String, func() {
			Description("Purpose of the payment.")
			Enum("SalaryPayment",
				"SalaryPaymentWithWithdrawalChargePaid",
				"BusinessPayment",
				"BusinessPaymentWithWithdrawalChargePaid",
				"PromotionPayment",
			)
		})

		// Use this field to send data that will map notifications to B2C requests.
		// It will be included in the notification we send once the B2C request is complete.
		Attribute("metadata", MapOf(String, String), func() {
			Description("Map of metadata associated with the request.")
		})
		Required("phoneNumber", "currencyCode", "amount")
	})
	Required("username", "productName", "recipients")
})

var MobileB2CResponse = Type("MobileB2CResponse", func() {
	Description("Mobile B2C HTTP response.")
	TypeName("MobileB2CResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("numQueued", Int, func() {
			Description("Number of B2C transactions that were successfully queued.")
			Example(1)
		})
		Attribute("totalValue", String, func() {
			Description("Total value of all the transactions that were successfully queued.")
			Example("KES 100000")
		})
		Attribute("totalTransactionFee", String, func() {
			Description("Total transaction fee charged for all the transactions that were successfully queued. ")
			Example("KES 2.00")
		})
		Attribute("entries", ArrayOf(B2CEntry), func() {
			Description("A list of B2C entries")
		})
		Attribute("errorMessage", String, func() {
			Description("Error message if the ENTIRE request was rejected by the API")
		})
	})
})

var B2CEntry = Type("B2CEntry", func() {
	Attribute("phoneNumber", String, func() {
		Description("The phone number of the B2C transaction recipient.")
		Example("+254711XXXYYY")
	})
	Attribute("status", String, func() {
		Description("The status of the B2C transaction.")
		Enum(

			// The transaction has been accepted and queued
			// for processing by the payment provider.
			"Queued",

			// We could not accept the request because one
			// of the fields was invalid.
			// The errorMessage field will contain a detailed
			// description of the requests status.
			"InvalidRequest",

			// B2C requests to the provided phone number
			// is not supported.
			"NotSupported",

			// The request failed for some other reason.
			// The errorMessage field will contain a detailed
			// description of the request status.
			"Failed",
		)
		Example("Queued")
	})
	Attribute("transactionId", String, func() {
		Description("Unique ID that our API generates for successful requests.")
		Example("ATPid_SampleTxnId123")
	})
	Attribute("provider", String, func() {
		Description("Provider used to process the B2C request.")
		Enum(
			// For all Safaricom phone numbers.
			"Mpesa",

			//  For ALL mobile subscribers in Uganda.
			"Segovia",

			// For ALL numbers when using the Sandbox environment.
			"Athena")
	})
	Attribute("providerChannel", String, func() {
		Description("Channel used to process the B2C request.")
		Example("525900")
	})
	Attribute("value", String, func() {
		Description("Value sent to the mobile subscriber.")
		Example("KES 100000")
	})
	Attribute("transactionFee", String, func() {
		Description("Transaction fee charged by Africa’s Talking for this transaction.")
		Example("KES 1.00")
	})
	Attribute("errorMessage", String, func() {
		Description("A more descriptive error message for the status of this transaction.")
		Example("Insufficient Credit")
	})
	Required("phoneNumber", "status")
})

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
		Example("MyAppUserName")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment product to initiate this transaction.")
		Example("myProductName")
	})
	Attribute("providerChannel", String, func() {
		Description("Provider channel the payment will be initiated from.")
		Example("myProviderChannel")
	})
	Attribute("phoneNumber", String, func() {
		Description("Phone number of the client that will complete this transaction.")
		Example("+254711XXXYYY")
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code.")
		Example("KES")
	})
	Attribute("amount", Float64, func() {
		Description("Amount client is expected to confirm.")
		Example(3000)
	})
	Attribute("metadata", MapOf(String, String), func() {
		Description("Map of any metadata associates with the request")
	})
	Required("username", "productName", "phoneNumber", "currencyCode", "amount")
})

var MobileCheckoutResponse = ResultType("MobileCheckoutResponse", func() {
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

var MobileB2CPayload = Type("MobileB2CPayload", func() {
	Description("Initiate Mobile B2C HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
		Example("MyAppUserName")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment Product to initiate this transaction.")
		Example("myProductName")
	})

	// Please note: You can currently pass a maximum of 10 recipient elements.
	// Recipient is a Map of a mobile subscribers account that will receive the B2C transaction.
	Attribute("recipients", ArrayOf(MobileRecipients), func() {
		Description("A list of B2C Mobile Recipients")
		MinLength(1)
		MaxLength(10)
	})
	Required("username", "productName", "recipients")
})

var MobileB2CResponse = ResultType("MobileB2CResponse", func() {
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

var MobileB2BPayload = Type("MobileB2BPayload", func() {
	Description("Initiate Mobile B2B HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
		Example("MyAppUserName")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment Product initiating transaction.")
		Example("myProductName")
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
		Example("Mpesa")
	})

	Attribute("transferType", String, func() {
		Description("Transfer type of the payment.")
		Enum("BusinessBuyGoods", "BusinessPayBill", "DisburseFundsToBusiness", "BusinessToBusinessTransfer")
		Example("BusinessBuyGoods")
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code")
		Example("KES")
	})
	Attribute("amount", Float64, func() {
		Description("Amount client is expected to confirm.")
		Example(3000)
	})
	Attribute("destinationChannel", String, func() {

		// E.g the Mobile Provider’s PayBill or Buy Goods number that belongs to your organization.
		Description("Name or number of the channel receiving payment by the provider.")
		Example("Buy Goods Number")
	})
	Attribute("destinationAccount", String, func() {
		Description("Account name used by the business to receive money on the provided destinationChannel.")
		Example("Account Name")
	})
	Attribute("metadata", MapOf(String, String), func() {
		Description("A map of any metadata associated with the request.")
		Key(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
		})
		Elem(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
		})
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
			Example("error")
		})
		Required("status")
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

var MobileRecipients = Type("MobileRecipients", func() {
	Description("List of Recipient elements of a B2C transaction request.")

	Attribute("name", String, func() {
		Description("Name of the B2C transaction recipient.")
		Example("name")
	})
	Attribute("phoneNumber", String, func() {
		Description("Phone number of the B2C transaction recipient.")
		Example("+254711XXXYYY")
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code.")
		Example("KES")
	})
	Attribute("amount", Float64, func() {
		Description("Amount that the client is expected to confirm.")
		Example(100.50)
	})

	// The payment channel must be mapped to your Africa’s Talking payment product.
	// If not specified, a default provider channel will be used.
	Attribute("providerChannel", String, func() {
		Description("Channel payment will be made from.")
		Example("myProviderChannel")
	})
	Attribute("reason", String, func() {
		Description("Purpose of the payment.")
		Enum(
			"SalaryPayment",
			"SalaryPaymentWithWithdrawalChargePaid",
			"BusinessPayment",
			"BusinessPaymentWithWithdrawalChargePaid",
			"PromotionPayment",
		)
		Example("SalaryPayment")
	})

	// Use this field to send data that will map notifications to B2C requests.
	// It will be included in the notification we send once the B2C request is complete.
	Attribute("metadata", MapOf(String, String), func() {
		Description("Map of metadata associated with the request.")
		Key(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
		})
		Elem(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
		})
	})
	Required("phoneNumber", "currencyCode", "amount")
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Supported Banks:
// FCMB Nigeria	    234001
// Zenith Nigeria	234002
// Access Nigeria	234003
// GTBank Nigeria	234004
// Ecobank Nigeria	234005
// Diamond Nigeria	234006
// Providus Nigeria	234007
// Unity Nigeria	234008
// Stanbic Nigeria	234009
// Sterling Nigeria	234010
// Parkway Nigeria	234011
// Afribank Nigeria	234012
// Enterprise Nigeria	234013
// Fidelity Nigeria	    234014
// Heritage Nigeria	    234015
// Keystone Nigeria  	234016
// Skye Nigeria      	234017
// Stanchart Nigeria	234018
// Union Nigeria	    234019
// UBA Nigeria	        234020
// Wema Nigeria	        234021
// First Nigeria	    234022
var BankTransferPayload = Type("BankTransferPayload", func() {
	Description("Bank Transfer HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment product to initiate this transaction.")
	})
	Attribute("recipients", func() {
		Description("A list of Recipient elements each corresponding to a bank transfer transaction request. ")

		Attribute("bankAccount", String, func() {
			Description("Details of a bank account to receive the bank transfer payment.")
		})
		Attribute("accountName", String, func() {
			Description(" Bank account name.")
		})
		Attribute("accountNumber", String, func() {
			Description("Bank account number.")
		})
		Attribute("bankCode", String, func() {
			Description("6-Digit Integer code for the bank that we allocate.")
		})
		Attribute("dateOfBirth", String, func() {
			Description("Date of birth of the account owner.") // Required for Zenith Nigeria.
		})
		Attribute("currencyCode", String, func() {
			Description(" 3-digit ISO format currency code")
		})
		Attribute("amount", String, func() {
			Description("Amount client is expected to receive.")
		})
		Attribute("narration", String, func() {
			Description("Short description of the transaction displayed on the clients statement.")
		})
		Attribute("metadata", MapOf(String, String), func() {
			Description("A map of any metadata associated with the request.")
		})
		Required("bankAccount", "accountNumber", "bankCode", "currencyCode", "amount")
	})
	Required("username", "productName", "recipients")
})

var BankTransferResponse = ResultType("BankTransferResponse", func() {
	Description("Bank Transfer HTTP response.")
	TypeName("BankTransferResponse")
	ContentType("application/json")

	Attributes(func() {

		// Each corresponds to a recipient object in the bank transfer transaction request.
		// Entry is a Map containing details of a bank transfer transaction result.
		Attribute("entries", func() {
			Description("A list of bank transfer entries.")

			Attribute("accountNumber", String, func() {
				Description("The account number of the bank transfer recipient.")
				Example("1234567890")
			})
			Attribute("status", String, func() {
				Description("The transaction has been accepted and queued for processing by the payment provider.")
				Enum(

					// We could not accept the request because one of the fields was invalid.
					// The errorMessage field will contain a detailed description of the request
					"InvalidRequest",

					// Bank transfer requests to the provided bank account is not supported.
					"NotSupported",

					// The request failed for some other reason.
					// The errorMessage field will contain a detailed description of the request status.
					"Failed")
				Example("Queued")
			})
			Attribute("transactionId", String, func() {
				Description("A unique ID that our API generates for successful requests.")
				Example("ATPid_SampleTxnId")
			})

			// Please note: The transaction fee will be deducted from
			// your Africa’s Talking credits NOT your payment wallet.
			// The format of this string is:
			// (3-digit Currency Code)(space)(Decimal Value)
			Attribute("transactionFee", String, func() {
				Description("Transaction fee charged by Africa’s Talking for this transaction. ")
				Example("NGN 50.00")
			})
			Attribute("errorMessage", String, func() {
				Description("A more descriptive error message for the status of this transaction.")
				Example("Insufficient Credit")
			})
		})
		Attribute("errorMessage", String, func() {
			Description("Error message if the ENTIRE request was rejected by the API.")
			Example("Insufficient Credit")
		})
	})
})

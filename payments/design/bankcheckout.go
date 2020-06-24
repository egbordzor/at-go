package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Supported Banks:
// FCMB Nigeria	234001
// Zenith Nigeria	234002
// Access Nigeria	234003
// Providus Nigeria	234007
// Sterling Nigeria	234010
var BankCheckoutPayload = Type("BankCheckoutPayload", func() {
	Description("Bank Checkout HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment Product to initiate this transaction.")
	})
	Attribute("bankAccount", func() {
		Attribute("accountName", String, "Bank account name.")
		Attribute("accountNumber", String, "Bank account number.")
		Attribute("bankCode", Int, "6-Digit Integer code for the bank that we allocate.")
		Attribute("dateOfBirth", String, func() {
			Description("Date of birth of the account owner.") // Required for Zenith Nigeria.
			Format(FormatDate)
		})
		Required("accountName", "accountNumber", "bankCode", "dateOfBirth")
	})
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code.")
		Example("KES")
	})
	Attribute("amount", Float64, func() {
		Description("Amount client is expected to confirm.")
	})
	Attribute("narration", String, func() {
		Description("Short description of the transaction displayed on the clients statement.")
	})

	// Use this field to send data that will map notifications
	// to bank checkout charge requests.
	// It will be included in the notification we send once the
	// client completes the bank checkout request.
	Attribute("metadata", MapOf(String, String), func() {
		Description("A map of any metadata that you would like us to associate with the request.")
	})
	Required("username", "productName", "bankAccount", "currencyCode", "amount", "narration")
})

var BankCheckoutResponse = ResultType("BankCheckoutResponse", func() {
	Description("Bank Checkout HTTP response.")
	TypeName("BankCheckoutResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("This corresponds to the status of the request.")
			Enum(

				//  The request has been accepted.
				//  Your application should provide the user a way to
				//  enter the OTP sent to them by the bank and finally
				//  send that to us using a bank checkout validate request
				"PendingConfirmation",


				// The request could not be accepted as one of the fields was invalid.
				// The description field will contain more information.
				"InvalidRequest",

				// Checkout to the provided bank account is not supported.
				"NotSupported",

				// The request failed for some other reason.
				// The description filed will contain more information.
				"Failed",
			)
			Example("PendingValidation")
		})
		Attribute("description", String, func() {
			Description("A detailed description of the request status.")
			Example("Waiting for user input")
		})
		Attribute("transactionId", String, func() {
			Description("Unique ID that our API generates for successful requests.")
			Example("ATPid_SampleTxnId123")
		})
	})
})

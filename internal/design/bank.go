package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var BankCheckoutPayload = Type("BankCheckoutPayload", func() {
	Description("Bank Checkout HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
		Example("MyAppUserName")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment Product to initiate this transaction.")
		Example("myPaymentProductName")
	})
	Attribute("bankAccount", BankAccount)
	Attribute("currencyCode", String, func() {
		Description("3-digit ISO format currency code.")
		Example("NGN")
	})
	Attribute("amount", Float64, func() {
		Description("Amount client is expected to confirm.")
		Example(39393)
	})
	Attribute("narration", String, func() {
		Description("Short description of the transaction displayed on the clients statement.")
		Example("payment for airtime")
	})

	// Use this field to send data that will map notifications
	// to bank checkout charge requests.
	// It will be included in the notification we send once the
	// client completes the bank checkout request.
	Attribute("metadata", MapOf(String, String), func() {
		Description("A map of any metadata that you would like us to associate with the request.")
		Key(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
		})
		Elem(func() {
			Pattern("[a-zA-Z]+") // Validates values of the map
		})
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

var BankTransferPayload = Type("BankTransferPayload", func() {
	Description("Bank Transfer HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("productName", String, func() {
		Description("Africa’s Talking Payment product to initiate this transaction.")
	})
	Attribute("recipients", ArrayOf(TransferRecipients), func() {
		Description("Transfer Recipients")
		MinLength(1)
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
		Attribute("entries", ArrayOf(TransferEntries), func() {
			Description("Transfer Entries")
			MinLength(1)

		})
		Attribute("errorMessage", String, func() {
			Description("Error message if the ENTIRE request was rejected by the API.")
			Example("Insufficient Credit")
		})
	})

	View("default", func() {
		Attribute("entries")
		Attribute("errorMessage")
	})
})

var TransferRecipients = Type("TransferRecipients", func() {
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

var TransferEntries = Type("TransferEntries", func() {
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

var BankAccount = Type("BankAccount", func() {

	Attribute("accountName", String, func() {
		Description("Bank account name.")
		Example("Test Bank Account")
	})
	Attribute("accountNumber", String, func() {
		Description("Bank account number.")
		Example("1234567890")
	})
	Attribute("bankCode", Int, func() {
		Description("6-Digit Integer code for the bank that we allocate.")
		Example(234001)
	})
	Attribute("dateOfBirth", String, func() {

		// Required for Zenith Nigeria.
		Description("Date of birth of the account owner.")
		Format(FormatDate)
	})

	Required("accountName", "accountNumber", "bankCode", "dateOfBirth")
})

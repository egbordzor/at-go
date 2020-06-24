package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Find Transaction
// Live: https://payments.africastalking.com/query/transaction/find
// Sandbox: https://payments.sandbox.africastalking.com/query/transaction/find
var FindTransactionPayload = Type("FindTransactionPayload", func() {
	Description("Find Transaction Query HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("transactionId", String, func() {
		Description("ID of the transaction you would like to find.")
	})
	Required("username", "transactionId")
})

var FindTransactionResponse = ResultType("FindTransactionResponse", func() {
	Description("Find Transaction Query HTTP response.")
	TypeName("FindTransactionResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Status of the request")
			Enum("Success", "Failed")
		})
		Attribute("data", TransactionResponse, func() {
			Description("Details of the transaction.")
		})
		Attribute("errorMessage", String, func() {
			Description("A message detailing what happened with a failed request.")
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("data")
		Attribute("errorMessage")
	})
})

// Fetch Product Transactions
var ProductTransactionsPayload = Type("ProductTransactionsPayload", func() {
	Description("Fetch Product Transaction Query HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("productName", String, func() {
		Description("Name of the payment product to fetch.")
	})
	Attribute("pageNumber", String, func() {
		Description("Number of the page you’d like to read results from.")
	})
	Attribute("count", String, func() {
		Description("Number of transaction results you would like for this query.")
	})
	Attribute("startDate", String, func() {
		Description("Transaction start date; in the format YYYY-MM-DD")
		Format(FormatDate)
	})
	Attribute("endDate", String, func() {
		Description("Transaction end date; in the format YYYY-MM-DD")
	})
	Attribute("category", String, func() {
		Description("Transaction category you would like to consider.")
		Enum("BankCheckout", "CardCheckout", "MobileCheckout", "MobileC2B", "MobileB2C", "MobileB2B", "BankTransfer", "WalletTransfer", "UserStashTopup")
	})
	Attribute("provider", String, func() {
		Description("Transaction provider you would like to consider.")
		Enum("Mpesa", "Segovia", "Flutterwave", "Admin", "Athena")
	})
	Attribute("status", String, func() {
		Description("Transaction status you would like to consider")
		Enum("Success", "Failed")
	})
	Attribute("source", String, func() {
		Description("Transaction source you would like to consider.")
		Enum("phoneNumber", "BankAccount", "Card", "Wallet")
	})
	Attribute("destination", String, func() {
		Description("Transaction destination you would like to consider.")
		Enum("PhoneNumber", "BankAccount", "Card", "Wallet")
	})

	// This could, for example, be the Mobile Provider’s Paybill
	// or Buy Goods number that belongs to your organization.
	Attribute("providerChannel", String, func() {
		Description("Transaction provider channel you would like to consider.")
	})
	Required("username", "productName", "pageNumber", "count")
})

var ProductTransactionsResponse = ResultType("ProductTransactionsResponse", func() {
	Description("Fetch Product Transaction Query HTTP response.")
	TypeName("ProductTransactionsResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("The status of the request.")
			Enum(
				"Success", "Failed")
		})

		// A list of response entries containing details of transactions that were found.
		Attribute("responses", ArrayOf(TransactionResponse))
	})

	View("default", func() {
		Attribute("status")
		Attribute("responses")
	})
})

// Fetch Wallet Transactions
var WalletTransactionsPayload = Type("WalletTransactionsPayload", func() {
	Description("Wallet Transaction HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
	})
	Attribute("pageNumber", String, func() {
		Description("Number of the page you’d like to read results from.")
	})
	Attribute("count", String, func() {
		Description("Number of transaction results you would like for this query.")
	})
	Attribute("startDate", String, func() {
		Description("Transaction start date; in the format YYYY-MM-DD")
		Format(FormatDate)
	})
	Attribute("endDate", String, func() {
		Description("Transaction end date; in the format YYYY-MM-DD")
	})
	Attribute("categories", String, func() {
		Description("List of transaction categories you would like to consider.")
		Enum("BankCheckout", "CardCheckout", "MobileCheckout", "MobileC2B", "MobileB2C", "MobileB2B", "BankTransfer", "WalletTransfer", "UserStashTopup")
	})
	Required("username", "pageNumber", "count")

})

var WalletTransactionsResponse = ResultType("WalletTransactionsResponse", func() {
	Description("Wallet Transaction HTTP response.")
	TypeName("WalletTransactionsResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("The status of the request")
			Enum("Success", "Failed")
			Example("Success")
		})
		Attribute("responses", ArrayOf(WalletEntry), func() {
			Description("List of response Entry corresponding to a transaction result.")
		})
		Attribute("errorMessage", String, func() {
			Description("A message detailing what happened with a failed request.")
		})
	})
	View("default", func() {
		Attribute("status")
		Attribute("responses")
		Attribute("errorMessage")
	})
})

// Fetch Wallet Balance
var WalletBalancePayload = Type("WalletBalancePayload", func() {
	Description("Wallet Balance HTTP request.")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username")
	})
})

var WalletBalanceResponse = ResultType("WalletBalanceResponse", func() {
	Description("Wallet Balance HTTP response.")
	TypeName("WalletBalanceResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("The status of the request.")
			Enum("Success", "Failed")
			Example("Success")
		})
		Attribute("balance", String, func() {
			Description("Balance of the payment wallet. ")
			Example("KES 47.7191")
		})
		Attribute("errorMessage", String, func() {
			Description("A message detailing what happened with a failed request.")
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("balance")
		Attribute("errorMessage")
	})
})

var WalletEntry = Type("WalletEntry", func() {
	Attribute("description", String, func() {
		Description("Detailed description of this transaction")
		Example("MobileB2C Payment Request to +254708663158")
	})
	Attribute("balance", String, func() {
		Description("The remaining wallet balance after the transaction was processed.")
		Example("KES 47.7191")
	})
	Attribute("category", String, func() {
		Description("Category of the payment")
		Enum("BankCheckout", "CardCheckout", "MobileCheckout", "MobileC2B", "MobileB2C", "MobileB2B", "BankTransfer", "WalletTransfer", "UserStashTopup")
	})
	Attribute("transactionData", CollectionOf(FindTransactionResponse), func() {
		Description("Contains details of the specific transaction")
	})
	Attribute("value", String, func() {
		Description("Value being exchanged in this transaction.")
		Example("KES 2900.0000")
	})
	Attribute("transactionId", String, func() {
		Description("A unique transactionId generated for every payment sent and received")
		Example("ATPid_b9379b671fee8ccf24b2c74f94da0ceb")
	})
})

var TransactionResponse = Type("TransactionResponse", func() {
	Attribute("requestMetadata", MapOf(String, String), func() {
		Description("Metadata sent by your application when it initiated this transaction.")
	})
	Attribute("sourceType", String, func() {
		Description("Type of party providing funds for this transaction (the Debit Party).")
		Enum(

			// Indicates that the funds are being provided by
			// a mobile subscriber through their mobile device.
			"PhoneNumber",

			// Indicates that the funds are being provided by
			// a customer through their Bank Account.
			"BankAccount",

			// Indicates that the funds are being provided by
			// a customer through their Debit or Credit Card.
			"Card",

			// Indicates that the funds are being provided by
			// your Africa’s Talking Wallet through one of your products.
			"Wallet",
		)
		Example("Wallet")
	})

	// This value will contain either the phone number, bank account number
	// or a card number of the customer who is sending funds to your application,
	// or the special value PaymentWallet that identifies your Africa’s Talking Payment Wallet.
	Attribute("source", String, func() {
		Description("Uniquely identifies the party providing the funds for this transaction")
		Example("PaymentWallet")
	})
	Attribute("provider", String, func() {
		Description("Payment provider that facilitated this transaction")
		Enum(

			// This identifies payments facilitated by Safaricom’s M-PESA’s APIs.
			"Mpesa",

			// This identifies payments facilitated by Segovia’s APIs.
			"Segovia",

			// This identifies payments facilitated by Flutterwaves’s APIs.
			"Flutterwave",

			// This identifies payments facilitated by our administrative APIs.
			"Admin",

			// This identifies payments facilitated by our Developer Sandbox.
			// This is not available on our production systems.
			"Athena",
		)
		Example("Mpesa")
	})
	Attribute("destinationType", String, func() {
		Description("Identifies party receiving funds in this transaction (the Credit Party)")
		Enum(

			// Indicates that the funds are being
			// sent to a mobile subscriber.
			"phoneNumber",

			// Indicates that the funds are being
			// sent to a customers Bank Account.
			"BankAccount",

			// Indicates that the funds are being sent
			// to a customers Debit or Credit Card.
			"Card",

			// Indicates that the funds are being sent
			// to your Africa’s Talking Wallet through
			// one of your products.
			"Wallet",
		)
		Example("PhoneNumber")
	})
	Attribute("description", String, func() {

		//  including a more detailed failure reason in the case of failures
		Description("Contains a detailed description of this transaction.")
		Example("The service request is processed successfully.")
	})
	Attribute("providerChannel", String, func() {
		Description("Name or number of the channel used to facilitate this payment by the provider.")
		Example("824879")
	})

	// Please note: The transaction fee will be deducted from
	// your Africa’s Talking Stash NOT your payment wallet.
	// The format of this string is:
	// (3-digit Currency Code)(space)(Decimal Value)
	// e.g KES 1.50.
	Attribute("transactionFee", String, func() {
		Description("Transaction fee charged by Africa’s Talking for this transaction.")
		Example("KES 1.0000")
	})
	Attribute("providerRefId", String, func() {

		// (e.g the M-PESA transactionId)
		Description("Unique ID generated by the payment provider for this transaction.")
		Example("SAMPLE_MPESA_CODE")
	})
	Attribute("providerMetadata", MapOf(String, String), func() {
		Description("Map of any additional data received from a payment provider.")
	})
	Attribute("status", String, func() {
		Description("Final status of this transaction.")
		Enum("Success", "Failed")
		Example("Success")
	})
	Attribute("productName", String, func() {
		Description("Identifies the Africa’s Talking Payment Product used.")
		Example("testing")
	})
	Attribute("category", String, func() {
		Description("Category of the payment")
		Enum(
			// For Consumer to Business payments initiated by
			// your application through our Bank Checkout APIs
			"BankCheckout",

			// For Consumer to Business payments initiated by
			// your application through our Card Checkout APIs
			"CardCheckout",

			// For Consumer to Business payments initiated by
			// your application through our Mobile Checkout APIs
			"MobileCheckout",

			// For Consumer to Business payments initiated by
			// a mobile subscriber through their device
			// (e.g using a paybill number)
			"MobileC2B",

			// For Business to Consumer payments initiated by
			// your application through our B2C APIs
			"MobileB2C",

			// For Business to Business payments initiated by
			// your application through our B2B APIs
			"MobileB2B",

			// For Business to Business payments initiated by
			// your application through our Bank Transfer APIs
			"BankTransfer",

			// For Wallet to Wallet payments initiated by your
			// application through our Wallet Transfer APIs
			"WalletTransfer",

			// For Wallet to application stash payments initiated by
			// your application through our User Stash Topup APIs
			"UserStashTopup",
		)
		Example("MobileB2C")
	})
	Attribute("transactionDate", String, func() {

		// This is only provided for successful transactions
		Description("Date and time when a successful transaction was completed.")
		Format(FormatDateTime)
		Example("12.05.2018 21:46:13")
	})

	//  This value will contain either a phone umber, bank account number
	//  or a card number of the customer who is sending funds to your application,
	//  or the special value PaymentWallet that identifies your Africa’s Talking Payment Wallet.
	Attribute("destination", String, func() {
		Description("Uniquely identifies the party receiving the funds for this transaction.")
		Example("+254708663158")
	})

	// he format of this string is: (3-digit Currency Code)(space)(Decimal Value) e.g KES 1.50
	Attribute("value", String, func() {
		Description("Value being exchanged in this transaction.")
		Example("KES 2900.0000")
	})
	Attribute("transactionId", String, func() {
		Description("Unique transactionId generated for every payment sent and received.")
		Example("ATPid_b9379b671fee8ccf24b2c74f94da0ceb")
	})
	Attribute("creationTime", String, func() {
		Description("Date and time when a transaction was accepted by our APIs.")
		Format(FormatDateTime)
		Example("2018-05-12 18:46:12")
	})
})

package design

import (
	. "github.com/wondenge/atgo/payments/design"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Africa's Talking have built simple REST APIs that allow your
// application to charge and send payments to more than 300 million
// mobile, bank and card account holders across Africa.
// You can quickly integrate and manage payments in your application.
// Africa's Talking also provides an online wallet that allows you to
// seamlessly receive payments from payment providers
// (such as Consumer-to-Business mobile payments), as well as initiate
// payments going out to your customers (such as Business-to-Consumer mobile payments).
// Africa's Talking also provide a robust reconciliation back-end helping
// you easily settle wallet balances into your mobile money wallet or bank account.
// These payments APIs can be used to for users in Kenya, Nigeria and Uganda.
var _ = Service("payments", func() {

	// Mobile Checkout APIs allow you to initiate Customer to
	// Business (C2B) payments on a mobile subscriber’s device.
	// This allows for a smoother checkout experience, since the
	// client will no longer need to remember the amount or an
	// account number to complete the transaction.
	Method("MobileCheckout", func() {
		Description("Initiate Mobile C2B payments on a mobile subscriber’s device.")
		Payload(MobileCheckoutPayload)
		Result(MobileCheckoutResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/mobile/checkout/request
			// Sandbox: https://payments.sandbox.africastalking.com/mobile/checkout/request
			POST("/mobile/checkout/request")
			Response(StatusCreated)
		})
	})

	Method("MobileB2C", func() {
		Description("Initiate Mobile B2C request")
		Payload(MobileB2CPayload)
		Result(MobileB2CResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/mobile/b2c/request
			// Sandbox: https://payments.sandbox.africastalking.com/mobile/b2c/request
			POST("/mobile/b2c/request")
			Response(StatusCreated)
		})
	})

	Method("MobileB2B", func() {
		Description("Initiate Mobile B2B request")
		Payload(MobileB2BPayload)
		Result(MobileB2BResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/mobile/b2b/request
			// Sandbox: https://payments.sandbox.africastalking.com/mobile/b2b/request
			POST("/mobile/b2b/request")
			Response(StatusCreated)
		})
	})

	// **************************************** BANK ***********************************************
	//
	// Bank Checkout:
	// Collect money into your payment wallet by initiating transactions
	// that deduct money from a customers bank account.
	Method("Bank Checkout", func() {
		Description("Collect money into your payment wallet.")
		Payload(BankCheckoutPayload)
		Result(BankCheckoutResponse)
		HTTP(func() {

			// Live: https://payments.africastalking.com/bank/checkout/charge
			// Sandbox: https://payments.sandbox.africastalking.com/bank/checkout/charge
			POST("/bank/checkout/charge")
			Response(StatusCreated)
		})
	})

	// Bank Checkout Validate:
	// Validate a bank checkout charge request.
	Method("BankCheckoutValidate", func() {
		Description("Validate a bank checkout charge request")
		Payload(BankCheckoutValidatePayload)
		Result(BankCheckoutValidateResponse)
		HTTP(func() {

			// Live: https://payments.africastalking.com/bank/checkout/validate
			// Sandbox: https://payments.sandbox.africastalking.com/bank/checkout/validate
			POST("/bank/checkout/validate")
			Response(StatusCreated)
		})
	})

	// Bank Transfer:
	// Move money from your Payment Wallet to a bank account.
	Method("BankTransfer", func() {
		Description("Initiate a bank transfer request.")
		Payload(BankTransferPayload)
		Result(BankTransferResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/bank/transfer
			// Sandbox: https://payments.sandbox.africastalking.com/bank/transfer
			POST("/bank/transfer")
			Response(StatusCreated)
		})
	})

	// **************************************** CARD ***********************************************
	// Card Checkout:
	// Collect money into your Payment Wallet by initiating transactions
	// that deduct money from a customers Debit or Credit Card.
	Method("CardCheckout", func() {
		Description("Collect money into your Payment Wallet")
		Payload(CardCheckoutPayload)
		Result(CardCheckoutResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/card/checkout/charge
			// Sandbox: https://payments.sandbox.africastalking.com/card/checkout/charge
			POST("/card/checkout/charge")
			Response(StatusCreated)
		})
	})

	// Card Checkout Validate:
	// Card checkout validation APIs allow your application to validate card checkout charge requests.
	Method("CardCheckoutValidate", func() {
		Description("Validate card checkout charge requests")
		Payload(CardCheckoutValidatePayload)
		Result(CardCheckoutValidateResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/card/checkout/validate
			// Sandbox: https://payments.sandbox.africastalking.com/card/checkout/validate
			POST("/card/checkout/validate")
			Response(StatusCreated)
		})
	})

	// **************************************** WALLET TRANSFER ***********************************************
	//
	// Wallet Transfer
	// Wallet transfer APIs allow you to transfer money from one Payment Product
	// to another Payment Product hosted on Africa’s Talking.
	Method("WalletTransfer", func() {
		Description("Transfer money from one product to another.")
		Payload(WalletTransferPayload)
		Result(WalletTransferResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/transfer/wallet
			// Sandbox: https://payments.sandbox.fricastalking.com/transfer/wallet
			POST("/transfer/wallet")
			Response(StatusCreated)
		})
	})

	// ****************************************   TOPUP STASH   ***********************************************
	//
	// Topup Stash
	// Topup stash APIs allow you to move money from a Payment Product to an Africa’s Talking application stash.
	// An application stash is the wallet that funds your service usage expenses.
	Method("TopupStash", func() {
		Description("Move money from a Payment Product to an application stash.")
		Payload(TopupStashPayload)
		Result(TopupStashResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/topup/stash
			// Sandbox: https://payments.sandbox.africastalking.com/topup/stash
			POST("/topup/stash")
			Response(StatusCreated)
		})
	})
})

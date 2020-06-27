package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("africastalking", func() {

	HTTP(func() {
		Path("/")
	})

	// ****************************************   1. SMS   ***********************************************

	// Send Bulk SMS
	Method("SendBulkSMS", func() {
		Description("Send Bulk SMS")
		Payload(BulkPayload)
		Result(BulkResponse)
		HTTP(func() {
			POST("version1/messaging")
			Response(StatusCreated)
		})
	})

	// Send Premium SMS
	Method("SendPremiumSMS", func() {
		Description("Send Premium SMS")
		Payload(PremiumPayload)
		Result(PremiumResponse)
		HTTP(func() {
			POST("version1/messaging")
			Response(StatusCreated)
		})
	})

	// Incrementally fetch messages from application inbox.
	Method("FetchSMS", func() {
		Description("Incrementally fetch messages from application inbox.")
		Payload(FetchMsgPayload)
		Result(FetchMsgResponse)
		HTTP(func() {
			GET("version1/messaging")
			Response(StatusOK)
		})
	})

	// Generate a checkout token
	Method("NewCheckoutToken", func() {
		Description("Generate a checkout token")
		Payload(CheckoutTokenPayload)
		Result(CheckoutTokenResponse)
		HTTP(func() {
			POST("checkout/token/create")
			Response(StatusCreated)
		})
	})

	// Subscribe a phone number
	Method("NewPremiumSubscription", func() {
		Description("Subscribe a phone number")
		Payload(NewSubPayload)
		Result(NewSubResponse)
		HTTP(func() {
			POST("version1/subscription/create")
			Response(StatusCreated)
		})
	})

	// Incrementally fetch your premium sms subscriptions.
	Method("FetchPremiumSubscription", func() {
		Description("Incrementally fetch your premium sms subscriptions.")
		Payload(FetchSubPayload)
		Result(FetchSubResponse)
		HTTP(func() {
			GET("version1/subscription")
			Response(StatusOK)
		})
	})

	// Delete a Premium SMS Subscription
	Method("PurgePremiumSubscription", func() {
		Description("Delete a Premium SMS Subscription")
		Payload(PurgeSubPayload)
		Result(PurgeSubResponse)
		HTTP(func() {
			POST("version1/subscription/delete")
			Response(StatusCreated)
		})
	})

	// ****************************************   3. VOICE   ***********************************************

	// Makes outbound calls.
	Method("MakeCall", func() {
		Description("Makes outbound calls.")
		Payload(MakeCallPayload)
		Result(MakeCallResponse)
		HTTP(func() {
			POST("call")
			Response(StatusCreated)
		})
	})

	// Transfers call to another number.
	Method("TransferCall", func() {
		Description("Transfers call to another number.")
		Payload(CallTransferPayload)
		Result(CallTransferResponse)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Used when you have more calls than you can handle at one time
	Method("Queue", func() {
		Description("Used when you have more calls than you can handle at one time")
		Payload(QueuedCallsPayload)
		Result(QueuedStatusResult)
		HTTP(func() {
			POST("queueStatus")
			Response(StatusCreated)
		})
	})

	// Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav
	Method("UploadMedia", func() {
		Description("Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav")
		Payload(UploadMediaFile)
		Result(String)
		HTTP(func() {
			POST("mediaUpload")
			Response(StatusCreated)
		})
	})

	// Set a text to be read out to the caller.
	Method("Say", func() {
		Description("Set a text to be read out to the caller.")
		Payload(SayPayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Play back an audio file located anywhere on the web.
	Method("Play", func() {
		Description("Play back an audio file located anywhere on the web.")
		Payload(PlayPayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Get digits a user enters on their phone in response to a prompt from application
	Method("GetDigits", func() {
		Description("Get digits a user enters on their phone in response to a prompt from application")
		Payload(GetDigitsPayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Connect the user who called your phone number to an external phone number.
	Method("Dial", func() {
		Description("Connect the user who called your phone number to an external phone number.")
		Payload(DialPayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Record a call session into an mp3 file.
	Method("Record", func() {
		Description("Record a call session into an mp3 file.")
		Payload(RecordPayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Pass an incoming call to a queue to be handled later.
	Method("Enqueue", func() {
		Description("Pass an incoming call to a queue to be handled later.")
		Payload(EnqueuePayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Pass the calls enqueued to a separate number to be handled.
	Method("Dequeue", func() {
		Description("Pass the calls enqueued to a separate number to be handled.") // e.g by an agent.
		Payload(DequeuePayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Transfer control of the call to the script whose URL is passed in
	Method("Redirect", func() {
		Description("Transfer control of the call to the script whose URL is passed in.")
		Payload(RedirectPayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// Reject an incoming call without incurring any usage charges.
	Method("Reject", func() {
		Description("Reject an incoming call without incurring any usage charges.")
		Payload(RejectPayload)
		Result(String)
		HTTP(func() {
			POST("callTransfer")
			Response(StatusCreated)
		})
	})

	// ****************************************   4. PAYMENTS   ***********************************************

	// Initiate Mobile C2B payments on a mobile subscriber’s device.
	Method("MobileCheckout", func() {
		Description("Initiate Mobile C2B payments on a mobile subscriber’s device.")
		Payload(MobileCheckoutPayload)
		Result(MobileCheckoutResponse)
		HTTP(func() {
			POST("mobile/checkout/request")
			Response(StatusCreated)
		})
	})

	// Initiate Mobile B2C request
	Method("MobileB2C", func() {
		Description("Initiate Mobile B2C request")
		Payload(MobileB2CPayload)
		Result(MobileB2CResponse)
		HTTP(func() {
			POST("mobile/b2c/request")
			Response(StatusCreated)
		})
	})

	// Initiate Mobile B2B request
	Method("MobileB2B", func() {
		Description("Initiate Mobile B2B request")
		Payload(MobileB2BPayload)
		Result(MobileB2BResponse)
		HTTP(func() {
			POST("mobile/b2b/request")
			Response(StatusCreated)
		})
	})

	// Collect money into your payment wallet.
	Method("Bank Checkout", func() {
		Description("Collect money into your payment wallet.")
		Payload(BankCheckoutPayload)
		Result(BankCheckoutResponse)
		HTTP(func() {
			POST("bank/checkout/charge")
			Response(StatusCreated)
		})
	})

	// Validate a bank checkout charge request
	Method("BankCheckoutValidate", func() {
		Description("Validate a bank checkout charge request")
		Payload(BankCheckoutValidatePayload)
		Result(BankCheckoutValidateResponse)
		HTTP(func() {
			POST("bank/checkout/validate")
			Response(StatusCreated)
		})
	})

	// Initiate a bank transfer request.
	Method("BankTransfer", func() {
		Description("Initiate a bank transfer request.")
		Payload(BankTransferPayload)
		Result(BankTransferResponse)
		HTTP(func() {
			POST("bank/transfer")
			Response(StatusCreated)
		})
	})

	// Collect money into your Payment Wallet by initiating transactions that deduct
	// money from a customers Debit or Credit Card.
	Method("CardCheckout", func() {
		Description("Collect money into your Payment Wallet by initiating transactions that deduct money from a customers Debit or Credit Card.")
		Payload(CardCheckoutPayload)
		Result(CardCheckoutResponse)
		HTTP(func() {
			POST("card/checkout/charge")
			Response(StatusCreated)
		})
	})

	// Allows your application to validate card checkout charge requests.
	Method("CardCheckoutValidate", func() {
		Description("Validate card checkout charge requests")
		Payload(CardCheckoutValidatePayload)
		Result(CardCheckoutValidateResponse)
		HTTP(func() {
			POST("card/checkout/validate")
			Response(StatusCreated)
		})
	})

	// Transfer money from one Payment Product to another Payment Product hosted on Africa’s Talking.
	Method("WalletTransfer", func() {
		Description("Transfer money from one Payment Product to another Payment Product hosted on Africa’s Talking.")
		Payload(WalletTransferPayload)
		Result(WalletTransferResponse)
		HTTP(func() {
			POST("transfer/wallet")
			Response(StatusCreated)
		})
	})

	// Move money from a Payment Product to an Africa’s Talking application stash.
	Method("TopupStash", func() {
		Description("Move money from a Payment Product to an Africa’s Talking application stash.")
		Payload(TopupStashPayload)
		Result(TopupStashResponse)
		HTTP(func() {
			POST("topup/stash")
			Response(StatusCreated)
		})
	})

	// Fetch transactions of a particular payment product.
	Method("FindTransaction", func() {
		Description("Fetch transactions of a particular payment product.")
		Payload(FindTransactionPayload)
		Result(FindTransactionResponse)
		HTTP(func() {
			GET("query/transaction/find")
			Response(StatusOK)
		})
	})

	// Fetch transactions of a particular payment product.
	Method("FetchProductTransactions", func() {
		Description("Fetch transactions of a particular payment product.")
		Payload(ProductTransactionsPayload)
		Result(ProductTransactionsResponse)
		HTTP(func() {
			GET("query/transaction/fetch")
			Response(StatusOK)
		})
	})

	// Fetch your wallet transactions.
	Method("FetchWalletTransactions", func() {
		Description("Fetch your wallet transactions")
		Payload(WalletTransactionsPayload)
		Result(WalletTransactionsResponse)
		HTTP(func() {
			GET("query/wallet/fetch")
			Response(StatusOK)
		})
	})

	// Fetch your wallet balance
	Method("FetchWalletBalance", func() {
		Description("Fetch your wallet balance")
		Payload(WalletBalancePayload)
		Result(WalletBalanceResponse)
		HTTP(func() {
			GET("query/wallet/balance")
			Response(StatusOK)
		})
	})

	// ****************************************   5. AIRTIME   ***********************************************

	// Send Airtime.
	Method("SendAirtime", func() {
		Description("Send Airtime.")
		Payload(AirtimePayload)
		Result(AirtimeResponse)
		HTTP(func() {
			POST("version1/airtime/send")
			Response(StatusOK)
		})
	})

	// ****************************************   6. IOT   ***********************************************

	// Publishes messages to remote devices.
	Method("PublishIoT", func() {
		Description("Publishes messages to remote devices.")
		Payload(IoTPayload)
		Result(IoTResponse)
		HTTP(func() {
			POST("data/publish")
			Response(StatusOK)
		})
	})

	// ****************************************   7. USER   ***********************************************

	// Initiate an application data request.
	Method("InitiateAppData", func() {
		Description("Initiate an application data request.")
		Payload(String)
		Result(UserResponse)
		HTTP(func() {
			GET("version1/user")
			Response(StatusOK)
		})
	})

	// Generates a valid auth token
	Method("Generate", func() {
		Description("Generates a valid auth token")
		Payload(func() {
			Attribute("username", String, func() {
				Description("Africa's Talking Username.")
				Example("sandbox")
				Default("sandbox")
			})
			Attribute("apiKey", String, func() {
				Description("Africa's Talking API Key.")
			})
			Required("username", "apiKey")
		})
		Result(AccessTokenResponse)
		HTTP(func() {
			POST("/auth-token/generate")
			Response(StatusCreated)
		})
	})
})

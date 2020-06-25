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

	Method("SendBulkSMS", func() {
		Description("Send Bulk SMS")
		Payload(BulkPayload)
		Result(BulkResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/messaging
			// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
			POST("version1/messaging")
			Response(StatusCreated)
		})
	})

	Method("SendPremiumSMS", func() {
		Description("Send Premium SMS")
		Payload(PremiumPayload)
		Result(PremiumResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/messaging
			// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
			POST("version1/messaging")
			Response(StatusCreated)
		})
	})

	Method("FetchSMS", func() {
		Description("Incrementally fetch messages from application inbox.")
		Payload(FetchMsgPayload)
		Result(FetchMsgResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/messaging
			// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
			GET("version1/messaging?{username}&{lastReceivedId}")
			Params(func() {
				Param("username:username", String, "Africa’s Talking application username.")
				Param("lastReceivedId:lastReceivedId", String, "ID of the message last processed.")
				Required("username")
			})
			Response(StatusOK)
		})
	})

	Method("NewCheckoutToken", func() {
		Description("Generate a checkout token")
		Payload(CheckoutTokenPayload)
		Result(CheckoutTokenResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/checkout/token/create
			// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
			POST("checkout/token/create")
			Response(StatusCreated)
		})
	})

	Method("NewPremiumSubscription", func() {
		Description("Subscribe a phone number")
		Payload(NewSubPayload)
		Result(NewSubResponse)
		HTTP(func() {

			// Live: https://content.africastalking.com/version1/subscription/create
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create
			POST("version1/subscription/create")
			Response(StatusCreated)
		})
	})

	Method("FetchPremiumSubscription", func() {
		Description("Incrementally fetch your premium sms subscriptions.")
		Payload(FetchSubPayload)
		Result(FetchSubResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/subscription
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription
			GET("version1/subscription?username={username}&shortCode={shortCode}&keyword={keyword}&lastReceivedId={lastReceivedId}")
			Params(func() {
				Param("username:username", String, "Africa’s Talking application username.")
				Param("shortCode:shortCode", String, "Premium short code mapped to your account")
				Param("keyword:keyword", String, "Premium keyword under short code mapped to your account.")
				Param("lastReceivedId:lastReceivedId", String, "ID of the message last processed.")
				Required("username", "shortCode", "keyword")
			})
			Response(StatusOK)
		})
	})

	Method("PurgePremiumSubscription", func() {
		Description("Delete a Premium SMS Subscription")
		Payload(PurgeSubPayload)
		Result(PurgeSubResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/subscription/delete
			// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete
			POST("version1/subscription/delete")
			Response(StatusCreated)
		})
	})

	// ****************************************   3. VOICE   ***********************************************
	Method("MakeCall", func() {
		Description("Makes outbound calls.")
		Payload(MakeCallPayload)
		Result(MakeCallResponse)
		HTTP(func() {

			// Live: https://voice.africastalking.com/call
			// Sandbox: https://voice.sandbox.africastalking.com/call
			POST("call")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("TransferCall", func() {
		Description("Transfers call to another number.")
		Payload(CallTransferPayload)
		Result(CallTransferResponse)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Say", func() {
		Description("Set a text to be read out to the caller.")
		Payload(Say)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Play", func() {
		Description("Play back an audio file located anywhere on the web.")
		Payload(Play)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("GetDigits", func() {
		Description("Get digits a user enters on their phone in response to a prompt from application")
		Payload(GetDigits)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
		})
	})

	Method("Dial", func() {
		Description("Connect the user who called your phone number to an external phone number.")
		Payload(Dial)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Record", func() {
		Description("Record a call session into an mp3 file.")
		Payload(Record)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Enqueue", func() {
		Description("Pass an incoming call to a queue to be handled later.")
		Payload(Enqueue)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Dequeue", func() {
		Description("Pass the calls enqueued to a separate number to be handled.") // e.g by an agent.
		Payload(Dequeue)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Redirect", func() {
		Description("Transfer control of the call to the script whose URL is passed in.")
		Payload(Redirect)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Reject", func() {
		Description("Reject an incoming call without incurring any usage charges.")
		Payload(Reject)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("Queue", func() {
		Description("Used when you have more calls than you can handle at one time")
		Payload(QueuedCallsPayload)
		Result(QueuedStatusResult)
		HTTP(func() {

			// Live: https://voice.africastalking.com/queueStatus
			// Sandbox: https://voice.sandbox.africastalking.com/queueStatus
			POST("queueStatus")
			Headers(func() {})
			Response(StatusCreated)
		})
	})

	Method("UploadMedia", func() {
		Description("Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav")
		Payload(UploadMediaFile)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/mediaUpload
			// Sandbox: https://voice.sandbox.africastalking.com/mediaUpload
			POST("mediaUpload")
			Headers(func() {})
			Response(StatusCreated)
		})
	})

	// ****************************************   4. PAYMENTS   ***********************************************

	Method("MobileCheckout", func() {
		Description("Initiate Mobile C2B payments on a mobile subscriber’s device.")
		Payload(MobileCheckoutPayload)
		Result(MobileCheckoutResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/mobile/checkout/request
			// Sandbox: https://payments.sandbox.africastalking.com/mobile/checkout/request
			POST("mobile/checkout/request")
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
			POST("mobile/b2c/request")
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
			POST("mobile/b2b/request")
			Response(StatusCreated)
		})
	})

	Method("Bank Checkout", func() {
		Description("Collect money into your payment wallet.")
		Payload(BankCheckoutPayload)
		Result(BankCheckoutResponse)
		HTTP(func() {

			// Live: https://payments.africastalking.com/bank/checkout/charge
			// Sandbox: https://payments.sandbox.africastalking.com/bank/checkout/charge
			POST("bank/checkout/charge")
			Response(StatusCreated)
		})
	})

	Method("BankCheckoutValidate", func() {
		Description("Validate a bank checkout charge request")
		Payload(BankCheckoutValidatePayload)
		Result(BankCheckoutValidateResponse)
		HTTP(func() {

			// Live: https://payments.africastalking.com/bank/checkout/validate
			// Sandbox: https://payments.sandbox.africastalking.com/bank/checkout/validate
			POST("bank/checkout/validate")
			Response(StatusCreated)
		})
	})

	Method("BankTransfer", func() {
		Description("Initiate a bank transfer request.")
		Payload(BankTransferPayload)
		Result(BankTransferResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/bank/transfer
			// Sandbox: https://payments.sandbox.africastalking.com/bank/transfer
			POST("bank/transfer")
			Response(StatusCreated)
		})
	})

	Method("CardCheckout", func() {
		Description("Collect money into your Payment Wallet")
		Payload(CardCheckoutPayload)
		Result(CardCheckoutResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/card/checkout/charge
			// Sandbox: https://payments.sandbox.africastalking.com/card/checkout/charge
			POST("card/checkout/charge")
			Response(StatusCreated)
		})
	})

	Method("CardCheckoutValidate", func() {
		Description("Validate card checkout charge requests")
		Payload(CardCheckoutValidatePayload)
		Result(CardCheckoutValidateResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/card/checkout/validate
			// Sandbox: https://payments.sandbox.africastalking.com/card/checkout/validate
			POST("card/checkout/validate")
			Response(StatusCreated)
		})
	})

	Method("WalletTransfer", func() {
		Description("Transfer money from one product to another.")
		Payload(WalletTransferPayload)
		Result(WalletTransferResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/transfer/wallet
			// Sandbox: https://payments.sandbox.fricastalking.com/transfer/wallet
			POST("transfer/wallet")
			Response(StatusCreated)
		})
	})

	Method("TopupStash", func() {
		Description("Move money from a Payment Product to an application stash.")
		Payload(TopupStashPayload)
		Result(TopupStashResponse)

		HTTP(func() {

			// Live: https://payments.africastalking.com/topup/stash
			// Sandbox: https://payments.sandbox.africastalking.com/topup/stash
			POST("topup/stash")
			Response(StatusCreated)
		})
	})

	// ****************************************   5. AIRTIME   ***********************************************
	Method("SendAirtime", func() {
		Description("Send Airtime.")
		Payload(AirtimePayload)
		Result(AirtimeResponse)
		HTTP(func() {

			// Live: https://api.africastalking.com/version1/airtime/send
			//Sandbox: https://api.sandbox.africastalking.com/version1/airtime/send
			POST("version1/airtime/send")
			Response(StatusOK)
		})
	})

	// ****************************************   6. IOT   ***********************************************
	Method("PublishIoT", func() {
		Description("Publishes messages to remote devices.")
		Payload(IoTPayload)
		Result(IoTResponse)
		HTTP(func() {

			// POST request to https://iot.africastalking.com/data/publish
			POST("data/publish")
			Response(StatusOK)
		})
	})

	// ****************************************   7. USER   ***********************************************
	Method("InitiateAppData", func() {
		Description("Initiate an application data request.")
		Payload(String)
		Result(UserResponse)
		HTTP(func() {

			// Initiate an application data request by making a HTTP GET request to the following endpoint:
			// Live: https://api.africastalking.com/version1/user
			// Sandbox: https://api.sandbox.africastalking.com/version1/user
			GET("version1/user")
			Response(StatusOK)
		})
	})

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

			// POST creates a route using the POST HTTP method.
			// POST request to https://api.africastalking.com/auth-token/generate
			POST("/auth-token/generate")
			Response(StatusOK)
		})
	})

})

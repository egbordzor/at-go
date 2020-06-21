package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("voice", func() {

	Method("make", func() {
		Description("Makes outbound calls.")
		Payload(MakeCallPayload)
		Result(MakeCallResponse)
		HTTP(func() {

			// Live: https://voice.africastalking.com/call
			// Sandbox: https://voice.sandbox.africastalking.com/call
			POST("/call")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("transfer", func() {
		Description("Transfers call to another number.")
		Payload(CallTransferPayload)
		Result(CallTransferResponse)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("say", func() {
		Description("Set a text to be read out to the caller.")
		Payload(Say)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("play", func() {
		Description("Play back an audio file located anywhere on the web.")
		Payload(Play)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("getDigits", func() {
		Description("Get digits a user enters on their phone in response to a prompt from application")
		Payload(GetDigits)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
		})
	})

	Method("dial", func() {
		Description("Connect the user who called your phone number to an external phone number.")
		Payload(Dial)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("record", func() {
		Description("Record a call session into an mp3 file.")
		Payload(Record)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("enqueue", func() {
		Description("Pass an incoming call to a queue to be handled later.")
		Payload(Enqueue)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("dequeue", func() {
		Description("Pass the calls enqueued to a separate number to be handled.") // e.g by an agent.
		Payload(Dequeue)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("redirect", func() {
		Description("Transfer control of the call to the script whose URL is passed in.")
		Payload(Redirect)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("reject", func() {
		Description("Reject an incoming call without incurring any usage charges.")
		Payload(Reject)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Headers(func() {})
			Response(StatusOK)
		})
	})

	Method("queue", func() {
		Description("Used when you have more calls than you can handle at one time")
		Payload(QueuedCallsPayload)
		Result(QueuedStatusResult)
		HTTP(func() {

			// Live: https://voice.africastalking.com/queueStatus
			// Sandbox: https://voice.sandbox.africastalking.com/queueStatus
			POST("/queueStatus")
			Headers(func() {})
			Response(StatusCreated)
		})
	})

	Method("upload", func() {
		Description("Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav")
		Payload(UploadMediaFile)
		Result(String)
		HTTP(func() {

			// Live: https://voice.africastalking.com/mediaUpload
			// Sandbox: https://voice.sandbox.africastalking.com/mediaUpload
			POST("/mediaUpload")
			Headers(func() {})
			Response(StatusCreated)
		})
	})
})

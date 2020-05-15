package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// 1. User calls via AT gateway.
// 2. AT Gateway sends POST request to our web application.
// 3. Our application sends an XML response to AT Gateway.
// 4. AT Gateway responds to User.
var _ = Service("voice", func() {
	Title("Voice Service")
	Description("")
	Error("unauthorized", String, "Credentials are invalid")
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("make", func() {
		Description("Makes an outbound calls through the Africa'sTalking Voice API")
		Payload(MakeCallPayload)
		Result(MakeCallResult)
		//Error()
		HTTP(func() {
			Headers(func() {
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/x-www-form-urlencoded")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("Content-Type")
			})

			// Live: https://voice.africastalking.com/call
			// Sandbox: https://voice.sandbox.africastalking.com/call
			POST("/call")
			Response(StatusOK)
		})
	})

	Method("transfer", func() {
		Description("Transfers call to another number")
		Payload(CallTransferPayload)
		Result(CallTransferResult)
		//Error()
		HTTP(func() {
			Headers(func() {
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/json")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Response(StatusOK)
		})
	})

	Method("queue", func() {
		Description("Used when you have more calls than you can handle at one time")
		Payload(QueuedCallsPayload)
		Result(QueuedStatusResult)
		//Error()
		HTTP(func() {
			Headers(func() {
				Attribute("apiKey", String, "Africa’s Talking application apiKey.")
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/x-www-form-urlencoded")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("Content-Type", "Accept")
			})

			// Live: https://voice.africastalking.com/queueStatus
			// Sandbox: https://voice.sandbox.africastalking.com/queueStatus
			POST("/queueStatus")

			// You can check the HTTP Response Code to determine whether the request was successful.
			// Any response code other than 201 (Created) indicates that the call was not initiated
			Response(StatusCreated)
		})
	})

	Method("upload", func() {
		Description("Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav")
		Payload(UploadMediaFile)
		Result(String)
		//Error()
		HTTP(func() {
			Headers(func() {
				Attribute("apiKey", String, "Africa’s Talking application apiKey.")
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/x-www-form-urlencoded")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("Content-Type", "Accept")
			})

			// Live: https://voice.africastalking.com/mediaUpload
			// Sandbox: https://voice.sandbox.africastalking.com/mediaUpload
			POST("/mediaUpload")

			// You can check the HTTP Response Code to determine whether the request was successful.
			// Any response code other than 201 (Created) indicates that the call was not initiated
			Response(StatusCreated)
		})
	})
})

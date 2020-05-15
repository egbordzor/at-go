package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("media", func() {
	Title("Upload Media using Africa'sTalking Voice API")
	Description("Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav")
	Method("add", func() {
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

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Live: https://voice.africastalking.com/mediaUpload
// Sandbox: https://voice.sandbox.africastalking.com/mediaUpload
// Any response code other than 201 (Created) indicates that the call was not initiated.
var UploadMediaFile = Type("UploadMediaFile", func() {
	// With the extension .mp3 or .wav only
	Description("Uploads media or audio files to Africa'sTalking servers")
	Attribute("username", String, func() {
		Description("Your Africa’s Talking application username.")
	})
	// This contains the audio file you want played during a call.
	Attribute("url", String, func() {
		// Don’t forget to start with http://
		Description("The url of the file to upload.")
	})
})

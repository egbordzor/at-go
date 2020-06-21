package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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

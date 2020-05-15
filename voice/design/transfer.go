package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Service describes a Calls Transfer Service
var _ = Service("transfer", func() {
	Title("Calls Transfer Service")

	// Method describes a service method (endpoint)
	Method("add", func() {
		Description("Transfers call to another number")

		// Payload describes the method payload.
		Payload(CallTransferPayload)

		// Result describes the method result.
		Result(CallTransferResult)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
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

			// Requests to the service consist of HTTP POST requests.
			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Response(StatusOK)
		})
	})
})

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Live: https://voice.africastalking.com/callTransfer
// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
var CallTransferPayload = Type("CallTransferPayload", func() {
	Description("Transfer calls to another number")
	Attribute("sessionId", String, func() {
		Description("Session Id of the ongoing call, it must have two legs")
	})
	Attribute("phoneNumber", String, func() {
		Description("Phone Number to transfer the call to")
	})
	Attribute("callLeg", String, func() {
		Description("Call leg to transfer the call to either caller or callee")
		Enum("caller", "callee")
		Default("callee") // (Defaults to callee)
	})
	Attribute("holdMusicUrl", String, func() {
		Description("The url of the media file to be played when the user is on hold.") // Don’t forget to start with http://
	})
	Required("sessionId", "phoneNumber")
})

var CallTransferResult = ResultType("CallTransferResult", func() {
	Attribute("status", String, func() {
		Description("can be either Success or Aborted")
		Enum("Success", "Aborted")
	})
	Attribute("errorMessage", String, "Why the transfer ws aborted None is successful")
})

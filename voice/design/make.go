package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("make", func() {
	Title("Make Call using Africa'sTalking Voice API")
	Description("Makes an outbound calls through the Africa'sTalking Voice API")
	Method("add", func() {
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
})

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Content-Type: application/x-www-form-urlencoded or multipart/form-data
// Live: https://voice.africastalking.com/call
// Sandbox: https://voice.sandbox.africastalking.com/call
var MakeCallPayload = Type("MakeCallPayload", func() {
	Description("Makes an outbound call through the Africa'sTalking  Voice API")
	Attribute("username", String, func() {
		Description("Africa’s Talking Username")
	})
	Attribute("from", String, func() {
		Description("Africa’s Talking phone number") // In international format i.e. +XXXYYYYYY
	})
	Attribute("to", String, func() {
		Description("A comma separated string of recipients’ phone numbers.")
	})
	Attribute("clientRequestId", String, func() {
		Description("Variable sent to Events Callback URL used to tag the call")
	})

	Required("username", "from", "to")
})

var MakeCallResult = ResultType("MakeCallResult", func() {
	Description("Voice call response")
	TypeName("MakeCallResult")
	ContentType("application/xml")

	// A list with multiple Entry each corresponding
	// to an individual phone number and their status.
	// Entry is a Map with details of queued numbers.
	Attribute("entries", ArrayOf(Entry))
	Attribute("errorMessage", String, func() {
		Description("Error message if ENTIRE request was rejected by the API.")
	})

	View("default", func() {
		Attribute("entries")
		Attribute("errorMessage")
	})
})

var Entry = Type("Entry", func() {
	Attribute("phoneNumber", String, func() {
		Description("The phone number queued.")
	})
	Attribute("status", String, func() {
		Enum(
			// The call request has been accepted and queued
			"Queued",

			// Recipient number is in an incorrect format
			"InvalidPhoneNumber",

			// Recipient number is outside the supported zone
			"DestinationNotSupported",

			// Your AfricasTalking account has insufficient balance.
			"InsufficientCredit",
		)
	})
	Attribute("sessionId", String, func() {
		Description("A unique id for the request associated to this phone number")
		Default("None") // Defaults to None if an error occurred.
	})
})

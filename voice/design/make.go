package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Content-Type: application/x-www-form-urlencoded or multipart/form-data
// Live: https://voice.africastalking.com/call
// Sandbox: https://voice.sandbox.africastalking.com/call

// Make an outbound call
var MakeCallPayload = Type("MakeCallPayload", func() {
	Description("Makes an outbound call.")
	Attribute("username", String, func() {
		Description("Africa’s Talking Username")
	})
	Attribute("from", String, func() {

		// In international format i.e. +XXXYYYYYY
		Description("Africa’s Talking phone number")
	})
	Attribute("to", String, func() {
		Description("A comma separated string of recipients’ phone numbers.")
	})
	Attribute("clientRequestId", String, func() {
		Description("Variable sent to Events Callback URL used to tag the call")
	})
	Required("username", "from", "to")
})

var MakeCallResponse = ResultType("MakeCallResponse", func() {
	Description("Outbound call HTTP response.")
	TypeName("MakeCallResponse")
	ContentType("application/xml")

	// A list with multiple Entry each corresponding
	// to an individual phone number and their status.
	// Entry is a Map with details of queued numbers.
	Attribute("entries", ArrayOf(VoiceEntry))
	Attribute("errorMessage", String, func() {
		Description("Error message if ENTIRE request was rejected by the API.")
	})

	View("default", func() {
		Attribute("entries")
		Attribute("errorMessage")
	})
})

var VoiceEntry = Type("VoiceEntry", func() {
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

		// Defaults to None if an error occurred.
		Default("None")
	})
})

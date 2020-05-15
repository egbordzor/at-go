package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Service describes an Outbound Calls Service
var _ = Service("outbound", func() {
	Title("Outbound Calls Service")

	// Method describes a service method (endpoint)
	Method("add", func() {
		Description("Makes outbound calls through the Africa'sTalking Voice API")

		// Payload describes the method payload.
		Payload(MakeCallPayload)

		// Result describes the method result.
		Result(MakeCallResult)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
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

			// Requests to the service consist of HTTP POST requests.
			// Live: https://voice.africastalking.com/call
			// Sandbox: https://voice.sandbox.africastalking.com/call
			POST("/call")

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
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

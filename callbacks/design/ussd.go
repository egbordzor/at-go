package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// USSD Callback Service
// This service hosts the following notifications.
// 1. USSD Notifications
var _ = Service("ussd", func() {

	HTTP(func() {
		Path("/africastalking/ussd/version1")
	})

	Method("newUSSDNotification", func() {
		Description("USSD Notifications")
		Payload(USSDPayload)
		Result(USSDResponse)
		HTTP(func() {
			POST("/sessions")
			Response(StatusCreated)
		})
	})
})
var USSDPayload = Type("USSDPayload", func() {
	Description("Request made whenever user dials a USSD code or responds to a menu.")

	// Sent every time a mobile subscriber response has been received.
	Attribute("sessionId", String, func() {
		Description("A unique value generated when the session starts.")
	})
	Attribute("phoneNumber", String, func() {
		Description("Mobile number of the subscriber interacting with USSD application.")
	})
	Attribute("networkCode", String, func() {
		Description("Telco of the mobile number interacting with USSD application.")
	})
	Attribute("serviceCode", String, func() {
		Description("USSD code assigned to application.")
	})

	// It is an empty string in the first notification of a session.
	// After that, it concatenates all the user input within the
	// session with a * until the session ends.
	Attribute("text", String, func() {
		Description("Shows the user input.")
	})
})

var USSDResponse = ResultType("USSDResponse", func() {
	Description("Echoed plain text response back to AT gateway")
	TypeName("USSDResponse")
	ContentType("text/plain")

	Attributes(func() {
		// Let the MSP know whether the session is complete or not.
		// If the session is ongoing, begin your response with CON.
		// If this is the last response for that session, begin your response with END.
		Attribute("response", String, func() {
			Description("Plain text response back to AT gateway")
		})
	})

	View("default", func() {
		Attribute("response")
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("ussd", func() {
	Title("USSD Callback")
	Description("USSD notifications sent from Africa'sTalking gateway")

	// USSD Callback
	Method("trigger", func() {
		Description("Callback URL that sends request data our App using HTTP POST.")
		Payload(USSDPayload)
		Result(String)
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
					Default("application/json")
				})
				Required("Content-Type")
			})
			POST("/ussd")
			Response(StatusOK) // 200 RFC 7231, 6.3.1
		})
	})

})

// 1. User dials USSD Code.
// 2. Africa'sTalking Gateway sends POST request to our web application.
// 3. Our application sends an plain text response to Africa'sTalking Gateway.
// 4. Africa'sTalking Gateway responds to user with USSD Menu.
var USSDPayload = Type("USSDPayload", func() {
	Description("This request is made when the user dials a USSD code and every time they respond to a menu.")

	// Sent every time a mobile subscriber response has been received.
	Attribute("sessionId", String, func() {
		Description("A unique value generated when the session starts.")
		// Let the Mobile Service Provider know whether the session is complete or not.
		// "CON" - Session is ongoing
		// "END" - Last session response
		Attribute("SessionStatus", String, func() {
			Enum("CON", "END")

		})
	})
	Attribute("phoneNumber", String, func() {
		Description("The number of the mobile subscriber interacting with your ussd application.")
	})
	Attribute("networkCode", String, func() {
		Description("The telco of the phoneNumber interacting with your ussd application.")
	})
	// Register a service code with AT.
	Attribute("serviceCode", String, func() {
		Description("This is the USSD code assigned to your application")
	})
	// It is an empty string in the first notification of a session.
	// After that, it concatenates all the user input within the
	// session with a * until the session ends.
	Attribute("text", String, func() {
		Description("This shows the user input.")
	})
	// Callback URL registered on AT
	// This can be called by AT whenever they get a request from a client coming into our system.
	Attribute("CallbackURL", String, func() {
		Description("Callback URL registered on AT")
		Pattern(`(?i)^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
		Format(FormatURI)
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("ussd", func() {
	Title("USSD Callback Service")
	Description("USSD Notifications sent from Africa'sTalking gateway")

	// 5xx Server errors
	// 500 Internal Server Error
	Error("internal_error", StatusInternalServerError)

	// 501 Not Implemented
	Error("not_implemented")

	// 502 Bad Gateway
	Error("bad_gateway")

	// 503 Service Unavailable
	Error("service_unavailable")

	// 504 Gateway Timeout
	Error("gateway_timeout")

	// USSD Callback
	Method("trigger", func() {
		Description("Callback URL that sends request data our App using HTTP POST.")
		Payload(USSDPayload)
		Result(USSDMedia)

		// 4xx Client errors
		// 400 Bad Request
		Error("bad_request")

		// 401 Unauthorized
		Error("unauthorized")

		// 403 Forbidden
		Error("forbidden")

		// 404 Not Found
		Error("not_found")

		// 405 Method Not Allowed
		Error("not_allowed")

		POST("/ussd")

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

			Response(StatusOK)

			// 4xx Client errors
			Response("bad_request", StatusBadRequest, func() {
				Tag("code", "bad_request")
			})
			Response("unauthorized", StatusUnauthorized, func() {
				Tag("code", "bad_request")
			})
			Response("forbidden", StatusForbidden, func() {
				Tag("code", "forbidden")
			})
			Response("not_found", StatusNotFound, func() {
				Tag("code", "not_found")
			})
			Response("not_allowed", StatusMethodNotAllowed, func() {
				Tag("code", "not_allowed")
			})
		})
	})

})

// 1. User dials USSD Code.
// 2. Africa'sTalking Gateway sends POST request to our web application.
// 3. Our application sends an plain text response to Africa'sTalking Gateway.
// 4. Africa'sTalking Gateway responds to user with USSD Menu.
var USSDPayload = Type("USSDPayload", func() {
	Description("Request made whenever user dials a USSD code or responds to a menu.")

	// Sent every time a mobile subscriber response
	// has been received.
	Attribute("sessionId", String, func() {
		Description("A unique value generated when the session starts.")
		// Let the Mobile Service Provider know whether
		// the session is complete or not.
		// "CON" - Session is ongoing
		// "END" - Last session response
		Attribute("SessionStatus", String, func() {
			Enum("CON", "END")

		})
	})
	Attribute("phoneNumber", String, func() {
		Description("Mobile number of the subscriber interacting with USSD application.")
	})
	Attribute("networkCode", String, func() {
		Description("Telco of the mobile number interacting with USSD application.")
	})
	// Register a service code with AT.
	Attribute("serviceCode", String, func() {
		Description("USSD code assigned to application.")
	})
	// It is an empty string in the first notification
	// of a session. After that, it concatenates all
	// the user input within the session with a * until
	// the session ends.
	Attribute("text", String, func() {
		Description("Shows the user input.")
	})
	// Callback URL registered on AT
	// This can be called by AT whenever they get a
	// request from a client coming into our system.
	Attribute("CallbackURL", String, func() {
		Description("Callback URL registered on AT")
		Pattern(`(?i)^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
		Format(FormatURI)
	})
})

var USSDMedia = ResultType("USSDMedia", func() {
	Description("Echoed plain text response back to AT gateway")
	TypeName("USSDMedia")
	ContentType("text/plain")

	Attribute("response", String, func() {
		Description("Plain text response back to AT gateway")
	})

	View("default", func() {
		Attribute("response")
	})
})

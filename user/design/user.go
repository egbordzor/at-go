package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// req.Header.Set("Accept", "application/json")
// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// req.Header.Set("Apikey", "MyAppApiKey")
var _ = Service("user", func() {
	Title("User Service")
	Description("Initiates an application data request")

	HTTP(func() {
		Path("/version1")
	})

	// Method describes a service method (endpoint)
	Method("fetch", func() {
		Description("Initiate an application data request.")

		// Payload describes the method payload.
		Payload(func() {
			Attribute("username", String, "username of the application making the request")
			Required("username")
		})

		// Result describes the method result.
		Result(UserMedia)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {

			// Requests to the service consist of HTTP GET requests.
			// Live: https://api.africastalking.com/version1/user
			//Sandbox: https://api.sandbox.africastalking.com/version1/user
			GET("/user?username={username}")

			Headers(func() {

				// Attribute describes an object field
				Attribute("apiKey", String, func() {
					Description("Africa’s Talking application apiKey.")
				})
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
				Required("apiKey", "Content-Type", "Accept")
			})

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})

})

// Live: https://api.africastalking.com/version1/user
// Sandbox: https://api.sandbox.africastalking.com/version1/user
var UserMedia = ResultType("UserMedia", func() {
	Description("Get application data response.")
	TypeName("UserMedia")
	ContentType("application/json")

	Attribute("UserData", MapOf(String, String, func() {
		Key(func() {
			Description("Your Africa’s Talking application balance")
			Default("balance")
		})
		Elem(func() {
			// Validates values of the map
			Pattern("[a-zA-Z]+")

			// The format of this string is:
			// (3-digit Currency Code)(space)(Decimal Value)
			Example("KES 1785.50")
		})
	}))
})

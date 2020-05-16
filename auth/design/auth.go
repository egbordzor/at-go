package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Service describes an Authentication MicroService
var _ = Service("auth", func() {
	Title("Authentication MicroService")

	// Error describes a method error return value.
	Error("unauthorized", String, "Credentials are invalid")

	// HTTP defines the HTTP transport specific properties
	// of an API, a service or a single method.
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Path("/auth-token")
	})

	// Method defines a single service method.
	Method("generate", func() {

		// Description sets the expression description.
		Description("Generates a valid auth token")

		// Payload defines the data type of an method input.
		Payload(func() {
			Attribute("username", String, func() {
				Description("Africa's Talking Username.")
				Example("sandbox")
				Default("sandbox")
			})
			Attribute("apiKey", String, func() {
				Description("Africa's Talking API Key.")
			})

			// Required adds a "required" validation to the attribute.
			Required("username", "username")
		})

		// Result defines the data type of a method output.
		Result(TokenMedia)

		// POST creates a route using the POST HTTP method.
		// POST request to https://api.africastalking.com/auth-token/generate
		POST("/generate")

		HTTP(func() {

			// Headers describes HTTP request/response or gRPC response headers.
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

				// Required adds a "required" validation to the attribute.
				Required("Content-Type", "Accept")
			})

			// Response describes a HTTP or a gRPC response.
			Response(StatusOK)
		})
	})
})

// TokenMedia defines a result type used to describe a method response
var TokenMedia = ResultType("TokenMedia", func() {
	Attribute("token", String, func() {
		Description("Generated Auth Token.")
		Example("ATtkn_abcdefghijklmnopqrstuvwxyz")
	})
	Attribute("lifetimeInSeconds", Int, func() {
		Description("Token Lifetime")
		Example(3600)
	})

	// Required adds a "required" validation to the attribute.
	Required("token", "lifetimeInSeconds")

	// View defines a view for the result type.
	View("default", func() {
		Attribute("token")
		Attribute("lifetimeInSeconds")
	})
})

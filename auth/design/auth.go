package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("auth", func() {
	Title("Authentication Service")

	Description("Authenticating with an Auth Token")

	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Path("/auth-token")
	})

	Method("generate", func() {
		Description("Generates a valid auth token")

		Payload(func() {
			Attribute("username", String, func() {
				Description("Africa's Talking Username.")
				Example("sandbox")
				Default("sandbox")
			})
			Attribute("apiKey", String, func() {
				Description("Africa's Talking API Key.")
			})

			Required("username", "username")
		})

		Result(TokenMedia)

		// POST request to https://api.africastalking.com/auth-token/generate
		POST("/generate")

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
				Required("Content-Type", "Accept")
			})

			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})
})

var TokenMedia = ResultType("TokenMedia", func() {
	Attribute("token", String, func() {
		Description("Generated Auth Token.")
		Example("ATtkn_abcdefghijklmnopqrstuvwxyz")
	})
	Attribute("lifetimeInSeconds", Int, func() {
		Description("Token Lifetime")
		Example(3600)
	})

	Required("token", "lifetimeInSeconds")

	View("default", func() {
		Attribute("token")
		Attribute("lifetimeInSeconds")
	})
})

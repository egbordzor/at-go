package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Authentication
// Requests made to AT APIs must be authenticated, there are two ways to do this:
// 1. Authenticating using your API Key and username
// 2. Authenticating using an Auth Token

var _ = Service("auth", func() {
	Title("Authentication Service")
	Description("Authenticating with an Auth Token")
	Error("unauthorized", String, "Credentials are invalid")
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("generate", func() {
		Description("Generates a valid auth token")
		Payload(TokenPayload)
		Result(TokenResponse)
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

			// POST request to https://api.africastalking.com/auth-token/generate
			// Username and API Key
			POST("/generate")

			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})
})

// Authenticating using your API Key and Username
var APIKeyHeader = Type("APIKeyHeader", func() {
	Attribute("apiKey", String, "Africa’s Talking application apiKey.")
	Attribute("Content-Type", String, func() {
		Description("The requests content type.")
		Enum("application/x-www-form-urlencoded", "application/json")
	})
	Attribute("Accept", String, func() {
		Description("The requests response type.")
		Enum("application/json", "application/xml")
		Default("application/xml")
	})
	Required("apiKey", "Content-Type")
})

// Authenticating with an Auth Token
var AuthTokenHeader = Type("AuthTokenHeader", func() {
	Attribute("authToken", String, "Generated Auth Token.")
	Attribute("Content-Type", String, func() {
		Description("The requests content type.")
		Enum("application/x-www-form-urlencoded", "application/xml")
	})
	Attribute("Accept", String, func() {
		Description("The requests response type.")
		Enum("application/json", "application/xml")
		Default("application/xml")
	})

})

var TokenPayload = Type("TokenPayload", func() {
	Attribute("username", String, "AT Username")
	Attribute("apiKey", String, "API Key ")

	Required("username", "username")
})

var TokenResponse = ResultType("TokenResponse", func() {
	Attribute("token", String, "Generated Auth Token.")
	Attribute("lifetimeInSeconds", Int, "Token Lifetime")

	Required("token", "lifetimeInSeconds")

	View("default", func() {
		Attribute("token")
		Attribute("lifetimeInSeconds")
	})
})

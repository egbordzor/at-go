package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("at", func() {
	Title("AfricasTalking API")
	Description("HTTP service for interacting with all AfricasTalking API.")
	Version("1.0")
	TermsOfService("https://github.com/wondenge/atgo/blob/master/LICENSE")
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/atgo/blob/master/LICENSE")
	})
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/atgo/blob/master/README.md")
	})
	Server("at", func() {
		Description("at hosts the AfricasTalking HTTP Service.")
		Services("africastalking", "health", "swagger")
		Host("development", func() {
			Description("Development hosts.")
			URI("https://api.sandbox.africastalking.com")
		})
		Host("production", func() {
			Description("Production hosts.")
			URI("https://{subdomain}.africastalking.com")
			Variable("subdomain", String, "Name of Sub-Domain", func() {
				Enum("api", "content")
				Default("api")
			})
		})
	})
})



var _ = Service("africastalking", func() {

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
		HTTP(func() {

			// POST creates a route using the POST HTTP method.
			// POST request to https://api.africastalking.com/auth-token/generate
			POST("/auth-token/generate")
			Response(StatusOK)
		})
	})

	Method("fetch", func() {
		Description("Initiate an application data request.")
		Payload(func() {
			Attribute("username", String, "username of the application making the request")
			Required("username")
		})
		Result(UserMedia)
		HTTP(func() {

			// Requests to the service consist of HTTP GET requests.
			// Live: https://api.africastalking.com/version1/user
			// Sandbox: https://api.sandbox.africastalking.com/version1/user
			GET("/user?username={username}")
			Response(StatusOK)
		})
	})

})

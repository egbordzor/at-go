package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Making an API call
// Include the API key in the request header as a field called apiKey.
// Include the Username as well, varied for different requests as following:
//
// 1. For GET requests e.g. fetch messages.
// The username should be passed as a query parameter.
//
// 2. For POST requests in which parameters are sent as a url encoded form
// e.g. in sending SMS.
// The username should be included as one of the parameters within the form.
//
// 3. For POST requests that require JSON in the request body
// e.g. in mobile checkout.
// The username should be included in the JSON sent in the body of the request.

// Authenticating an API using username and API Key.

var AccessTokenResponse = ResultType("AccessTokenResponse", func() {
	Description("Access Token Response")
	TypeName("AccessTokenResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("token", String, func() {
			Description("Generated Auth Token.")
			Example("ATtkn_abcdefghijklmnopqrstuvwxyz")
		})
		Attribute("lifetimeInSeconds", Int, func() {
			Description("Token Lifetime")
			Example(3600)
		})
		Required("token", "lifetimeInSeconds")
	})

	View("default", func() {
		Attribute("token")
		Attribute("lifetimeInSeconds")
	})
})

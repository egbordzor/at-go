package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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
var APIKeyHeader = Type("APIKeyHeader", func() {
	Description("Authenticating an API using username and API Key.")

	Attribute("apiKey", String, func() {
		Description("Africa’s Talking application apiKey.")
	})
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

// Authenticating an API using an Auth Token
var AuthTokenHeader = Type("AuthTokenHeader", func() {
	Description("Authenticating an API using an Auth Token")

	Attribute("authToken", String, func() {
		Description("Generated Auth Token.")
	})
	Attribute("Content-Type", String, func() {
		Description("The requests content type.")
		Enum("application/x-www-form-urlencoded", "application/xml")
	})
	Attribute("Accept", String, func() {
		Description("The requests response type.")
		Enum("application/json", "application/xml")
		Default("application/xml")
	})
	Required("authToken", "Content-Type")
})

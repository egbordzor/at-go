package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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

	// Required adds a "required" validation to the attribute.
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

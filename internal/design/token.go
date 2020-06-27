package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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

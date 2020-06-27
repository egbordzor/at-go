package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

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
		Services("africastalking")
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

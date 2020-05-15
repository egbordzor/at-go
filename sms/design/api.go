package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("messaging", func() {

	// API title.
	Title("Messaging API")

	// Description of API
	Description("HTTP service for interacting with AT's SMS API")

	// Version of API
	Version("1.0")

	// Terms of use of API
	TermsOfService("https://github.com/wondenge/atgo/blob/master/LICENSE")

	// Contact info for Author and Maintainer
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})

	// License for Library usage
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/atgo/blob/master/LICENSE")
	})

	// Library Documentation
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/atgo/blob/master/README.md")
	})

	// Server describes a single process listening for client requests.
	Server("messaging", func() {
		Description("messaging hosts the Messaging Service.")

		// List services hosted by this server.
		Services(
			"sms",
			"subscription",
			"swagger",
		)

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			URI("https://voice.sandbox.africastalking.com")
		})

		Host("production", func() {
			Description("Production hosts.")
			URI("https://api.africastalking.com")
		})
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("callbacks", func() {

	// API title.
	Title("Callback HTTP service")

	// Description of API
	Description("HTTP service for interacting with callbacks supporting AfricasTalking APIs")

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
	Server("atcallbacksvr", func() {
		Description("atcallbacksvr hosts callbacks that support AfricasTalking APIs.")

		// List services hosted by this server.
		Services("sms", "voice", "ussd", "airtime", "payments", "iot")

		// List the Hosts and their transport URLs.
		Host("local", func() {
			Description("Localhost")
			URI("http://localhost:3000/callbacks/africastalking")
		})
	})
})

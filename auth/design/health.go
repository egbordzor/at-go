package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("health", func() {

	// HTTP defines the HTTP transport specific
	// properties of an API, a service or a single method.
	HTTP(func() {
		Path("/health")
	})

	// Defines a single service method
	Method("show", func() {
		Description("Health check endpoint.")
		Result(String)
		HTTP(func() {
			GET("/")
			Response(func() {
				Code(StatusOK)
				ContentType("text/plain")
			})
		})
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var USSD = ResultType("USSD", func() {
	Description("Echo plain text response back to AT gateway")
	TypeName("USSD")
	ContentType("text/plain")

	Attribute("response", String, func() {
		Description("Plain text response back to AT gateway")
	})
})

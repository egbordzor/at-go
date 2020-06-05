package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Live: https://api.africastalking.com/version1/user
// Sandbox: https://api.sandbox.africastalking.com/version1/user
var UserMedia = ResultType("UserMedia", func() {
	Description("Get application data response.")
	TypeName("UserMedia")
	ContentType("application/json")

	Attribute("UserData", MapOf(String, String, func() {
		Key(func() {
			Description("Your Africa’s Talking application balance")
			Default("balance")
		})
		Elem(func() {
			// Validates values of the map
			Pattern("[a-zA-Z]+")

			// The format of this string is:
			// (3-digit Currency Code)(space)(Decimal Value)
			Example("KES 1785.50")
		})
	}))
})

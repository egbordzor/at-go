package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var UserResponse = ResultType("UserMedia", func() {
	Description("A User HTTP response.")
	TypeName("UserMedia")
	ContentType("application/json")

	Attributes(func() {
		Attribute("UserData", UserData)
	})
	View("default", func() {
		Attribute("UserData")
	})
})

var UserData = Type("UserData", func() {

	// The format of this string is: (3-digit Currency Code)(space)(Decimal Value)
	Attribute("balance", String, func() {
		Description("Your Africa’s Talking application balance.")
		Example("KES 1785.50")
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("user", func() {
	Title("User Service")
	Description("Initiates an application data request")

	Method("get", func() {

	})

})

// Live: https://api.africastalking.com/version1/user
// Sandbox: https://api.sandbox.africastalking.com/version1/user
// Initiate an application data request by making a HTTP GET request to the following endpoint:
var UserResponse = ResultType("UserResponse", func() {
	Attribute("UserData", MapOf(String, String, func() {

		Key(func() {
			Description("Your Africa’s Talking application balance")
			Default("balance")
		})
		Elem(func() {
			Pattern("[a-zA-Z]+")   // Validates values of the map
			Example("KES 1785.50") // The format of this string is: (3-digit Currency Code)(space)(Decimal Value)
		})
	}))
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("ussd", func() {
	Title("USSD Callback")
	Description("USSD notifications sent from Africa'sTalking gateway")

	// USSD Callback
	Method("trigger", func() {
		Description("Callback URL that sends request data our App using HTTP POST.")
		Payload(USSDPayload)
		Result(String)
		//Error()
		HTTP(func() {
			Headers(func() {
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/json")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("Content-Type")
			})
			POST("/ussd")
			Response(StatusOK) // 200 RFC 7231, 6.3.1
		})
	})

})

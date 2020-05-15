package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("voice", func() {
	Title("Voice Callback")
	Description("Voice notifications sent from Africa'sTalking gateway")
	Method("handle", func() {
		Description("Makes an outbound calls through the Africa'sTalking Voice API")
		Payload(VoiceNotificationPayload)
		Result(VoiceNotificationResult)
		//Error()
		HTTP(func() {
			Headers(func() {
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})
			POST("/voice")
			Response(StatusOK)
		})
	})

})

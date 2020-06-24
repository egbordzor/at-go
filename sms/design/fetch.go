package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var FetchMsgPayload = Type("FetchMsgPayload", func() {
	Description("Incrementally fetches our application inbox")

	Attribute("username", String, func() {
		Description("Africa’s Talking application username.")
		Default("sandbox")
	})
	Attribute("lastReceivedId", String, func() {
		Description("This is the id of the message that you last processed.")
		Default("0") // The default is 0.
	})
	Required("username")
})

var FetchMsgResponse = ResultType("FetchMsgResponse", func() {
	Description("Fetch Messages Response")
	TypeName("FetchMsgResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("SMSMessageData", FetchSMSMessageData)
	})

	View("default", func() {
		Attribute("SMSMessageData")
	})
})

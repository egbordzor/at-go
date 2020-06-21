package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Live: https://voice.africastalking.com/queueStatus
// Sandbox: https://voice.sandbox.africastalking.com/queueStatus

var QueuedCallsPayload = Type("QueuedCallsPayload", func() {
	Description("Handles more calls than you can handle at one time")
	Attribute("username", String, "Your Africa’s Talking application username.")
	Attribute("phoneNumbers", String, "List of one or more numbers mapped to your Africa’s Talking account.")
	Required("username", "phoneNumbers")
})

// Queue status response for a successful request:
var QueuedStatusResult = ResultType("QueuedStatusResult", func() {
	Attribute("Entries", ArrayOf(QueuedStatusEntry))
	Attribute("errorMessage", String, "Error Message")

})

var QueuedStatusEntry = Type("QueuedStatusEntry", func() {
	Attribute("phoneNumber", String)
	Attribute("queueName", String)
	Attribute("numCalls", String)
})

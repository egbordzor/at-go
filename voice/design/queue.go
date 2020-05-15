package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("queue", func() {
	Title("Queue Calls using Africa'sTalking Voice API")
	Description("Used when you have more calls than you can handle at one time")
	Method("add", func() {
		Payload(QueuedCallsPayload)
		Result(QueuedStatusResult)
		//Error()
		HTTP(func() {
			Headers(func() {
				Attribute("apiKey", String, "Africa’s Talking application apiKey.")
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/x-www-form-urlencoded")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("Content-Type", "Accept")
			})

			// Live: https://voice.africastalking.com/queueStatus
			// Sandbox: https://voice.sandbox.africastalking.com/queueStatus
			POST("/queueStatus")

			// You can check the HTTP Response Code to determine whether the request was successful.
			// Any response code other than 201 (Created) indicates that the call was not initiated
			Response(StatusCreated)
		})
	})
})

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

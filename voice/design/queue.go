package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Service describes a Queue Calls Service
var _ = Service("queue", func() {
	Title("Queue Calls Service")

	// Method describes a service method (endpoint)
	Method("add", func() {
		Description("Used when you have more calls than you can handle at one time")

		// Payload describes the method payload.
		Payload(QueuedCallsPayload)

		// Result describes the method result.
		Result(QueuedStatusResult)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
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

			// Requests to the service consist of HTTP POST requests.
			// Live: https://voice.africastalking.com/queueStatus
			// Sandbox: https://voice.sandbox.africastalking.com/queueStatus
			POST("/queueStatus")

			// Responses use a "201 (Created)" HTTP status.
			// Any response code other than 201 (Created)
			// indicates that the call was not initiated.
			// The result is encoded in the response body.
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

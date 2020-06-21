package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("iot", func() {

	// 6. IoT API
	Method("publish", func() {
		Description("Publishes messages to remote devices.")
		Payload(IoTPayload)
		Result(IoTResponse)
		HTTP(func() {

			// POST request to https://iot.africastalking.com/data/publish
			POST("/data/publish")
			Response(StatusOK)
		})
	})
})

// Messages can be sent between internet IoT enabled devices via MQTT or HTTP
// This IoTPayload enables interaction via HTTP.
var IoTPayload = Type("IoTPayload", func() {
	Description("IoT request payload")
	Attribute("username", String, func() {
		Description("Africa’s Talking application username")
	})
	Attribute("deviceGroup", String, func() {
		Description("Device group to which the message is to be sent")
	})
	Attribute("topic", String, func() {
		Description("Messaging channel to which the message is to be sent.")
		Example("<username>/<device-group>/<the-topic>")
	})
	Attribute("payload", String, func() {
		Description("Message packet to be sent to the subscribed devices")
	})
	Required("username", "deviceGroup", "topic", "payload")
})

var IoTResponse = ResultType("IoTResponse", func() {
	Description("An IoT HTTP response")
	TypeName("IoTResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("Response status of the API request. ")
			Enum("true", "false")
		})
		Attribute("description", String, func() {
			Description("Verbose response message detailing the status of the HTTP response")
			Enum(

				// Message was successfully published.
				"Message processed successfully",

				// The device group access level does not allow publishing.
				// The group is a Sub group
				"Publishing not allowed",

				// The topic to which the message was published
				// to is not part of the target device group.
				"The provided topic does not belong to the user",
			)
		})
	})

	View("default", func() {
		Attribute("status")
		Attribute("description")
	})
})

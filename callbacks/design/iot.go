package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// IoT Notifications
//The IoT API sends a notification when a device publish event occurs.
// To receive these notifications you need to setup a callback URL depending on the type of notification.
var IoTNotification = Type("IoTNotification", func() {
	Description("")

	Attribute("payload", String, func() {
		Description("The MQTT packet sent by the publishing device.")
		Example("42")
	})
	Attribute("topic", String, func() {
		Description("Message channel to which the message was sent by the publishing MQTT client")
		Example("myusername/devicegroup/sensor/id/1/temperature")
	})
})

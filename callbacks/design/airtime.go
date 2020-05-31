package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Airtime Validation Notifications
// The Airtime API provides optional functionality to validate airtime requests from your application.
// To receive these notifications you need to setup an airtime validation callback URL.
// From the dashboard select Airtime -> Airtime Callback URLs -> Validation Callback URL.
//
// Validation notification content
// Airtime validation notifications are sent as HTTP POST requests to the
// validation callback URL provided and contain the following parameters:
var AirtimeValidationPayload = Type("AirtimeValidationPayload", func() {
	Attribute("transactionId", String, func() {
		Description("The transaction id within Africa’s Talking.")
		Example("SomeTransactionID")
	})
	Attribute("phoneNumber", String, func() {
		Description("The phone number of the mobile subscriber receiving the airtime.")
		Example("+254711XXXYYY")
	})
	Attribute("currencyCode", String, func() {
		Description("The 3-digist ISO format currency for the value of this transaction")
		Example("KES")
	})
	Attribute("amount", UInt64, func() {
		Description("Amount - in the provided currency - that the client will receive.")
		Example(500.00)
	})
})

// Validation notification response
// Once you receive a validation callback notification you’ll be expected to
// send back a JSON response that marks the transaction as Validated or Failed.
// If validated we will proceed to send the airtime, if failed, we will block the airtime transaction.
var AirtimeValidationResponse = Type("AirtimeValidationResponse", func() {
	Attribute("status", String, func() {
		Enum("Validated", "Failed")
		Example("Validated")
	})
})

// Airtime Status Notifications
// The Airtime API sends delivery status notification from the mobile service
// provider to your application indicating success or failure of the request.
// To receive these notifications you need to setup an airtime status callback URL.
// From the dashboard select Airtime -> Airtime Callback URLs -> Status Callback URL.
//
// Status notification content
// Airtime status notifications are sent as HTTP POST requests to the status callback
// URL provided and contain the following parameters:
var AirtimeStatus = Type("AirtimeStatus", func() {
	Attribute("requestId", String, func() {
		Description("The request ID sent back as a response to the airtime send request.")
		Example("ATQid_SampleTxnId123")
	})
	Attribute("status", String, func() {
		Description("The Transaction status.")
		Enum("Success", "Failed")
		Example("Success")
	})
})

package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var MakeCallPayload = Type("MakeCallPayload", func() {
	Description("Makes an outbound call.")

	Attribute("username", String, func() {
		Description("Africa’s Talking Username")
	})
	Attribute("from", String, func() {

		// In international format i.e. +XXXYYYYYY
		Description("Africa’s Talking phone number")
	})
	Attribute("to", String, func() {
		Description("A comma separated string of recipients’ phone numbers.")
	})
	Attribute("clientRequestId", String, func() {
		Description("Variable sent to Events Callback URL used to tag the call")
	})

	Required("username", "from", "to")
})

var MakeCallResponse = ResultType("MakeCallResponse", func() {
	Description("Outbound call HTTP response.")
	TypeName("MakeCallResponse")
	ContentType("application/json")

	Attributes(func() {
		// A list with multiple Entry each corresponding
		// to an individual phone number and their status.
		// Entry is a Map with details of queued numbers.
		Attribute("entries", ArrayOf(VoiceEntry))
		Attribute("errorMessage", String, func() {
			Description("Error message if ENTIRE request was rejected by the API.")
		})

	})
	View("default", func() {
		Attribute("entries")
		Attribute("errorMessage")
	})
})

var VoiceEntry = Type("VoiceEntry", func() {
	Attribute("phoneNumber", String, func() {
		Description("The phone number queued.")
	})
	Attribute("status", String, func() {
		Enum(
			// The call request has been accepted and queued
			"Queued",

			// Recipient number is in an incorrect format
			"InvalidPhoneNumber",

			// Recipient number is outside the supported zone
			"DestinationNotSupported",

			// Your AfricasTalking account has insufficient balance.
			"InsufficientCredit",
		)
	})
	Attribute("sessionId", String, func() {
		Description("A unique id for the request associated to this phone number")

		// Defaults to None if an error occurred.
		Default("None")
	})
})

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Live: https://voice.africastalking.com/callTransfer
// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
var CallTransferPayload = Type("CallTransferPayload", func() {
	Description("Transfer calls to another number")

	Attribute("sessionId", String, func() {
		Description("Session Id of the ongoing call, it must have two legs")
	})
	Attribute("phoneNumber", String, func() {
		Description("Phone Number to transfer the call to")
	})
	Attribute("callLeg", String, func() {
		Description("Call leg to transfer the call to either caller or callee")
		Enum("caller", "callee")
		Default("callee") // (Defaults to callee)
	})
	Attribute("holdMusicUrl", String, func() {
		Description("The url of the media file to be played when the user is on hold.") // Don’t forget to start with http://
	})

	Required("sessionId", "phoneNumber")
})

var CallTransferResponse = ResultType("CallTransferResponse", func() {
	Description("Call Transfer HTTP response.")
	TypeName("CallTransferResponse")
	ContentType("application/json")

	Attributes(func() {
		Attribute("status", String, func() {
			Description("can be either Success or Aborted")
			Enum("Success", "Aborted")
		})
		Attribute("errorMessage", String, func() {
			Description("Why the transfer ws aborted None is successful")
		})
	})
})

var QueuedCallsPayload = Type("QueuedCallsPayload", func() {
	Description("Handles more calls than you can handle at one time")

	Attribute("username", String, func() {
		Description("Your Africa’s Talking application username.")
	})
	Attribute("phoneNumbers", String, func() {
		Description("List of one or more numbers mapped to your Africa’s Talking account.")
	})

	Required("username", "phoneNumbers")
})

// Queue status response for a successful request:
var QueuedStatusResult = ResultType("QueuedStatusResult", func() {
	Description("Queued Status HTTP response result.")
	TypeName("QueuedStatusResult")
	ContentType("application/json")

	Attributes(func() {
		Attribute("Entries", ArrayOf(QueuedStatusEntry))
		Attribute("errorMessage", String, func() {
			Description("Error Message")
		})
	})

	View("default", func() {
		Attribute("Entries")
		Attribute("errorMessage")
	})
})

var QueuedStatusEntry = Type("QueuedStatusEntry", func() {
	Attribute("phoneNumber", String)
	Attribute("queueName", String)
	Attribute("numCalls", String)
})

var UploadMediaFile = Type("UploadMediaFile", func() {
	Description("Uploads media or audio files to Africa'sTalking servers")

	Attribute("username", String, func() {
		Description("Your Africa’s Talking application username.")
	})
	// This contains the audio file you want played during a call.
	Attribute("url", String, func() {
		// Don’t forget to start with http://
		Description("The url of the file to upload.")
	})
})

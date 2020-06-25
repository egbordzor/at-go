package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Content-Type: application/x-www-form-urlencoded or multipart/form-data
// Live: https://voice.africastalking.com/call
// Sandbox: https://voice.sandbox.africastalking.com/call

// Make an outbound call
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
	ContentType("application/xml")

	// A list with multiple Entry each corresponding
	// to an individual phone number and their status.
	// Entry is a Map with details of queued numbers.
	Attribute("entries", ArrayOf(VoiceEntry))
	Attribute("errorMessage", String, func() {
		Description("Error message if ENTIRE request was rejected by the API.")
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
	Attribute("status", String, func() {
		Description("can be either Success or Aborted")
		Enum("Success", "Aborted")
	})
	Attribute("errorMessage", String, "Why the transfer ws aborted None is successful")
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

// Sends a HTTP POST request to Africa'sTalking Voice API.
// Live: https://voice.africastalking.com/mediaUpload
// Sandbox: https://voice.sandbox.africastalking.com/mediaUpload
// Any response code other than 201 (Created) indicates that the call was not initiated.
var UploadMediaFile = Type("UploadMediaFile", func() {
	// With the extension .mp3 or .wav only
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

// One can send AT some text (in English) to read to the user.
// E.g Let us say we want to tell the user what their account balance is.
var Say = Type("Say", func() {

	// The possible values are man
	// or woman (default: woman).
	// Note that the voice will
	// read the text in English.
	Attribute("voice", String, func() {
		Description("This parameter specifies the voice to use")
		Enum("man", "woman")
		Default("woman")
		Example("woman")
	})
	Attribute("playBeep", Boolean, func() {
		Description("Instructs AT to play a beep after reading the text contained in the request")
		// The possible values
		// are true and false
		// (default: false)
		Default(false)
	})
})

var Play = Type("Play ", func() {
	Description("Plays back an audio file that is located anywhere on the web.")

	// Note that we will cache the content after
	// the first playback request, and will check
	// for request header information that indicates
	// the need to re-download the file later on.
	Attribute("url", String, func() {
		Description("A valid URL that contains a link to the file to be played.")
	})
	Required("url")
})

// This action to get the digits that a user enters
// on their phone in response to a prompt from your
// application. For example, you can ask the user to
// enter their account number, followed by the hash
// sign, and then read back their balance to them.
// Once we get this response, our server will read out
// the text contained in the Say request to the caller
// and wait for the configured timeout for the response.
// If the user successfully completed the request and
// presses the hash sign, our API will forward this
// response to your web server and fetch the next set
// of instructions.
// Otherwise, if a GetDigits action receives no input
// at all, it will move on to the next action without
// triggering a notification. This means that if it is
// the last action the call will be terminated if the
// user does not include any input. If you would like us
// to say something before hanging up, you can add
// another element after the GetDigits request
var GetDigits = Type("GetDigits", func() {

	// If absent, our API will forward the request
	// to the default URL for this phone number
	// (or to the redirected URL if a redirect has
	// been issued). Default: None
	Attribute("callbackUrl", String, func() {
		Description("Instructs AT to forward results of the GetDigits action to the URL value passed in. ")
	})

	// Default: None
	Attribute("numDigits", String, func() {
		Description("This shows the number of digits you would like to grab from the user input.")
	})

	// If there is no other element to execute, the system will hang up. Default: 30 Seconds
	Attribute("timeout", String, func() {
		Description("Timeout (in seconds) for getting the digits, after which the system moves on to the next element.")
	})

	// Default: None
	Attribute("finishOnKey", String, func() {
		Description("The key which will terminate the action of getting digits.")
	})
})

// You can use this element to connect the user who
// called your phone number to an external phone number.
// For example, if you use your Africa’s Talking phone
// number for customer service queries, you can forward
// calls received on that phone number to an actual phone
// line located in your office. No action will be executed
// after this if the call is picked.
var Dial = Type("Dial", func() {

	// A comma separated list of phone
	// numbers (in international format)
	// to call.
	Attribute("phoneNumbers")

	// Record the conversation (true or false).
	// The URL for the recording will
	// be sent to you once the call
	// is complete. Default: false
	Attribute("record")

	// If more than one phone number is
	// provided, this determines whether the
	// phone numbers will be dialed one after
	// the other, or whether they will all ring
	// at the same time. If set to true, the
	// numbers will be dialed one after the other.
	// If set to false, they will ring at the same
	// time and will be handled in the order of
	// who picks up first. Default: false.
	Attribute("sequential")

	// This contains the Africa’s Talking number
	// you want to dial out with. It is mainly
	// important when you call using a sip number.
	// If not specified, the number called by the
	// user will be used.
	Attribute("callerId")

	// This contains a URL location of a media
	// playback you would want the user to listen
	// to when the call has been placed before
	// its picked up.
	Attribute("ringBackTone")
	Attribute("maxDuration", String, func() {
		Description("This contains the maximum amount of time in seconds a call should take.")
	})
	Required("phoneNumbers")
})

//  This element lets you record a call session
//  into an mp3 file that you can then retrieve
//  and play later. Our API supports partial
//  recording or terminal recording. Partial recording
//  is useful where a particular user response is
//  required like while prompting for a name. For
//  terminal recording, the recording starts when
//  the record action is called until the mobile
//  user hangs up.
//
//  1. Partial Recording: A recording stops when
//  finishOnKey button is pressed or maxLength time
//  is reached or the user hangs up. When the
//  recording ends, the URL of the recorded file will
//  be sent in the final response from our APIs in a
//  field named recordingUrl. If no file was recorded,
//  this variable will contain an empty string. The
//  call proceeds based on the response to that notification.
//
//  2. Terminal Recording: Terminal recording doesn’t
//  have any attributes. No other action will be executed
//  after. The URL of the recorded file will be sent in the
//  final response from our APIs in a field named recordingUrl.
//  If no file was recorded, this variable will contain an
//  empty string.
var Record = Type("Record", func() {

	// This shows the number of digits
	// you would like to grab from the
	// user input. Default: None
	Attribute("finishOnKey")

	// This specifies the maximum amount
	// of time in seconds a recording
	// should take. Default: None
	Attribute("maxLength")

	// This specifies how long to wait
	// until ending the record session after
	// silence is detected. The record session
	// will be ended after time has elapsed.
	// Default: 3600 seconds.
	Attribute("timeout")

	// This is a boolean attribute which specifies
	// whether you want to remove the final parts
	// of a recording where the user was silent.
	// Default: false
	Attribute("trimSilence")

	// This is a boolean attribute which specifies
	// whether the API should play a beep when the
	// recording starts.
	// Default: false
	Attribute("playBeep")

	// This is the URL of the script to process the
	// recording mp3 file sent by the API. If absent,
	// our API will forward the request to the default
	// URL for this phone number.
	// Default: None
	Attribute("callbackUrl")
})

// Enqueuing is passing an incoming call to a queue
// to be handled later.
var Enqueue = Type("Enqueue", func() {

	// A valid URL that contains a link to the file
	// to be played while the user is on hold. Note
	// that we will cache the content after the first
	// playback request, and will check for request
	// header information that indicates the need to
	// re-download the file later on.
	// Default: None
	Attribute("holdMusic")

	// You can put a call in a particular queue.
	// eg. support, general, technical. This would
	// help you dequeue the call to a particular
	// person or team.
	// Default: None
	Attribute("name")
})

// Dequeuing is passing the calls enqueued to a separate number so
// that it can be handled e.g by an agent.
var Dequeue = Type("Dequeue", func() {

	// The Phone number which users called to join the queue.
	// Our gateway will check to see whether there is anyone in
	// the queue that corresponds to that phone number and connect
	// the two calls. If there is no one in the queue, the gateway
	// just hangs up.
	Attribute("phoneNumber")

	// This is the name of the queue you want to dequeue from.
	// Default: None
	Attribute("name")
})

// This action will transfer control of the call to the script whose
// URL is passed in. This can help you better organize your call
// handling logic by spreading the logic across multiple scripts.
// Any action after Redirect won’t execute since the API will be
// handling events from the script it was redirected to. Once we
// get this response, our server will generate a POST request to
// the specified URL and proceed with the call based on the returned
// XML. The new url will be used as the callback url for the rest of
// the call.
var Redirect = Type("Redirect", func() {
	Attribute("Reject", String, "Reject")
})

// This action lets you reject an incoming call without incurring
// any usage charges from AT's APIs. Note that this should be the
// only element in your response.
var Reject = Type("Reject", func() {
	Attribute("Redirect", String, "Redirect")
})

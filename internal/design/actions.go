package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// One can send AT some text (in English) to read to the user.
// E.g Let us say we want to tell the user what their account balance is.
var SayPayload = Type("SayPayload", func() {

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

var PlayPayload = Type("PlayPayload", func() {
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
var GetDigitsPayload = Type("GetDigitsPayload", func() {

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
var DialPayload = Type("DialPayload", func() {

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
var RecordPayload = Type("RecordPayload", func() {

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
var EnqueuePayload = Type("EnqueuePayload", func() {

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
var DequeuePayload = Type("DequeuePayload", func() {

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
var RedirectPayload = Type("RedirectPayload", func() {
	Attribute("Reject", String, "Reject")
})

// This action lets you reject an incoming call without incurring
// any usage charges from AT's APIs. Note that this should be the
// only element in your response.
var RejectPayload = Type("RejectPayload", func() {
	Attribute("Redirect", String, "Redirect")
})

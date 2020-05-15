package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)


// Making an API call
// You need to include the API key in the request header as a field called apiKey.
//
// The place where the username should be included depends on the type of request.
// For GET requests: Username should be passed as a query parameter.
// For POST requests:  URL encoded form parameters, username included as one of the parameters within the form.
// For POST requests: Require JSON in the request body, Username included in the JSON sent in the body of the request.

// Authenticating using your API Key and Username
var APIKeyHeader = Type("APIKeyHeader", func() {
	Attribute("apiKey", String, "Africa’s Talking application apiKey.")
	Attribute("Content-Type", String, func() {
		Description("The requests content type.")
		Enum("application/x-www-form-urlencoded", "application/json")
	})
	Attribute("Accept", String, func() {
		Description("The requests response type.")
		Enum("application/json", "application/xml")
		Default("application/xml")
	})
	Required("apiKey", "Content-Type")
})

// Authenticating with an Auth Token
var AuthTokenHeader = Type("AuthTokenHeader", func() {
	Attribute("authToken", String, "Generated Auth Token.")
	Attribute("Content-Type", String, func() {
		Description("The requests content type.")
		Enum("application/x-www-form-urlencoded", "application/xml")
	})
	Attribute("Accept", String, func() {
		Description("The requests response type.")
		Enum("application/json", "application/xml")
		Default("application/xml")
	})

})

var TokenPayload = Type("TokenPayload", func() {
	Attribute("username", String, "AT Username")
	Attribute("apiKey", String, "API Key ")

	Required("username", "username")
})

var TokenResponse = ResultType("TokenResponse", func() {
	Attribute("token", String, "Generated Auth Token.")
	Attribute("lifetimeInSeconds", Int, "Token Lifetime")

	Required("token", "lifetimeInSeconds")

	View("default", func() {
		Attribute("token")
		Attribute("lifetimeInSeconds")
	})
})

// Live: https://api.africastalking.com/version1/user
// Sandbox: https://api.sandbox.africastalking.com/version1/user
// Initiate an application data request by making a HTTP GET request to the following endpoint:
var UserResponse = ResultType("UserResponse", func() {
	Attribute("UserData", MapOf(String, String, func() {

		Key(func() {
			Description("Your Africa’s Talking application balance")
			Default("balance")
		})
		Elem(func() {
			Pattern("[a-zA-Z]+")   // Validates values of the map
			Example("KES 1785.50") // The format of this string is: (3-digit Currency Code)(space)(Decimal Value)
		})
	}))
})

// Live: https://api.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
// Send SMS through your application by making a HTTP POST request to the above endpoints.
var SendSMS = Type("SendSMS", func() {
	Attribute("username", String, "Africa’s Talking application username.")
	Attribute("to", String, "Recipients’ phone numbers.")
	Attribute("message", String, "Message to be sent.")
	Attribute("from", String, "Registered Short Code or Alphanumeric.")

	// This is used by the Mobile Service Provider to determine who
	// gets billed for a message sent using a Mobile-Terminated ShortCode.
	// This parameter will be ignored for messages sent using alphanumerics
	// or Mobile-Originated shortcodes.
	Attribute("bulkSMSMode", Int, func() {
		Description("Used by MSP to determine who gets billed for a message")

		// The value must be set to 1 for bulk messages.
		Enum(1)

		// The default value is 1(which means that the sender
		// - Africa’s Talking account being used - gets charged).
		Default(1)
	})

	// This is used for Bulk SMS clients that would like to deliver
	// as many messages to the API before waiting for an
	// acknowledgement from the Telcos. If enabled, the API will store
	// the messages in a queue and send them out asynchronously after responding
	// to the request. The default value is 1
	Attribute("enqueue", Int, func() {
		Description("Used for Bulk SMS clients")

		// Possible values are 1 to enable and 0 to disable.
		Enum(0, 1)
		Default(1)
	})
	Attribute("keyword", String, "The keyword to be used for a premium service.")
	// This is used for premium services to send OnDemand messages. AT forwards
	// the linkId to our application when the user sends a message to our service.
	Attribute("linkId")
	Attribute("retryDurationInHours", String, func() {
		Description("No. of hours subscription message should be retried in case it’s not delivered to the subscriber.")
	})
	Required("username", "to", "message")
})

// Live: https://api.africastalking.com/version1/messaging
// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
// To do so, you make a HTTP GET request to the above endpoints.
// Incrementally fetches our application inbox.
var FetchMessages = Type("FetchMessages", func() {
	Attribute("username", String, "Africa’s Talking application username.")
	Attribute("lastReceivedId", String, func() {
		Description("This is the id of the message that you last processed.")

		// The default is 0.
		Default("0")
	})
})

var USSD = ResultType("USSD", func() {
	Description("Echo plain text response back to AT gateway")
	TypeName("USSD")
	ContentType("text/plain")

	Attribute("response", String, func() {
		Description("Plain text response back to AT gateway")
	})
})

// Making Calls: Send a POST request to Africa'sTalking Voice API.
//
// When the call is picked, Africa'sTalking sends a notification
// to our callback url so that our application can let Africa'sTalking
// know how to handle the call.
//
// The next action on such calls will depend on the kind of response
// Africa'sTalking gets from the notification sent from our application.
// If the action. requires a user’s input, Africa'sTalking will send
// another notification to our callback url to submit the input and
// find out the next action.
//
// This continues until either the user ends the call or we respond with
// an action that terminates the call.
//
// Once the call is terminated, Africa'sTalking sends a final notification
// to our callback url with some extra information about the call such
// as cost of the call and its duration.

// Content-Type: The requests content type. Can be application/x-www-form-urlencoded or multipart/form-data

// MakeCall makes an outbound call through the Africa'sTalking  Voice API
// by sending HTTP POST request to one of the following endpoints.
// Live: https://voice.africastalking.com/call
// Sandbox: https://voice.sandbox.africastalking.com/call
var MakeCallPayload = Type("MakeCallPayload", func() {
	Attribute("username", String, "Africa’s Talking Username")
	Attribute("from", String, "Africa’s Talking phone number") // In international format i.e. +XXXYYYYYY
	Attribute("to", String, "A comma separated string of recipients’ phone numbers.")
	Attribute("clientRequestId", String, "Variable sent to Events Callback URL used to tag the call")

	Required("username", "from", "to")
})

var MakeCallResult = ResultType("MakeCallResult", func() {
	Description("Voice call response")
	TypeName("MakeCallResult")
	ContentType("application/xml")

	// A list with multiple Entry each corresponding
	// to an individual phone number and their status.
	// Entry is a Map with details of queued numbers.
	Attribute("entries", ArrayOf(Entry))
	Attribute("errorMessage", String, "Error message if ENTIRE request was rejected by the API.")

	View("default", func() {
		Attribute("entries")
		Attribute("errorMessage")
	})
})

var Entry = Type("Entry", func() {
	Attribute("phoneNumber", String, "The phone number queued.")
	Attribute("status", String, func() {
		Enum(
			"Queued",                  // Queued: The call request has been accepted and queued
			"InvalidPhoneNumber",      //InvalidPhoneNumber: Recipient number is in an incorrect format
			"DestinationNotSupported", //DestinationNotSupported: Recipient number is outside the supported zone
			"InsufficientCredit",      //InsufficientCredit: Your AfricasTalking account has insufficient balance.
		)
	})
	// Defaults to None if an error occurred.
	Attribute("sessionId", String, "A unique id for the request associated to this phone number")
})

// Steps Involved In Handling A Call Session
//
// 1. Africa'sTalking receives a call to our phone number on their voice gateways,
// or we successfully initiate a call using their calling API.
//
// 2. Africa'sTalking Voice API sends a POST request to the notification callback URL
// that we have set for that phone number in our Voice Dashboard.
//
// 3. Our application responds with XML that tells Africa'sTalking Voice API how
// to handle the call. This XML will typically contain a list of actions
// that Africa'sTalking API will execute in sequence.
//
// 4. Our API translates those actions into events or messages relayed back to the caller.

// HandleCall handles calls made to our Africa’s Talking phone number.
var HandleCall = ResultType("HandleCall", func() {

	// The API will set a value of 0 in the final request to your application.
	// That request will contain details about the call’s duration and cost.
	Attribute("isActive", String, "Know whether the call is in session state.")

	// This variable will stay the same throughout the call
	Attribute("sessionId", String, "A unique identifier that we will generate during each call session.")

	// Inbound calls are initiated by a phone user.
	// Outbound calls are initiated by your application.
	Attribute("direction", String, "Whether this is an inbound or outbound call.")
	Attribute("callerNumber", String, "The phone number of the phone user in the call.")
	Attribute("destinationNumber", String, "Your Africa’s Talking phone number.")
	Attribute("dtmfDigits", String, "Digits a user enters in response to a getDigits request")

	// (using either the Record element, or the record attribute of the Bridge element)
	Attribute("recordingUrl", String, "The URL of the recording made for this call")
	Attribute("durationInSeconds", String, "The duration of the call in seconds.")
	Attribute("currencyCode", String, "The currency used to bill this call")
	Attribute("amount", String, "The total cost of the call.")

	View("AllRequests", func() {
		Attribute("isActive")
		Attribute("sessionId")
		Attribute("direction")
		Attribute("callerNumber")
		Attribute("destinationNumber")
	})

	View("GetDigits", func() {
		Attribute("isActive")
		Attribute("sessionId")
		Attribute("direction")
		Attribute("callerNumber")
		Attribute("destinationNumber")
		Attribute("dtmfDigits")

	})

	View("FinalRequest", func() {
		Attribute("isActive")
		Attribute("sessionId")
		Attribute("direction")
		Attribute("callerNumber")
		Attribute("destinationNumber")
		Attribute("recordingUrl")
		Attribute("durationInSeconds")
		Attribute("currencyCode")
		Attribute("amount")
	})

})

// Call Transfer
// You can transfer your call to another number by making a HTTP POST request to one of the following endpoints:
//
// Endpoints
// Live: https://voice.africastalking.com/callTransfer
// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
var CallTransferPayload = Type("CallTransferPayload", func() {
	Attribute("sessionId", String, "Session Id of the ongoing call, it must have two legs")
	Attribute("phoneNumber", String, "Phone Number to transfer the call to")
	Attribute("callLeg", String, func() {
		Description("Call leg to transfer the call to either caller or callee")
		Enum("caller", "callee")
		Default("callee") // (Defaults to callee)
	})
	// Don’t forget to start with http://
	Attribute("holdMusicUrl", String, "The url of the media file to be played when the user is on hold.")
	Required("sessionId", "phoneNumber")
})

var CallTransferResult = ResultType("CallTransferResult", func() {
	Attribute("status", String, func() {
		Description("can be either Success or Aborted")
		Enum("Success", "Aborted")
	})
	Attribute("errorMessage", String, "Why the transfer ws aborted None is successful")
})

// Event Notifications
// When the transfer has been initiated we will send any of these events to your
// event notifications URL, you can check from these form fields in the events
var EventPayload = ResultType("EventPayload", func() {
	Attribute("callTransferParam", String, "+2347XXXXXXXXX:20, (20 is the duration in seconds)")
	Attribute("status", String, func() {
		Enum("Success")
	})
	Attribute("callSessionState", String, func() {
		Enum("Active", "Transferred", "TransferCompleted")
	})
	Attribute("isActive", String, func() {
		Enum("0", "1")
		Default("1")
	})
	Attribute("callTransferredToNumber", String, "Number call was transferred to")
	Attribute("callTransferState", String, func() {
		Enum(" Active", "Completed", " CallerHangup", "CalleeHangup")
	})
	Attribute("callTransferHangupCause", String, func() {
		Enum("DestinationNotSupported", "InvalidPhoneNumber", "NoActiveClient", "NotAllowed")
	})
})

var EventResult = ResultType("EventResult", func() {
	Description("Event notifications sent to URL when call transfer has been initiated.")
	TypeName("EventResult")
	ContentType("application/xml")

	Attributes(func() {
		Extend(EventPayload)
	})
	Attribute("callTransferParam", String, "+2347XXXXXXXXX:20, (20 is the duration in seconds)")
	Attribute("status", String, func() {
		Enum("Success")
	})
	Attribute("callSessionState", String, func() {
		Enum("Active", "Transferred", "TransferCompleted")
	})
	Attribute("isActive", String, func() {
		Enum("0", "1")
		Default("1")
	})
	Attribute("callTransferredToNumber", String, "Number call was transferred to")
	Attribute("callTransferState", String, func() {
		Enum(" Active", "Completed", " CallerHangup", "CalleeHangup")
	})
	Attribute("callTransferHangupCause", String, func() {
		Enum("DestinationNotSupported", "InvalidPhoneNumber", "NoActiveClient", "NotAllowed")
	})

	// Call transferred but failed
	View("CallTransferredButFailed", func() {
		Attribute("callSessionState")
		Attribute("isActive")
		Attribute("callTransferHangupCause")
	})

	// Call successfully transferred
	View("CallTransferredSuccessfully", func() {
		Attribute("callSessionState")
		Attribute("isActive")
		Attribute("callTransferredToNumber")
		Attribute("callTransferState")
	})

	// Call transfer ends
	View("CallTransferEnds", func() {
		Attribute("callSessionState")
		Attribute("isActive")
		Attribute("callTransferredToNumber")
		Attribute("callTransferState")
	})

	// Caller hang up
	View("CallerHangup", func() {
		Attribute("callSessionState")
		Attribute("isActive")
		Attribute("callTransferState")
	})

	// Callee hangup
	View("CalleeHangup", func() {
		Attribute("callSessionState")
		Attribute("isActive")
		Attribute("callTransferState")
	})

	// End of call
	View("CallEnd", func() {
		Attribute("callTransferParam")
		Attribute("status")
		Attribute("callSessionState")
		Attribute("isActive")
	})
})

// Voice Actions
// A voice action is what your application sends in response to a notification
// that tells the Voice API how to handle the current call.
//
// Say: With this action, you set a text and we will read it out to the caller.
// Play: This element lets you play back an audio file that is located anywhere on the web.
// Get Digits: You can use this element to get the digits that a user enters on their phone
// in response to a prompt from your application.
// Dial: You can use this action to connect the user who called your phone number to an
// external phone number.
// Record: This element lets you record a call session into an mp3 file that you can
// then retrieve and play later. Our API supports terminal recording or partial recording.
// For final recording, the recording starts when the record action is called until the
// mobile user hangs up. Partial recording may be useful where a particular user response
// is required like while prompting for a name.
// Enqueue: Lets you pass an incoming call to a queue to be handled later.
// Dequeue: Lets you pass the calls enqueued to a separate number so that it can be handled e.g by an agent.
// Redirect: This action will transfer control of the call to the script whose URL is passed in.
// This can help you better organize your call handling logic by spreading the logic across multiple scripts.
// Reject: This action lets you reject an incoming call without incurring any usage charges from our APIs.

// Say: You can send us some text (in English) and we will read it out to the user.
// For example, let us say we want to tell the user what their account balance is.
// Your web server can respond to our request with the following XML:
var Say = Type("Say", func() {
	// The possible values are man or woman (default: woman).
	// Note that the voice will read the text in English.
	Attribute("voice", String, func() {
		Description("This parameter specifies the voice to use")
		Enum("man", "woman")
		Default("woman")
		Example("woman")
	})
	Attribute("playBeep", Boolean, func() {
		Description("This parameter instructs our API to play a beep after reading the text contained in the request")
		Default(false) // The possible values are true and false (default: false)
	})
})

// Play: This action lets you play back an audio file that is located anywhere on the web.
var Play = Type("Play ", func() {
	// Note that we will cache the content after the first playback request,
	// and will check for request header information that indicates the
	// need to re-download the file later on.
	Attribute("url", String, "A valid URL that contains a link to the file to be played.")
	Required("url")
})

// Get Digits: You can use this action to get the digits that a user enters on their phone
// in response to a prompt from your application. For example, you can ask the user to enter
// their account number, followed by the hash sign, and then read back their balance to them.
// Once we get this response, our server will read out the text contained in the Say request
// to the caller and wait for the configured timeout for the response. If the user successfully
// completed the request and presses the hash sign, our API will forward this response to your
// web server and fetch the next set of instructions.
//
// Otherwise, if a GetDigits action receives no input at all, it will move on to the next
// action without triggering a notification. This means that if it is the last action the
// call will be terminated if the user does not include any input. If you would like us
// to say something before hanging up, you can add another element after the GetDigits request
var GetDigits = Type("GetDigits", func() {

	// If absent, our API will forward the request to the default URL for this phone number
	// (or to the redirected URL if a redirect has been issued). Default: None
	Attribute("callbackUrl", String, "Instructs AT to forward results of the GetDigits action to the URL value passed in. ")

	// Default: None
	Attribute("numDigits", String, "This shows the number of digits you would like to grab from the user input.")

	// If there is no other element to execute, the system will hang up. Default: 30 Seconds
	Attribute("timeout", String, "Timeout (in seconds) for getting the digits, after which the system moves on to the next element.")

	// Default: None
	Attribute("finishOnKey", String, "The key which will terminate the action of getting digits.")
})

// Dial: You can use this element to connect the user who called your phone number
// to an external phone number. For example, if you use your Africa’s Talking phone number
// for customer service queries, you can forward calls received on that phone number to an
// actual phone line located in your office. No action will be executed after this if the call is picked.
var Dial = Type("Dial", func() {

	// A comma separated list of phone numbers (in international format) to call.
	Attribute("phoneNumbers")

	// Record the conversation (true or false).
	// The URL for the recording will be sent to you once the call is complete.
	// Default: false
	Attribute("record")

	// If more than one phone number is provided, this determines whether the phone numbers will
	// be dialed one after the other, or whether they will all ring at the same time.
	// If set to true, the numbers will be dialed one after the other.If set to false,
	// they will ring at the same time and will be handled in the order of who picks up first.
	// Default: false.
	Attribute("sequential")

	// This contains the Africa’s Talking number you want to dial out with.
	// It is mainly important when you call using a sip number.
	// If not specified, the number called by the user will be used.
	Attribute("callerId")

	// This contains a URL location of a media playback you would want
	// the user to listen to when the call has been placed before its picked up.
	Attribute("ringBackTone")
	Attribute("maxDuration", String, "This contains the maximum amount of time in seconds a call should take.")
	Required("phoneNumbers")
})

//  Record: This element lets you record a call session into an mp3 file that you
//  can then retrieve and play later. Our API supports partial recording or terminal
//  recording. Partial recording is useful where a particular user response is
//  required like while prompting for a name. For terminal recording, the recording
//  starts when the record action is called until the mobile user hangs up.
//  1. Partial Recording: A recording stops when finishOnKey button is pressed
//  or maxLength time is reached or the user hangs up.  When the recording ends,
//  the URL of the recorded file will be sent in the final response from our APIs
//  in a field named recordingUrl. If no file was recorded, this variable will contain
//  an empty string. The call proceeds based on the response to that notification.
//  2. Terminal Recording: Terminal recording doesn’t have any attributes.
//  No other action will be executed after. The URL of the recorded file will be sent
//  in the final response from our APIs in a field named recordingUrl.
//  If no file was recorded, this variable will contain an empty string.
var Record = Type("Record", func() {

	// This shows the number of digits you would like to grab from the user input.
	// Default: None
	Attribute("finishOnKey")

	// This specifies the maximum amount of time in seconds a recording should take.
	// Default: None
	Attribute("maxLength")

	// This specifies how long to wait until ending the record session after silence is detected.
	// The record session will be ended after time has elapsed.
	// Default: 3600 seconds.
	Attribute("timeout")

	// This is a boolean attribute which specifies whether you want to remove the final
	// parts of a recording where the user was silent.
	// Default: false
	Attribute("trimSilence")

	// This is a boolean attribute which specifies whether the API should play a beep
	// when the recording starts.
	// Default: false
	Attribute("playBeep")

	// This is the URL of the script to process the recording mp3 file sent by the API.
	// If absent, our API will forward the request to the default URL for this phone number.
	// Default: None
	Attribute("callbackUrl")
})

// Enqueue: Enqueuing is passing an incoming call to a queue to be handled later.
var Enqueue = Type("Enqueue", func() {

	// A valid URL that contains a link to the file to be played while the user is on hold.
	// Note that we will cache the content after the first playback request, and will check
	// for request header information that indicates the need to re-download the file later on.
	// Default: None
	Attribute("holdMusic")

	// You can put a call in a particular queue. eg. support, general, technical.
	// This would help you dequeue the call to a particular person or team.
	// Default: None
	Attribute("name")
})

// Dequeue: Dequeuing is passing the calls enqueued to a separate number so
// that it can be handled e.g by an agent.
var Dequeue = Type("Dequeue", func() {
	// The Phone number which users called to join the queue.
	// Our gateway will check to see whether there is anyone in the queue
	// that corresponds to that phone number and connect the two calls.
	// If there is no one in the queue, the gateway just hangs up.
	Attribute("phoneNumber")

	// This is the name of the queue you want to dequeue from.
	// Default: None
	Attribute("name")
})

// Redirect: This action will transfer control of the call to the script whose
// URL is passed in. This can help you better organize your call handling logic
// by spreading the logic across multiple scripts. Any action after Redirect won’t
// execute since the API will be handling events from the script it was redirected to.
// Once we get this response, our server will generate a POST request to the specified
// URL and proceed with the call based on the returned XML. The new url will be used
// as the callback url for the rest of the call.
var Redirect = Type("Redirect", func() {
	Attribute("Reject", String, "Reject")
})

// Reject: This action lets you reject an incoming call without incurring any
// usage charges from our APIs.
// Note that this should be the only element in your response.
var Reject = Type("Reject", func() {
	Attribute("Redirect", String, "Redirect")
})

// Queued Calls
// Queuing is a feature mainly used when you have more calls than you can handle at one time.
// Incoming calls will be put in a queue and handled one by one until all of them are out of the queue (dequeued).
// You can find the number of queued calls by sending a HTTP POST request to one of the following endpoints:
//
// Endpoints
// Live: https://voice.africastalking.com/queueStatus
// Sandbox: https://voice.sandbox.africastalking.com/queueStatus
var QueuedCallsPayload = Type("QueuedCallsPayload", func() {
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

// Upload Media File
//
// This API uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav only.
// Media files are played when called upon by one of Africa'sTalking voice actions.
//
// Play: "url" - This contains the audio file you want played during a call.
// Call queueing: "holdMusic" - This contains the audio file you want played when the user has been queued waiting to be dequeued.
// Dial: "holdMusic" - This contains the audio file you want played when a number has been dialed before it’s picked.
//
// Live: https://voice.africastalking.com/mediaUpload
// Sandbox: https://voice.sandbox.africastalking.com/mediaUpload
// You can check the HTTP Response Code to determine whether the request was successful.
// Any response code other than 201 (Created) indicates that the call was not initiated.
// You can upload media by sending a HTTP POST request to one of the following endpoints:
var UploadMediaFile = Type("UploadMediaFile", func() {
	Attribute("username", String, func() {
		Description("Your Africa’s Talking application username.")
	})
	Attribute("url", String, func() {
		// Don’t forget to start with http://
		Description("The url of the file to upload.")
	})
})

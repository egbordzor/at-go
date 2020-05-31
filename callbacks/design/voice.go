package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Voice Notifications
// The Voice API sends a notification when a specific event happens.
// To receive these notifications you need to setup a voice callback URL.
// From the dashboard select Voice -> Phone Numbers -> Actions -> Callback.
//
// Voice API notifications are sent for various categories as shown below:
// Outbound calls: Sent whenever you make a call from a registered SIP number.
// Inbound calls: Sent when a call comes to your virtual or SIP number.
// After input: Sent whenever an action in your response requires user input
// (such as GetDigits and Record)
// When call ends: Sent after a call ends.
// This is the final notification and contains some extra information
// about the call like the cost and duration.
var VoiceNotification = Type("VoiceNotification", func() {

	// The API will set a value of 0 in the
	// final request to your application.
	// That request will contain details
	// about the call’s duration and cost.
	Attribute("isActive", String, func() {
		Description("Lets us know whether the call is in session state")
		Default("0")
	})
	// This variable will stay the same
	// throughout the call
	Attribute("sessionId", String, func() {
		Description("A unique identifier generated during each call session")
	})
	// Inbound calls are initiated by a
	// phone user. Outbound calls are
	// initiated by our application.
	Attribute("direction", String, func() {
		Description("Whether this is an inbound or outbound call")
	})
	Attribute("destinationNumber", String, func() {
		Description("Africa’s Talking phone number, in international format")
		Example("+254711XXXYYY")
	})
	Attribute("callerNumber", String, func() {
		Description("The phone number of the phone user in the call, in international format.")
		Example("+254711XXXYYY")
	})
	Attribute("callerCountryCode", String, func() {
		Description("The code of the country the callerNumber is calling from.")
	})
	Attribute("callStartTime", String, func() {
		Description("The time the call began.")
	})
	// Only present in a notification  following a
	// GetDigits response.
	Attribute("dtmfDigits", String, func() {
		Description("The digits that a user enters in response to a getDigits request")
	})
	// The URL of the recording made for this call
	// (using either the Record element, or the
	// record attribute of the Dial element). Only
	// present in the notification following a
	// partial recording, or in the final notification
	// if it is a terminal recording.
	Attribute("recordingUrl", String, func() {
		Description("The URL of the recording made for this call")
	})
	//  Only present in the final notification.
	Attribute("durationInSeconds", String, func() {
		Description("The duration of the call in seconds.")
	})
	// Only present in the final notification.
	Attribute("currencyCode", String, func() {
		Description("The currency used to bill this call (e.g KES, USD, GBP).")
	})
	// Only present in the final notification.
	Attribute("amount", String, func() {
		Description("The total cost of the call.")
	})
	// Only present in the final notification.
	Attribute("callSessionState", String, func() {
		Description("The final status of the call.")
	})
	// Only present in the final notification.
	Attribute("dialDestinationNumber", String, func() {
		Description("The number which a call was forwarded to if the Dial action was used.")
	})
	// Only present in the final notification.
	Attribute("dialDurationInSeconds", String, func() {
		Description("The duration of the dialed call if the Dial action was used.")
	})
	// Only present in the final notification.
	Attribute("dialStartTime", String, func() {
		Description("The time the dial action began if the Dial action was used.")
	})
	Attribute("hangupCause", String, func() {
		Description("The reason a call could have ended")
		Enum(
			// This cause indicates that the call is
			// being cleared because one of the users
			// involved in the call has requested that
			// the call be cleared. This also means
			// the call was successfully answered and
			// successfully ended.
			"NORMAL_CLEARING",

			// This cause indicates the called party
			// does not wish to accept this call.
			"CALL_REJECTED",

			// This cause indicates that the network
			// is not functioning correctly and that
			// the condition is not likely to last a
			// long period of time. The user may wish
			// to try another call attempt almost
			// immediately.
			"NORMAL_TEMPORARY_FAILURE",

			// This cause indicates an expiration of a
			// request, to the called party. This is
			// often associated with NAT problems.
			// It affects mostly soft phones with
			// SIP numbers.
			"RECOVERY_ON_TIMER_EXPIRE",

			// The caller initiated a call and then
			// hang up before the recipient picked up.
			// This normally happens when using the
			// Dial call action.
			"ORIGINATOR_CANCEL",

			// This occurs when a call is initiated
			// to multiple phone numbers. Once one
			// recipient picks up, the others will
			// have a LOSE_RACE hangup cause. You
			// can get this when using the Dial
			// call action.
			"LOSE_RACE",

			// This cause is used to indicate that
			// the called party is unable to accept
			// another call because the user busy
			// condition has been encountered/engaged
			// on another call.
			"USER_BUSY",

			// This cause is used when the called party
			// has been alerted but does not respond
			// with a connect indication within a
			// prescribed period of time.
			"NO_ANSWER",

			// This cause is used when a called party
			// does not respond to a call establishment
			// message with either an alerting or
			// connect indication within the prescribed
			// period of time allocated.
			"NO_USER_RESPONSE",

			// This cause value is used when a mobile
			// station has logged off, radio contact is
			// not obtained with a mobile station or if
			// a personal telecommunication user is
			// temporarily not addressable at any
			// user-network interface.
			"SUBSCRIBER_ABSENT",

			// This cause is used to report a service
			// not available.
			"SERVICE_UNAVAILABLE",

			// This means you tried to originate a call to
			// a SIP user who forgot to register/hasn’t
			// registered.
			"USER_NOT_REGISTERED",

			// This cause indicates that the called party
			// cannot be reached because, although the
			// called party number is in a valid format,
			// it is not currently allocated (assigned).
			"UNALLOCATED_NUMBER",

			// This cause happens on very rare occasions
			// when a valid hangup cause can’t be obtained.
			// We (AfricasTalking) are usually alerted for
			// this and we look into it immediately.
			"UNSPECIFIED",
		)
	})
})

// When the transfer has been initiated AT sends any of these events
// to the event notifications URL, you can check from these form
// fields in the events.
var CallTransferEvent = Type("CallTransferEvent", func() {
	Description("Event Notifications sent from AT after call transfer initiated.")

	Attribute("callSessionState", String, func() {
		Enum("Active",
			"Transferred",
			"TransferCompleted",
		)
	})
	Attribute("isActive", String, func() {
		Enum("0", "1")
		Default("1")
	})
	Attribute("status", String, func() {
		Enum("Success")
	})
	Attribute("callTransferParam", String, func() {
		Description("+2347XXXXXXXXX:20, (20 is the duration in seconds)")
	})
	Attribute("callTransferredToNumber", String, func() {
		Description("Number call was transferred to")
	})
	Attribute("callTransferState", String, func() {
		Enum(" Active",
			"Completed",
			"CallerHangup",
			"CalleeHangup",
		)
	})
	Attribute("callTransferHangupCause", String, func() {
		Enum("DestinationNotSupported",
			"InvalidPhoneNumber",
			"NoActiveClient",
			"NotAllowed",
		)
	})
})

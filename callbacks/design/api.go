package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// 1a. SMS POST request to AT gateway
// 1b. JSON Response from AT gateway
// 2. AT gateway sends SMS to user.
// 3. Status update from Telco
// 4. Status update from AT gateway to Callback

// Delivery Report notification contents
var DeliveryReport = Type("DeliveryReport", func() {
	Description("Sent whenever the MSP confirms or rejects delivery of a message")
	Attribute("id", String, func() {
		Description("A unique identifier for each message.") // Same id as the one in the response when a message is sent
	})
	Attribute("status", String, func() {
		Description("The status of the message.")
		Enum(
			// The message has successfully been sent by our network.
			"Sent",

			// The message has successfully been submitted to the MSP (Mobile Service Provider).
			"Submitted",

			// The message has been queued by the MSP.
			"Buffered",

			// The message has been rejected by the MSP.
			// This is a final status.
			"Rejected",

			// The message has successfully been delivered to the receiver’s handset.
			// This is a final status.
			"Success",

			// The message could not be delivered to the receiver’s handset.
			// This is a final status.
			"Failed",
		)
	})
	Attribute("phoneNumber", String, func() {
		Description("This is phone number that the message was sent out to.")
	})
	Attribute("networkCode", String, func() {
		Description("A unique identifier for the telco that handled the message.")
		Enum(
			"62120", // 62120: Airtel Nigeria
			"62130", // 62130: MTN Nigeria
			"62150", // 62150: Glo Nigeria
			"62160", // 62160: Etisalat Nigeria
			"63510", // 63510: MTN Rwanda
			"63513", // 63513: Tigo Rwanda
			"63514", // 63514: Airtel Rwanda
			"63902", // 63902: Safaricom
			"63903", // 63903: Airtel Kenya
			"63907", // 63907: Orange Kenya
			"63999", // 63999: Equitel Kenya
			"64002", // 64002: Tigo Tanzania
			"64003", // 64003: Zantel Tanzania
			"64004", // 64004: Vodacom Tanzania
			"64005", // 64005: Airtel Tanzania
			"64007", // 64007: TTCL Tanzania
			"64009", // 64009: Halotel Tanzania
			"64101", // 64101: Airtel Uganda
			"64110", // 64110: MTN Uganda
			"64111", // 64111: UTL Uganda
			"64114", // 64114: Africell Uganda
			"65001", // 65001: TNM Malawi
			"65010", // 65010: Airtel Malawi
			"99999", // 99999: Athena (sandbox environment).
		)
	})
	Attribute("failureReason", String, func() {
		Description("Only provided if status is Rejected or Failed.")
		Enum(
			// This occurs when the subscriber doesn’t have enough airtime for
			// a premium subscription service/message
			"InsufficientCredit",

			// This occurs when a message is sent with an invalid linkId for an
			// onDemand service
			"InvalidLinkId",

			// This occurs when the subscriber is inactive or the account deactivated
			// by the MSP (Mobile Service Provider).
			"UserIsInactive",

			// This occurs if the user has been blacklisted not to receive messages
			// from a paricular service (shortcode or keyword)
			"UserInBlackList",

			// This occurs when the mobile subscriber has been suspended by the MSP.
			"UserAccountSuspended",

			// This occurs when the message is passed to an  MSP where the subscriber
			// doesn’t belong.
			"NotNetworkSubscriber",

			// This occurs when the message from a subscription product is sent to a
			// phone number that has not subscribed to the product.
			"UserNotSubscribedToProduct",

			// This occurs when the message is sent to a non-existent mobile number.
			"UserDoesNotExist",

			// This occurs when message delivery fails for any reason not listed above
			// or where the MSP didn’t provide a delivery failure reason.
			"DeliveryFailure",
		)
	})
	// Note: This only applies for premium SMS messages.
	Attribute("retryCount", String, func() {
		Description("Number of times the request to send a message to the device was retried before it succeeded or definitely failed.")
	})
})

// Incoming message notification contents
var IncomingMessage = Type("IncomingMessage", func() {
	Description("Sent whenever a message is sent to any of your registered shortcodes.")
	Attribute("date", String, func() {
		Description("The date and time when the message was received.")
		Format(FormatDate)
	})
	Attribute("from", String, "The number that sent the message.")
	Attribute("id", String, "The internal ID that we use to store this message.")
	Attribute("linkId", String, func() {
		Description("Parameter required when responding to an on-demand user request with a premium message.")
	})
	Attribute("text", String, "The message content.")
	Attribute("to", String, "The number to which the message was sent.")
	Attribute("networkCode", String, func() {
		Description("A unique identifier for the telco that handled the message.")
		Enum(
			"62120", // 62120: Airtel Nigeria
			"62130", // 62130: MTN Nigeria
			"62150", // 62150: Glo Nigeria
			"62160", // 62160: Etisalat Nigeria
			"63510", // 63510: MTN Rwanda
			"63513", // 63513: Tigo Rwanda
			"63514", // 63514: Airtel Rwanda
			"63902", // 63902: Safaricom
			"63903", // 63903: Airtel Kenya
			"63907", // 63907: Orange Kenya
			"63999", // 63999: Equitel Kenya
			"64002", // 64002: Tigo Tanzania
			"64003", // 64003: Zantel Tanzania
			"64004", // 64004: Vodacom Tanzania
			"64005", // 64005: Airtel Tanzania
			"64007", // 64007: TTCL Tanzania
			"64009", // 64009: Halotel Tanzania
			"64101", // 64101: Airtel Uganda
			"64110", // 64110: MTN Uganda
			"64111", // 64111: UTL Uganda
			"64114", // 64114: Africell Uganda
			"65001", // 65001: TNM Malawi
			"65010", // 65010: Airtel Malawi
			"99999", // 99999: Athena (This is a custom networkCode that only applies when working in the sandbox environment).
		)
	})
})

// The instructions on how to opt out are automatically appended to the
// first message you send to the mobile subscriber. From then onwards,
// any other message will be sent ‘as is’ to the subscriber.
var BulkSMSOptOut = Type("BulkSMSOptOut", func() {
	Description("Sent whenever a user opts out of receiving messages from your alphanumeric sender ID")
	Attribute("senderId", String, func() {
		Description("This is the shortcode/alphanumeric sender id the user opted out from.")
	})
	Attribute("phoneNumber", String, func() {
		Description("This will contain the phone number of the subscriber who opted out.")
	})
})

var SubscriptionNotification = Type("SubscriptionNotification", func() {
	Description("Sent whenever someone subscribes or unsubscribes from any of your premium SMS products.")

	Attribute("phoneNumber", String, func() {
		Description("Phone number to subscribe or unsubscribe.")
	})
	Attribute("shortCode", String, func() {
		Description("The short code that has this product.")
	})
	Attribute("keyword", String, func() {
		Description("The keyword of the product that the user has subscribed or unsubscribed from.")
	})
	Attribute("updateType", String, func() {
		Description("The type of the update.")
		Enum("addition", "deletion")
	})
})

// 1. User dials USSD Code.
// 2. Africa'sTalking Gateway sends POST request to our web application.
// 3. Our application sends an plain text response to Africa'sTalking Gateway.
// 4. Africa'sTalking Gateway responds to user with USSD Menu.
var USSDPayload = Type("USSDPayload", func() {
	Description("This request is made when the user dials a USSD code and every time they respond to a menu.")

	// Sent every time a mobile subscriber response has been received.
	Attribute("sessionId", String, func() {
		Description("A unique value generated when the session starts.")
		// Let the Mobile Service Provider know whether the session is complete or not.
		// "CON" - Session is ongoing
		// "END" - Last session response
		Attribute("SessionStatus", String, func() {
			Enum("CON", "END")

		})
	})
	Attribute("phoneNumber", String, func() {
		Description("The number of the mobile subscriber interacting with your ussd application.")
	})
	Attribute("networkCode", String, func() {
		Description("The telco of the phoneNumber interacting with your ussd application.")
	})
	// Register a service code with AT.
	Attribute("serviceCode", String, func() {
		Description("This is the USSD code assigned to your application")
	})
	// It is an empty string in the first notification of a session.
	// After that, it concatenates all the user input within the
	// session with a * until the session ends.
	Attribute("text", String, func() {
		Description("This shows the user input.")
	})
	// Callback URL registered on AT
	// This can be called by AT whenever they get a request from a client coming into our system.
	Attribute("CallbackURL", String, func() {
		Description("Callback URL registered on AT")
		Pattern(`(?i)^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
		Format(FormatURI)
	})
})

// The Voice API from Africa'sTalking sends a notification when a
// specific event happens under the following categories.
// 1. Outbound Calls: Sent whenever you make a call from a registered SIP number.
// 2. Inbound Calls: Sent when a call comes to your virtual or SIP number.
// 3. After Input: Sent whenever an action in your response requires user input (such as GetDigits and Record)
// 4. When Call Ends: Sent after a call ends. This is the final notification and contains some extra
// information about the call like the cost and duration.

// VoiceNotification are voice notification contents sent to our callback URL.
var VoiceNotificationPayload = Type("VoiceNotificationPayload", func() {
	Attribute("isActive", String, func() {
		Description("Lets us know whether the call is in session state")
		Default("0")
	})
	Attribute("sessionId", String, func() {
		Description("A unique identifier generated during each call session")
	})
	// Inbound calls are initiated by a phone user.
	// Outbound calls are initiated by our application.
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
	// Only present in a notification  following a GetDigits response.
	Attribute("dtmfDigits", String, func() {
		Description("The digits that a user enters in response to a getDigits request")
	})
	// The URL of the recording made for this call (using either the Record element,
	// or the record attribute of the Dial element).
	// Only present in the notification following a partial recording,
	// or in the final notification if it is a terminal recording.
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
	Attribute("callSessionState", String, "The final status of the call.")
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
			// This cause indicates that the call is being cleared because one of the users
			// involved in the call has requested that the call be cleared.
			// This also means the call was successfully answered and successfully ended.
			"NORMAL_CLEARING",

			// This cause indicates the called party does not wish to accept this call.
			"CALL_REJECTED",

			// This cause indicates that the network is not functioning correctly and that the condition is not likely to
			// last a long period of time. The user may wish to try another call attempt almost immediately.
			"NORMAL_TEMPORARY_FAILURE",

			// This cause indicates an expiration of a request, to the called party.
			// This is often associated with NAT problems.
			// It affects mostly soft phones with SIP numbers.
			"RECOVERY_ON_TIMER_EXPIRE",

			// The caller initiated a call and then hang up before the recipient picked up.
			// This normally happens when using the Dial call action.
			"ORIGINATOR_CANCEL",

			// This occurs when a call is initiated to multiple phone numbers.
			// Once one recipient picks up, the others will have a LOSE_RACE hangup cause.
			// You can get this when using the Dial call action.
			"LOSE_RACE",

			// This cause is used to indicate that the called party is unable to accept another call because
			// the user busy condition has been encountered/engaged on another call.
			"USER_BUSY",

			// This cause is used when the called party has been alerted but does not respond with a connect indication
			//  within a prescribed period of time.
			"NO_ANSWER",

			// This cause is used when a called party does not respond to a call establishment message with
			// either an alerting or connect indication within the prescribed period of time allocated.
			"NO_USER_RESPONSE",

			// This cause value is used when a mobile station has logged off, radio contact is not obtained with
			// a mobile station or if a personal telecommunication user is temporarily not addressable at any user-network interface.
			"SUBSCRIBER_ABSENT",

			// This cause is used to report a service not available.
			"SERVICE_UNAVAILABLE",

			// This means you tried to originate a call to a SIP user who forgot to register/hasn’t registered.
			"USER_NOT_REGISTERED",

			// This cause indicates that the called party cannot be reached because, although the called
			// party number is in a valid format, it is not currently allocated (assigned).
			"UNALLOCATED_NUMBER",

			// This cause happens on very rare occasions when a valid hangup cause can’t be obtained.
			// We (AfricasTalking) are usually alerted for this and we look into it immediately.
			"UNSPECIFIED",
		)
	})
})

var VoiceNotificationResult = ResultType("VoiceNotificationResult", func() {
	Description("Voice Notification delivered to our /notifications callback URL")
	TypeName("VoiceNotification")
	ContentType("application/xml")

	Attributes(func() {
		Extend(VoiceNotificationPayload)
	})
	View("final", func() {
		Attribute("isActive")
		Attribute("sessionId")
		Attribute("direction")
		Attribute("destinationNumber")
		Attribute("callerNumber")
		Attribute("callerCountryCode")
		Attribute("callStartTime")
		Attribute("dtmfDigits")
		Attribute("recordingUrl")
		Attribute("durationInSeconds")
		Attribute("currencyCode")
		Attribute("amount")
		Attribute("callSessionState")
		Attribute("dialDestinationNumber")
		Attribute("dialDurationInSeconds")
		Attribute("dialStartTime")
	})
	View("recording", func() {
		Attribute("isActive")
		Attribute("sessionId")
		Attribute("direction")
		Attribute("destinationNumber")
		Attribute("callerNumber")
		Attribute("callerCountryCode")
		Attribute("callStartTime")
		Attribute("recordingUrl")
	})
	View("dtmf", func() {
		Attribute("isActive")
		Attribute("sessionId")
		Attribute("direction")
		Attribute("destinationNumber")
		Attribute("callerNumber")
		Attribute("callerCountryCode")
		Attribute("callStartTime")
		Attribute("dtmfDigits")
	})
})

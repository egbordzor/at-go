package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("at", func() {

	// API title.
	Title("AfricasTalking API")

	// Description of API.
	Description("HTTP service for interacting with all AfricasTalking API.")

	// Version of API
	Version("1.0")

	// Terms of use of API
	TermsOfService("https://github.com/wondenge/atgo/blob/master/LICENSE")

	// Contact info for Author and Maintainer.
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})

	// License for Library usage.
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/atgo/blob/master/LICENSE")
	})

	// Library Documentation.
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/atgo/blob/master/README.md")
	})

	// Server describes a single process listening for client requests.
	Server("atsvr", func() {
		Description("atsvr hosts the AfricasTalking HTTP Service.")

		// List services hosted by this server.
		Services()

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			URI("https://api.sandbox.africastalking.com")
		})

		Host("production", func() {
			Description("Production hosts.")
			URI("https://{subdomain}.africastalking.com")

			// Variable describes a URI variable.
			Variable("subdomain", String, "Name of Sub-Domain", func() {
				// URI parameters must have a default value
				// and/or an enum validation.
				Enum("api", "content")
				Default("api")
			})
		})
	})
})

var _ = Service("at", func() {

	HTTP(func() {
		Path("/")
	})

	// Method defines a single service method.
	Method("generate", func() {

		// Description sets the expression description.
		Description("Generates a valid auth token")

		// Payload defines the data type of an method input.
		Payload(func() {
			Attribute("username", String, func() {
				Description("Africa's Talking Username.")
				Example("sandbox")
				Default("sandbox")
			})
			Attribute("apiKey", String, func() {
				Description("Africa's Talking API Key.")
			})

			// Required adds a "required" validation to the attribute.
			Required("username", "username")
		})

		// Result defines the data type of a method output.
		Result(TokenMedia)

		// POST creates a route using the POST HTTP method.
		// POST request to https://api.africastalking.com/auth-token/generate
		POST("/generate")

		HTTP(func() {

			// Headers describes HTTP request/response or gRPC response headers.
			Headers(func() {
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/json")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})

				// Required adds a "required" validation to the attribute.
				Required("Content-Type", "Accept")
			})

			// Response describes a HTTP or a gRPC response.
			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("fetch", func() {
		Description("Initiate an application data request.")

		// Payload describes the method payload.
		Payload(func() {
			Attribute("username", String, "username of the application making the request")
			Required("username")
		})

		// Result describes the method result.
		Result(UserMedia)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {

			// Requests to the service consist of HTTP GET requests.
			// Live: https://api.africastalking.com/version1/user
			//Sandbox: https://api.sandbox.africastalking.com/version1/user
			GET("/user?username={username}")

			Headers(func() {

				// Attribute describes an object field
				Attribute("apiKey", String, func() {
					Description("Africa’s Talking application apiKey.")
				})
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/json")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/json")
				})
				Required("apiKey", "Content-Type", "Accept")
			})

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})

	// Send Bulk SMS
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/version1/messaging
	// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
	Method("SendBulkSMS", func() {
		Description("Send Bulk SMS")
		Payload(BulkSMSPayload)
		Result(BulkSMSResponse)
		HTTP(func() {
			POST("version1/messaging")

			// Authenticating using an API using username and API Key
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Send Premium SMS
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/version1/messaging
	// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
	Method("SendPremiumSMS", func() {
		Description("Send Premium SMS")
		Payload(PremiumSMSPayload)
		Result(PremiumSMSResponse)
		HTTP(func() {
			POST("version1/messaging")

			// Authenticating using an API using username and API Key
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Fetch Messages
	// Requests to the service consist of HTTP GET requests.
	// Live: https://api.africastalking.com/version1/messaging
	// Sandbox: https://api.sandbox.africastalking.com/version1/messaging
	Method("fetchMessages", func() {
		Description("Incrementally fetch messages from application inbox.")
		Payload(FetchMsgPayload)
		Result(FetchMsgResponse)
		HTTP(func() {
			GET("version1/messaging?{username}&{lastReceivedId}")
			Response(StatusOK)
			Params(func() {
				Param("username:username", String, "Africa’s Talking application username.")
				Param("lastReceivedId:lastReceivedId", String, "ID of the message last processed.")
				Required("username")
			})
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
		})
	})

	// Generate a checkout token
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/checkout/token/create
	// Sandbox: https://api.sandbox.africastalking.com/checkout/token/create
	Method("checkout", func() {
		Description("Generate a checkout token")
		Payload(CreateCheckoutTokenPayload)
		Result(CreateCheckoutTokenResponse)
		HTTP(func() {
			POST("checkout/token/create")
			Headers(func() {
				Header("Content-Type:Content-Type", String, "Content-Type")
				Required("Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Subscribe a phone number
	// Requests to the service consist of HTTP POST requests.
	// Live: https://content.africastalking.com/version1/subscription/create
	// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create
	Method("newSub", func() {
		Description("Subscribe a phone number")
		Payload(CreateSubPayload)
		Result(CreateSubResponse)
		HTTP(func() {
			POST("version1/subscription/create")
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Incrementally fetch your premium sms subscriptions.
	// Requests to the service consist of HTTP GET requests.
	// Live: https://api.africastalking.com/version1/subscription
	// Sandbox: https://api.sandbox.africastalking.com/version1/subscription
	Method("fetchSub", func() {
		Description("Incrementally fetch your premium sms subscriptions.")
		Payload(FetchSubPayload)
		Result(FetchSubResponse)
		HTTP(func() {
			GET("/version1/subscription?username={username}&shortCode={shortCode}&keyword={keyword}&lastReceivedId={lastReceivedId}")
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Params(func() {
				Param("username:username", String, "Africa’s Talking application username.")
				Param("shortCode:shortCode", String, "Premium short code mapped to your account")
				Param("keyword:keyword", String, "Premium keyword under short code mapped to your account.")
				Param("lastReceivedId:lastReceivedId", String, "ID of the message last processed.")
				Required("username", "shortCode", "keyword")
			})
			Response(StatusOK)
		})
	})

	// Delete a Premium SMS Subscription
	// Requests to the service consist of HTTP POST requests.
	// Live: https://api.africastalking.com/version1/subscription/delete
	// Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete
	Method("removeSub", func() {
		Description("Delete a Premium SMS Subscription")
		Payload(PurgeSubPayload)
		Result(PurgeSubResponse)
		HTTP(func() {
			POST("version1/subscription/delete")
			Headers(func() {
				Attribute("apiKey:apiKey", String, "Africa’s Talking application apiKey")
				Attribute("Content-Type:Content-Type", String, "The requests content type")
				Attribute("Accept:Accept", String, "The requests response type")
				Required("apiKey", "Content-Type")
			})
			Response(StatusCreated)
		})
	})

	// Outbound Calls Service//

	// Method describes a service method (endpoint)
	Method("add", func() {
		Description("Makes outbound calls through the Africa'sTalking Voice API")

		// Payload describes the method payload.
		Payload(MakeCallPayload)

		// Result describes the method result.
		Result(MakeCallResult)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
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
				Required("Content-Type")
			})

			// Requests to the service consist of HTTP POST requests.
			// Live: https://voice.africastalking.com/call
			// Sandbox: https://voice.sandbox.africastalking.com/call
			POST("/call")

			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})
	})

	// Calls Transfer Service //
	// Method describes a service method (endpoint)
	Method("add", func() {
		Description("Transfers call to another number")

		// Payload describes the method payload.
		Payload(CallTransferPayload)

		// Result describes the method result.
		Result(CallTransferMedia)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/json")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			// Requests to the service consist of HTTP POST requests.
			// Live: https://voice.africastalking.com/callTransfer
			// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
			POST("/callTransfer")
			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("say", func() {
		Description("Set a text to be read out to the caller.")

		// Payload describes the method payload.
		Payload(Say)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("play", func() {
		Description("Play back an audio file located anywhere on the web.")

		// Payload describes the method payload.
		Payload(Play)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("getDigits", func() {
		Description("Get digits a user enters on their phone in response to a prompt from application")

		// Payload describes the method payload.
		Payload(GetDigits)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("dial", func() {
		Description("Connect the user who called your phone number to an external phone number.")

		// Payload describes the method payload.
		Payload(Dial)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("record", func() {

		// Can be retrieved and played later.
		Description(" Record a call session into an mp3 file.")

		// Payload describes the method payload.
		Payload(Record)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("enqueue", func() {
		Description("Pass an incoming call to a queue to be handled later.")

		// Payload describes the method payload.
		Payload(Enqueue)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("dequeue", func() {
		Description("Pass the calls enqueued to a separate number to be handled.") // e.g by an agent.

		// Payload describes the method payload.
		Payload(Dequeue)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("redirect", func() {
		Description("Transfer control of the call to the script whose URL is passed in.")

		// Payload describes the method payload.
		Payload(Redirect)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Method describes a service method (endpoint)
	Method("reject", func() {
		Description("Reject an incoming call without incurring any usage charges.")

		// Payload describes the method payload.
		Payload(Reject)

		// Result describes the method result.
		Result(String)

		// Requests to the service consist of HTTP POST requests.
		// Live: https://voice.africastalking.com/callTransfer
		// Sandbox: https://voice.sandbox.africastalking.com/callTransfer
		POST("/callTransfer")

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			Headers(func() {

				// Attribute describes an object field
				Attribute("Content-Type", String, func() {
					Description("The requests content type.")
					Enum("application/x-www-form-urlencoded", "application/json", "application/xml")
					Default("application/xml")
				})
				Attribute("Accept", String, func() {
					Description("The requests response type.")
					Enum("application/json", "application/xml")
					Default("application/xml")
				})
				Required("Content-Type")
			})

			Response(StatusOK)
		})
	})

	// Queue Calls Service //
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

	// Media Upload Service.//
	// Method describes a service method (endpoint)
	Method("add", func() {
		Description("Uploads media or audio files to Africa'sTalking servers with the extension .mp3 or .wav")

		// Payload describes the method payload.
		Payload(UploadMediaFile)

		// Result describes the method result.
		Result(String)

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
			// Live: https://voice.africastalking.com/mediaUpload
			// Sandbox: https://voice.sandbox.africastalking.com/mediaUpload
			POST("/mediaUpload")

			// You can check the HTTP Response Code to determine whether the request was successful.
			// Any response code other than 201 (Created) indicates that the call was not initiated
			Response(StatusCreated)
		})
	})
})

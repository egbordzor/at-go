# Voice

Take advantage of fully web based voice infrastructure and deploy fully fledged services without the need for expensive physical equipment. Build dynamic, scalable, fully-featured voice applications that reside entirely in the cloud using our Voice APIs without the need to purchase and maintain expensive voice equipment. This reduces your business costs while increasing the efficiency of how you reach your users.

Some possible use cases

- A call center.
- User authentication (Two-Factor Authentication)
- Call Blasting/Robocalls
- Conducting Surveys
- Receiving Feedback

## Product Overview

Programmable Voice allows you to enrich your call experience by adding features that would have otherwise be inaccessible using the normal physical lines. Our fully featured API gives you a wide number of actions that you can use together to build intuitive solutions

- VOIP Capabilities
  Use our SIP endpoints to make and receive calls on your browser. This lowers infrastructure costs significantly

* Call Queues
  Dynamically enqueue and dequeue callers to increase your call handling capacity. With this feature, calls can be put in a queue to be handled later or forwarded to another agent

- Call Transfer
  Send a caller to another phone number or agent when the need arises e.g wrong recipient, need to transfer to a different department etc

* Call Recording
  Record a call session into an MP3 file that you can retrieve and play later. This will prove useful when conducting surveys or ensuring quality of service in call centres.

- Say and Play
  Automatically read text information to callers using our free Text to Speech Engine. Using the play action, you can playback an audio file that is located anywhere on the web

- SIP Trunking
  Direct incoming and outgoing calls to a virtual phone provider or your company's private network. This will lower your call charges and reduce need for more physical infrastructure.

## Use cases

1. Two Factor Authentication
   A phone number has incredible reach and is growing in importance as a useful tool for verifying user identity. You can enhance security in your applications by calling your users directly with authentication tokens

2. Premium Services
   You build and monetise services that run on voice using our Premium Rated numbers. Services such as interactive dating or marketing audio content can be enhanced by taking advantage of the personal nature of voice

3. Build a Call Center
   Enhance our customer support experience by building an Interactive Voice Response(IVR)functionality into your support workflow. Add dynamic queuing, call recording and routing capabilities to differentiate how your business responds to client needs

4. Customer Notifications
   Keep your customers in the know about issues concerning them. This will in turn improve your customer satisfaction rating and increase customer loyalty.

5. Carrying out Surveys/Data Collection
   Receive and store responses from a data collection exercise. You can use the IVR (Interactive Voice Response) feature for caller prompts and the record action to store all responses onto a database.

6. IVR
   Make it easier for your callers to more efficiently get the help they need through voice prompts. Used mostly in customer service to choose language, to speak to an agent,to suggest common solutions for recurring problems etc.

7. Robocalls/Call blasting
   Reach out to hundreds of people in your contact list with a single call blast and play a pre-recorded message. Works well for marketing/promotional campaigns and Customer notifications.

8. Call Tagging/Identification
   Make your work of analysing calls easier by tagging them using our ClientRequestID feature. You can tag the phone calls depending on function e.g support call, by action e.g failed calls, by agent or any other way depending on what you want to measure.

## How the SMS service works

1. User calls via AT gateway.
2. AT Gateway sends POST request to our web application.
3. Our application sends an XML response to AT Gateway.
4. AT Gateway responds to User.

To make a call from your application, you will have to send a POST request to our Voice API. When the call is picked, we will send a notification to your callback url so that your application can let us know how to handle the call.

The next action on such calls will depend on the kind of response we get from the notification sent to your application.

If the action. requires a user’s input, we will send another notification to your callback url to submit the input and find out the next action. This continues until either the user ends the call or you respond with an action that terminates the call.

Once the call is terminated, we will send a final notification to your callback url with some extra information about the call such as cost of the call and its duration.

## Making a call

You can make an outbound call through our Voice API by sending a HTTP POST requsest to one of the following endpoints:

### Endpoints

Live: https://voice.africastalking.com/call
Sandbox: https://voice.sandbox.africastalking.com/call

## Handling a call

Handling calls made to your Africa’s Talking phone number is as easy as implementing a script on your web server that handles HTTP POST requests.

### Steps Involved In Handling A Call Session

Here is how it all comes together…

- We receive a call to your phone number on our voice gateways, or you successfully initiate a call using our calling API
- Our API sends a POST request to the notification callback URL that you have set for that phone number in your Voice Dashboard.
- Your application responds with XML that tells our API how to handle the call. This XML will typically contain a list of actions that our API will execute in sequence.
- Our API translates those actions into events or messages relayed back to the caller.

## Call Transfer

You can transfer your call to another number by making a HTTP POST request to one of the following endpoints:

### Endpoints

Live: https://voice.africastalking.com/callTransfer
Sandbox: https://voice.sandbox.africastalking.com/callTransfer

### Voice Actions

A voice action is what your application sends in response to a notification that tells the Voice API how to handle the current call.

- Say: With this action, you set a text and we will read it out to the caller.
- Play: This element lets you play back an audio file that is located anywhere on the web.
- Get Digits: You can use this element to get the digits that a user enters on their phone in response to a prompt from your application.
- Dial: You can use this action to connect the user who called your phone number to an external phone number.
- Record: This element lets you record a call session into an mp3 file that you can then retrieve and play later. Our API supports terminal recording or partial recording. For final recording, the recording starts when the record action is called until the mobile user hangs up. Partial recording may be useful where a particular user response is required like while prompting for a name.
- Enqueue: Lets you pass an incoming call to a queue to be handled later.
- Dequeue: Lets you pass the calls enqueued to a separate number so that it can be handled e.g by an agent.
- Redirect: This action will transfer control of the call to the script whose URL is passed in. This can help you better organize your call handling logic by spreading the logic across multiple scripts.
- Reject: This action lets you reject an incoming call without incurring any usage charges from our APIs.

Note
Any action that takes input from the user such as [Get Digits](/docs/voice/actions/getdigits) and [Record](/docs/voice/actions/record) will trigger another notification to your application with those values so we can know what next action to take. This is so that you can perform some logic depending on the user's input.

## Queued Calls

Queuing is a feature mainly used when you have more calls than you can handle at one time. Incoming calls will be put in a queue and handled one by one until all of them are out of the queue (dequeued).

You can find the number of queued calls by sending a HTTP POST request to one of the following endpoints:

### Endpoints

Live: https://voice.africastalking.com/queueStatus
Sandbox: https://voice.sandbox.africastalking.com/queueStatus

## Upload media file

You can upload media / audio file to our servers with the extension .mp3 or .wav only. This media files will be played when called upon by one of our voice actions.

- Play: `url`: This contains the audio file you want played during a call.
- Call queueing `holdMusic`: This contains the audio file you want played when the user has been queued waiting to be dequeued.
- Dial `holdMusic`: This contains the audio file you want played when a number has been dialed before it’s picked.
  You can upload media by sending a HTTP POST requsest to one of the following endpoints:

### Endpoints

Live: https://voice.africastalking.com/mediaUpload
Sandbox: https://voice.sandbox.africastalking.com/mediaUpload

You can check the HTTP Response Code to determine whether the request was successful. Any response code other than 201 (Created) indicates that the call was not initiated

## Voice Notifications

The Voice API sends a notification when a specific event happens. To receive these notifications you need to setup a voice callback URL. From the dashboard select Voice -> Phone Numbers -> Actions -> Callback.

### Types of voice notifications

Voice API notifications are sent for various categories as shown below:

- Outbound calls: Sent whenever you make a call from a registered SIP number.
- Inbound calls: Sent when a call comes to your virtual or SIP number.
- After input: Sent whenever an action in your response requires user input (such as GetDigits and Record)
- When call ends: Sent after a call ends. This is the final notification and contains some extra information about the call like the cost and duration.

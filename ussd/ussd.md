# USSD

Build dynamic, scalable, fully featured USSD applications.

## How the SMS service works

1. User dials USSD code.
2. AT gateway sends POST request to our application.
3. Application sends plain text response to AT gateway.
4. AT gateway sends USSD menu to user.

Processing USSD requests using AT's API is very easy once your account is set up. In particular, you will need to:

- Register a service code with them.
- Register a URL that AT can call whenever they get a request from a client coming into their system.

Once you register your callback URL, any requests that they receive belonging to you will trigger a callback that sends the request data to that URL using HTTP POST.

All you have to do at this point is print the string response that you would like them to send back to the user.

A few things to note about USSD:

- USSD is session driven. Every request AT sends you will contain a sessionId, and this will be maintained until that session is completed
- You will need to let the Mobile Service Provider know whether the session is complete or not. If the session is ongoing, please begin your response with CON. If this is the last response for that session, begin your response with END.
- If AT gets a HTTP error response (Code 40X) from your script, or a malformed response (does not begin with CON or END), we will terminate the USSD session gracefully.

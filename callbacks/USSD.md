# USSD

Build real time interactive experiences for your users that is accessible on both feature and smartphones. Easily integrate our simple and reliable USSD API and start delivering scalable solutions to your users!

Build dynamic, scalable, fully featured USSD applications.

## Product Overview

Unstructured Supplementary Service Data or USSD in short is a service that allows mobile phone users to interact with a remote application from their device in real time. A predefined session is started once a user dials in to facilitate the transfer of information between your application and the user. It is a highly scalable service as it does not require an internet connection and is supported by both feature and smartphones

- Dedicated USSD
  This is a three digit code for example \*384# that is used solely by one user for their company or service

- Shared USSD
  This is a code that is used by multiple companies or services. Each company is allocated a channel on our dedicated codes, for example *384*65#

## Use cases

1. Query services
   Give access to structured information from your application with USSD menus. The interactive nature can layer leading questions to deliver the right information

2. Data collection
   USSD is a powerful tool for data collection because it is universally accessible on any mobile phone. The session-based nature of USSD services also make them better suited for structured surveys as compared to SMS

3. User Registration
   USSD is a powerful marketing tool for leading users to quickly register for your service. Enhance the customer on boarding experience by adding services such as SMS and Voice to your USSD application.

## Features

1. High Capacity
   Our USSD gateway is built to handle thousands of concurrent API requests at any one time

2) Flexible Pricing
   Choose between pre-paid and post paid billing options depending on your business needs.

3. Painless Integration
   We have made it easy for a developer to build and deploy a USSD application within minutes

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

# 1. SMS Callback Service

SMSs are sent from an application by making a HTTP POST request to the AfricasTalking SMS API. For each request AfricasTalking receives, they respond with a notification back indicating whether the sms transaction was successful or failed.

Other events AfricasTalking sends notifications for are;

- Whenever they receive a message for an application.
- When a user opts out of receiving messages from the alphanumeric code, sender id, or if they subscribe/unsubscribe from the premium SMS service.

To receive these notifications you need to setup a callback URL depending on the type of notification you would like to receive. SMS API notifications are sent for various SMS categories as shown below:

## 1.1 Delivery Reports

These are sent whenever the mobile service provider confirms or rejects delivery of a message. To receive delivery reports, you need to set a delivery report callback URL. From the dashboard select SMS -> SMS Callback URLs -> Delivery Reports.

## 1.2 Incoming Messages.

These are sent whenever a message is sent to any of your registered shortcodes. To receive incoming messages, you need to set an incoming messages callback URL. From the dashboard select SMS -> SMS Callback URLs -> Incoming Messages.

## 1.3 Bulk SMS Opt Out

These are sent whenever a user opts out of receiving messages from your alphanumeric sender ID. To receive bulk sms opt out notifications, you need to set a bulk sms opt out callback URL. From the dashboard select SMS -> SMS Callback URLs -> Bulk SMS Opt Out. The instructions on how to opt out are automatically appended to the first message you send to the mobile subscriber. From then onwards, any other message will be sent ‘as is’ to the subscriber.

## 1.4 Subscription Notification.

These are sent whenever someone subscribes or unsubscribes from any of your premium SMS products. To receive premium sms subscription notifications, you need to set asubscription notification callback URL. From the dashboard select SMS -> SMS Callback URLs -> Subscription Notifications.

# 2. Voice Callback Service

The Voice API sends a notification when a specific event happens. To receive these notifications you need to setup a voice callback URL. From the dashboard select Voice -> Phone Numbers -> Actions -> Callback.

Voice API notifications are sent for;

## 2.1 Outbound calls

These are sent whenever you make a call from a registered SIP number.

## 2.2 Inbound calls

These are sent when a call comes to your virtual or SIP number.

## 2.3 After Input

These are sent whenever an action in your response requires user input (such as GetDigits and Record)

## 2.4 When Call Ends

These are sent after a call ends. This is the final notification and contains some extra information about the call like the cost and duration.

# 3. USSD Callback Service

This service is hit when the user dials a USSD code and every time they respond to a menu. Processing USSD requests using the USSD API is very easy once your account is set up. In particular, you will need to:

- Register a service code with AfricasTalking.
- Register a URL that they can call whenever they get a request from a client coming into their system.

Once you register your callback URL, any requests that they receive belonging to you will trigger a callback that sends the request data to that URL using HTTP POST.

# 4. Airtime Callback Service

## 4.1 Airtime Validation Notifications

The Airtime API provides optional functionality to validate airtime requests from your application. To receive these notifications you need to setup an airtime validation callback URL. From the dashboard select Airtime -> Airtime Callback URLs -> Validation Callback URL. Airtime validation notifications are sent as HTTP POST requests to the validation callback URL provided.

Once you receive a validation callback notification you’ll be expected to send back a JSON response that marks the transaction as Validated or Failed. If validated we will proceed to send the airtime, if failed, we will block the airtime transaction

## 4.2 Airtime Status Notifications

The Airtime API sends delivery status notification from the mobile service provider to your application indicating success or failure of the request. To receive these notifications you need to setup an airtime status callback URL. From the dashboard select Airtime -> Airtime Callback URLs -> Status Callback URL. Status notification content Airtime status notifications are sent as HTTP POST requests to the status callback URL provided.

# 5. Payment Callback Service

The Payment API sends a notification when a specific event happens. To receive these notifications you need to setup a callback URL depending on the type of notification.

## 5.1 BankCheckout

These are sent once the provided bank confirms or rejects the checkout request, or once the checkout request expires.

## 5.2 CardCheckout

These are sent once the card provider confirms or rejects the checkout request, or once the checkout request expires.

## 5.3 MobileCheckout

These are sent once the mobile subscriber confirms or rejects the checkout request, or once the checkout request expires.

## 5.4 MobileC2B

These are sent once funds are moved from the mobile subscriber’s account to your payment wallet.

## 5.5 MobileB2C

These are sent once funds are successfully moved from your payment wallet to the mobile subscriber’s account.
If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

## 5.6 MobileB2B

These are sent once funds are successfully moved from your payment wallet to the recipeints business account. If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

## 5.7 BankTransfer

These are sent once funds are successfully moved from your payment wallet to the provided bank account account. If AT are not able to complete the transaction, they will refund your payment wallet with the value of the transaction, and also refund your Africa’s Talking Stash with any transaction Fees.

## 5.8 WalletTransfer

These are sent once funds are successfully moved from your payment wallet to the target payment wallet.

## 5.9 UserStashTopup

These are sent once funds are successfully moved from your payment wallet to your Africa’s Talking Stash.

# 6. IoT Callback Service

The IoT API sends a notification when a device publish event occurs. To receive these notifications you need to setup a callback URL depending on the type of notification. When creating your device group, you have the option to supply your application callback URL. This is the remote endpoint to which device messages will be sent.

# SMS

Send and receive SMS to more than 300 million mobile subscribers across Africa.

## How the SMS service works

1. SMS POST request to AT gateway
2. AT gateway returns JSON response
3. AT sends SMS to user.
4. Status update from Telco to AT gateway.
5. AT gateway sends status update to callback url.

## Send SMS

Send SMS through your application by making a HTTP POST request to the following endpoints:

### Endpoints

Live: https://api.africastalking.com/version1/messaging
Sandbox: https://api.sandbox.africastalking.com/version1/messaging

## Fetch Messages

You can incrementally fetch your application inbox. To do so, you make a HTTP GET request to the following endpoints:

### Endpoints

Live: https://api.africastalking.com/version1/messaging
Sandbox: https://api.sandbox.africastalking.com/version1/messaging

## SMS Notifications

The SMS API sends a notification when a specific event happens. To receive these notifications you need to setup a callback URL depending on the type of notification you would like to receive.

### Types of SMS Notifications

SMS API notifications are sent for various SMS categories as shown below:

- Delivery reports: Sent whenever the mobile service provider confirms or rejects delivery of a message.
- Incoming messages: Sent whenever a message is sent to any of your registered shortcodes.
- Bulk SMS Opt Out: Sent whenever a user opts out of receiving messages from your alphanumeric sender ID.
- Subscription Notifications: Sent whenever someone subscribes or unsubscribes from any of your premium SMS products.

### Delivery Reports
To receive delivery reports, you need to set a delivery report callback URL. From the dashboard select SMS -> SMS Callback URLs -> Delivery Reports.

### Incoming Messages
To receive incoming messages, you need to set an incoming messages callback URL. From the dashboard select SMS -> SMS Callback URLs -> Incoming Messages.

### Bulk SMS Opt Out
To receive bulk sms opt out notifications, you need to set a bulk sms opt out callback URL. From the dashboard select SMS -> SMS Callback URLs -> Bulk SMS Opt Out.

The instructions on how to opt out are automatically appended to the first message you send to the mobile subscriber. From then onwards, any other message will be sent ‘as is’ to the subscriber.

### Subscription Notification
To receive premium sms subscription notifications, you need to set a subscription notification callback URL. From the dashboard select SMS -> SMS Callback URLs -> Subscription Notifications.
# 1. SMS Notifications

The SMS API sends a notification when a specific event happens. To receive these notifications you need to setup a callback URL depending on the type of notification you would like to receive.

## Types of SMS Notifications

SMS API notifications are sent for various SMS categories as shown below:

- Delivery reports: Sent whenever the mobile service provider confirms or rejects delivery of a message.

- Incoming messages: Sent whenever a message is sent to any of your registered shortcodes.

- Bulk SMS Opt Out: Sent whenever a user opts out of receiving messages from your alphanumeric sender ID.

- Subscription Notifications: Sent whenever someone subscribes or unsubscribes from any of your premium SMS products.

## 4.4.1 Delivery Reports

To receive delivery reports, you need to set a delivery report callback URL. From the dashboard select SMS -> SMS Callback URLs -> Delivery Reports.

## 4.4.2 Incoming Messages

To receive incoming messages, you need to set an incoming messages callback URL. From the dashboard select SMS -> SMS Callback URLs -> Incoming Messages.

## 4.4.3 Bulk SMS Opt Out

To receive bulk sms opt out notifications, you need to set a bulk sms opt out callback URL. From the dashboard select SMS -> SMS Callback URLs -> Bulk SMS Opt Out.

The instructions on how to opt out are automatically appended to the first message you send to the mobile subscriber. From then onwards, any other message will be sent ‘as is’ to the subscriber.

## 4.4.4 Subscription Notification

To receive premium sms subscription notifications, you need to set a subscription notification callback URL. From the dashboard select SMS -> SMS Callback URLs -> Subscription Notifications.

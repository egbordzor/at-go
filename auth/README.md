# Authentication

Requests made to Africa’s Talking APIs must be authenticated. There are two ways to do this:

1. Authenticating using your API Key and Username.
2. Authenticating using an Auth Token.

## 1. Authenticating using your API Key and Username

### 1.1 Your Username

When working with the sandbox ( Africa’s Talking development environment ), the username is always sandbox. When in the live environment, this is the specific username of the application making the request.

### 1.2 Your API Key

When a new API key is generated, you can no longer use the old one. After you generate your API key, Africa’s Talking strongly adviseS that you copy it and keep it somewhere safe. It will not be displayed again because Africa’s Talking does not log or save your API Key for security reasons. If you lose it, you’ll have to generate a new one.

### 1.3 Request Headers

|  Parameter   |    Type    |                         Description                          |
| :----------: | :--------: | :----------------------------------------------------------: |
|    apiKey    |   String   |             Africa’s Talking application apiKey.             |
|              | (Required) |                                                              |
|              |            |                                                              |
| Content-Type |   String   |                  The requests content type.                  |
|              | (Required) | Can be application/x-www-form-urlencoded or application/json |
|              |            |                                                              |
|    Accept    |   String   |                 The requests response type.                  |
|              |            |         Can be application/json or application/xml.          |
|              | (Optional) |                 Defaults to application/xml                  |

### 1.4 Making an API call

You need to include the API key in the request header as a field called apiKey. The place where the username should be included depends on the type of request.

- For GET requests e.g. fetch messages, the username should be passed as a query parameter.

- For POST requests in which parameters are sent as a url encoded form e.g. in sending SMS, the username should be included as one of the parameters within the form.

- For POST requests that require JSON in the request body e.g. in mobile checkout, the username should be included in the JSON sent in the body of the request.

## 2. Authenticating with an Auth Token

For instances where it may not be possible to include your APIKey in your application such as in a mobile application, Africa’s Talking provide a way to authenticate using temporary auth tokens.

### 2.1 Getting the token

To generate the auth token, make a POST request to https://api.africastalking.com/auth-token/generate with your username and API Key. This request should be made from your server as you should not inclue your apiKey in client code.

You will receive a JSON response that looks like this:

```json
{
  "token": "ATtkn_abcdefghijklmnopqrstuvwxyz",
  "lifetimeInSeconds": 3600
}
```

You will be able to use that token to make API calls. The token will be valid for value of lifetimeInSeconds and you should generate a new token before it expires.

You need to include the Auth Token in the request header as a field called authToken. The place where the username should be included depends on the type of request you’re making.

### 2.2 Request Headers

|  Parameter   |    Type     |                         Description                         |
| :----------: | :---------: | :---------------------------------------------------------: |
|  authToken   |   String    |                    Generated Auth Token.                    |
|              | (Required ) |                                                             |
|              |             |                                                             |
| Content-Type |   String    |                 The requests content type.                  |
|              | (Required)  | Can be application/x-www-form-urlencoded or application/xml |
|              |             |                                                             |
|    Accept    |   String    |                 The requests response type.                 |
|              | (Optional)  |         Can be application/json or application/xml.         |
|              |             |                 Defaults to application/xml                 |

## 3. Idempotent Requests

Africa’s Talking APIs protect your application against cases where you might end up sending unintended repeat requests. This could be caused by a communication breakdown (mainly network issues) or your application having broken logic.

For example, imagine a scenario where you initiate a POST request to top up phoneNumber1 with KES 100 worth of credits. However, due to network issues, you do not receive a valid response from our APIs, even though the airtime was delivered to phoneNumber1. In this case, you might actually want to retry the POST request but also ensure that we do not send the airtime again.

However, you might run into cases where you actually want to send the same request more than once. In that case, you can set a special header Idempotency-Key to a unique value of your choosing and resend the request. If you send another request with the same Idempotency-Key within a given period of time, we will respond with a failure status notifying you that we have detected a duplicate request.

You can add idempotent keys to your requests to ensure that we send a request from your application once. This feature is currently supported for Airtime APIs and Payment Disbursement APIs.

### 3.1 Request Headers

|    Parameter    |    Type    |                    Description                    |
| :-------------: | :--------: | :-----------------------------------------------: |
| Idempotency-Key |   String   |  A unique value of your choosing that identifies  |
|                 | (Optional) | a request sent Africa’s Talking APIs e.g req-1234 |

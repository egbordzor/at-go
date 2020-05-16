# Authentication

All requests must be authenticated. There are currently two ways to do this.

> 1. Authenticating using your API Key and Username

API keys must be included in the request header, as a field called `apiKey`. The place where the username should be included depends on the type of request you’re making.

## Request Headers

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

> 2. Authenticating using an Auth Token

You need to include the Auth Token in the request header as a field called `authToken`. The place where the username should be included depends on the type of request you’re making.

## Request Headers

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

> 3. Username Scenarios

- For GET requests: The username is passed as a query parameter.
- For POST requests, where parameters are sent as a url encoded form, the username is included as one of the parameters within the form.
- For POST requests, that require JSON in the request body, the username should be included in the JSON sent in the body of the request.

> 4. Idempotent Requests

Africa’s Talking APIs protect applications against cases where you might end up sending unintended repeat requests. This could be caused by a communication breakdown (mainly network issues) or your application having broken logic.

For example, imagine a scenario where you initiate a POST request to top up phoneNumber1 with KES 100 worth of credits. However, due to network issues, you do not receive a valid response from our APIs, even though the airtime was delivered to phoneNumber1. In this case, you might actually want to retry the POST request but also ensure that we do not send the airtime again.

However, you might run into cases where you actually want to send the same request more than once. In that case, you can set a special header Idempotency-Key to a unique value of your choosing and resend the request. If you send another request with the same Idempotency-Key within a given period of time, we will respond with a failure status notifying you that we have detected a duplicate request.

You can add idempotent keys to your requests to ensure that we send a request from your application once. This feature is currently supported for Airtime APIs and Payment Disbursement APIs.

## Request Headers

|    Parameter    |    Type    |                    Description                    |
| :-------------: | :--------: | :-----------------------------------------------: |
| Idempotency-Key |   String   |  A unique value of your choosing that identifies  |
|                 | (Optional) | a request sent Africa’s Talking APIs e.g req-1234 |

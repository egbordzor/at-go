# Authentication

All requests must be authenticated. There are currently two ways to do this.

## 1. Authenticating using your API Key and Username

API keys must be included in the request header, as a field called `apiKey`. The place where the username should be included depends on the type of request you’re making.

### Request Headers

|  Parameter   |       Type        | Description                                                    |
| :----------: | :---------------: | :------------------------------------------------------------- |
|    apiKey    | String (Required) | Africa’s Talking application apiKey.                           |
| Content-Type | String (Required) | The requests content type.                                     |
|              |                   | - Can be application/x-www-form-urlencoded or application/json |
|    Accept    | String (Optional) | The requests response type.                                    |
|              |                   | - Can be application/json or application/xml.                  |
|              |                   | - Defaults to application/xml                                  |

## 2. Authenticating using an Auth Token

You need to include the Auth Token in the request header as a field called `authToken`. The place where the username should be included depends on the type of request you’re making.

### Request Headers

|  Parameter   |        Type        | Description                                                   |
| :----------: | :----------------: | :------------------------------------------------------------ |
|  authToken   | String (Required ) | Generated Auth Token.                                         |
| Content-Type | String (Required)  | The requests content type.                                    |
|              |                    | - Can be application/x-www-form-urlencoded or application/xml |
|    Accept    | String (Optional)  | The requests response type.                                   |
|              |                    | - Can be application/json or application/xml.                 |
|              |                    | - Defaults to application/xml                                 |

## 3. Username Scenarios

- For GET requests: The username is passed as a query parameter.
- For POST requests, where parameters are sent as a url encoded form, the username is included as one of the parameters within the form.
- For POST requests, that require JSON in the request body, the username should be included in the JSON sent in the body of the request.

## 4. Idempotent Requests

Africa’s Talking APIs protect applications against cases where you might end up sending unintended repeat requests. This could be caused by a communication breakdown (mainly network issues) or your application having broken logic.

You can add idempotent keys to your requests to ensure that we send a request from your application once. This feature is currently supported for Airtime APIs and Payment Disbursement APIs.

### Request Headers

|    Parameter    |       Type        | Description                                       |
| :-------------: | :---------------: | :------------------------------------------------ |
| Idempotency-Key | String (Optional) | A unique value of your choosing that identifies   |
|                 |                   | a request sent Africa’s Talking APIs e.g req-1234 |

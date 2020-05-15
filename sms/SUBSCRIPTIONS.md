# 1. Premium Subscriptions

## 1.1 Create Subscription

Two endpoints are required in order for your application to complete a premium sms create subscription request. These are:

- **Checkout Token request:** The checkoutToken is used to authorize a premium sms subscription. It is required in order to subscribe a phone number.

- **Subscription request:** Once a checkoutToken is received it should be sent along with the rest of the parameters to initiate the actual subscription request.

## 1.1.1 Generate a checkout token

Generate a checkoutToken by making a HTTP POST request to the following endpoints:

- Live: https://api.africastalking.com/checkout/token/create
- Sandbox: https://api.sandbox.africastalking.com/checkout/token/create

### Request parameters

This is an open endpoint and does not need your authentication credentials.

|  Parameter  |    Type    |                      Description                       |
| :---------: | :--------: | :----------------------------------------------------: |
| phoneNumber |   String   | The phone number you want to create a subscription for |
|             | (Required) |                                                        |

### API response

The body of the response will be a JSON object containing the following fields:

|  Parameter  |  Type  |                 Description                 |
| :---------: | :----: | :-----------------------------------------: |
| description | String | A description of the status of the request. |
|             |        |   Possible values are: Success or Failed    |
|    token    | String |        The checkout token to be used        |

Below is a sample generate checkout token response for a successful request:

```json
{
  "description": "Success",
  "token": "CkTkn_SampleCkTknId123"
}
```

This token shall be expected later when initiating the actual create subscription request. The tokens expire after 10 minutes and are limited to 2 tokens for a 5 minute window for each source IP Address.

## 1.1.2 Subscribe a phone number

You can subscribe a phone number by making a HTTP POST request to the following endpoints:

- Live: https://api.africastalking.com/version1/subscription/create
- Sandbox: https://api.sandbox.africastalking.com/version1/subscription/create

### Request Parameters

In addition to the standard request headers, the body of the request should contain the following fields:

|   Parameter   |      Type       |                              Description                               |
| :-----------: | :-------------: | :--------------------------------------------------------------------: |
|   username    | String Required |                 Africa’s Talking application username.                 |
|   shortCode   | String Required |             The premium short code mapped to your account.             |
|    keyword    | String Required | The premium keyword under the above short code mapped to your account. |
|  phoneNumber  | String Required |                    The phoneNumber to be subscribed                    |
| checkoutToken | String Required |       This is a token used to validate the subscription request        |

### API response

The body of the response will be a JSON object containing the following fields:

|  Parameter  |  Type  |                                         Description                                         |
| :---------: | :----: | :-----------------------------------------------------------------------------------------: |
|   status    | String | Indicates whether the prompt to subscribe to this shortcode was successfully raised or not. |
|             |        |                           Possible values are: Success or Failed                            |
| description | String |                  Describes the status of the create subscription request.                   |

Below is a sample premium sms create subscription response for a successful request:

```json
{
  "status": "Success",
  "description": "Waiting for user input"
}
```

## 1.2 Fetch Subscriptions

You can incrementally fetch your premium sms subscriptions. To do so, you make a HTTP GET request to the following endpoints:

- Live: https://api.africastalking.com/version1/subscription
- Sandbox: https://api.sandbox.africastalking.com/verson1/subscription

### Request parameters

In addition to the standard request headers, the body of the request should contain the following fields:

|   Parameter    |    Type    |                              Description                              |
| :------------: | :--------: | :-------------------------------------------------------------------: |
|    username    |   String   |                Africa’s Talking application username.                 |
|                | (Required) |                                                                       |
|   shortCode    |   String   |             The premium short code mapped to your account             |
|                | (Required) |
|    keyword     |   String   | The premium keyword under the above short code mapped to your account |
|                | (Required) |
| lastReceivedId |   String   |          ID of the subscription you believe to be your last.          |
|                | (Optional) |                  Set it to 0 to for the first time.                   |

### API response

The body of the response will be a JSON object containing the following fields:

|   Parameter   | Type |                         Description                          |
| :-----------: | :--: | :----------------------------------------------------------: |
| Subscriptions | List |         A list of subscriptions made to the product.         |
|               |      |      Each subscription will have the following fields:       |
|               |      |           `id` Integer: The id of the subscription           |
|               |      | `number` String: The phone number subscribed to the product. |
|               |      |   `Date` String: Timestamp when the subscription was made.   |

Below is a sample fetch subscriptions response for a successful request:

```json
{
  "responses": [
    {
      "id": 100,
      "phoneNumber": "+254711XXXYYY",
      "date": "Timestamp"
    },
    {
      "id": 200,
      "phoneNumber": "+254733YYYZZZ",
      "date": "Timestamp"
    }
  ]
}
```

## 1.3. Delete Subscription

Delete a premium sms subscription by making a HTTP POST request to the following endpoints:

- Live: https://api.africastalking.com/version1/subscription/delete
- Sandbox: https://api.sandbox.africastalking.com/version1/subscription/delete

### Request parameters

In addition to the standard request headers, the body of the request should contain the following fields:

|  Parameter  |      Type       |                              Description                              |
| :---------: | :-------------: | :-------------------------------------------------------------------: |
|  username   |     String      |                Africa’s Talking application username.                 |
|             |   (Required)    |                                                                       |
|  shortCode  | String Required |             The premium short code mapped to your account             |
|             |   (Required)    |                                                                       |
|   keyword   | String Required | The premium keyword under the above short code mapped to your account |
|             |   (Required)    |                                                                       |
| phoneNumber | String Required |                  The phoneNumber to be unsubscribed                   |

### API response

The body of the response will be a JSON object containing the following fields:

|  Parameter  |  Type  |                               Description                                |
| :---------: | :----: | :----------------------------------------------------------------------: |
|   status    | String | Indicates whether the phone number was successfully unsubscribed or not. |
|             |        |                  Possible values are: Success or Failed                  |
| description | String |         Describes the status of the delete subscription request.         |

Below is a sample delete subscription response for a successful request:

```json
{
  "status": "Success",
  "description": "Succeeded"
}
```

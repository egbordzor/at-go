package atgo

import (
	"context"
)

type (
	TopUp interface {

		// Move money from a Payment Product to an application stash.
		// An application stash is the wallet that funds your service usage expenses.
		TopupStash(ctx context.Context, p *TopupStashPayload) (res *TopupStashResponse, err error)
	}
)

// TopupStashPayload is the payload type of the africastalking service
// TopupStash method.
type TopupStashPayload struct {
	// Africa’s Talking application username.
	Username string
	// Africa’s Talking Payment product initiating transaction.
	ProductName string
	// 3-digit ISO format currency code.
	CurrencyCode string
	// Amount application will be topped up with.
	Amount float64
	// Metadata associated with the request.
	Metadata map[string]string
}

// TopupStashResponse is the result type of the africastalking service
// TopupStash method.
type TopupStashResponse struct {
	// Corresponds to the status of the request
	Status string
	// A detailed description of the request status.
	Description string
	// Unique ID for successful requests.
	TransactionID string
}

// TopupStashRequestBody is the type of the "africastalking" service
// "TopupStash" endpoint HTTP request body.
type TopupStashRequestBody struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Africa’s Talking Payment product initiating transaction.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// 3-digit ISO format currency code.
	CurrencyCode string `form:"currencyCode" json:"currencyCode" xml:"currencyCode"`
	// Amount application will be topped up with.
	Amount float64 `form:"amount" json:"amount" xml:"amount"`
	// Metadata associated with the request.
	Metadata map[string]string `form:"metadata" json:"metadata" xml:"metadata"`
}

// TopupStashResponseBody is the type of the "africastalking" service
// "TopupStash" endpoint HTTP response body.
type TopupStashResponseBody struct {
	// Corresponds to the status of the request
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A detailed description of the request status.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Unique ID for successful requests.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

package airtyme

import (
	"context"
)

type (
	Airtime interface {

		// Send Airtime.
		SendAirtime(ctx context.Context, p *AirtimePayload) (res *AirtimeResponse, err error)
	}
)

// AirtimePayload is the payload type of the SendAirtime method.
type AirtimePayload struct {

	// Africa’s Talking application username.
	Username   string              `json:"username"`
	Recipients []AirtimeRecipients `json:"recipients"`
}

// A url encoded json list of Recipients
type AirtimeRecipients struct {

	// Phone number that will be topped up.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Value of airtime to send together with the currency code.
	Amount string `json:"amount ,omitempty"`
}

// AirtimeResponse is the result type of the SendAirtime method.
type AirtimeResponse struct {

	// Number of requests sent to the provider
	NumSent int `json:"numSent,omitempty"`

	// Total value of airtime sent to the provider.
	TotalAmount string `json:"totalAmount,omitempty"`

	// Total discount applied on the airtime.
	TotalDiscount string `json:"totalDiscount,omitempty"`

	Responses []AirtimeEntry `json:"responses,omitempty"`

	// Error message if the ENTIRE request was rejected by the API.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type AirtimeEntry struct {

	// Phone number for this transaction.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Value of airtime requested.
	Amount string `json:"amount,omitempty"`

	// Discount applied to the requested airtime amount.
	Discount string `json:"discount,omitempty"`

	// Status of the request associated to this phone number
	Status string `json:"status,omitempty"`

	// Unique ID for  the request associated to this phone number
	RequestID string `json:"requestId,omitempty"`

	// Error message for the request associated to this phone number.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

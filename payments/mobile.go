package atgo

import (
	"context"
)

type (
	Mobile interface {

		// Initiate C2B payments on a mobile subscriber’s device.
		MobileCheckout(ctx context.Context, p *MobileCheckoutPayload) (res *MobileCheckoutResponse, err error)

		// Send payments to mobile subscribers from your Payment Wallet.
		MobileB2C(ctx context.Context, p *MobileB2CPayload) (res *MobileB2CResponse, err error)

		// Send payments to businesses e.g banks from your Payment Wallet.
		MobileB2B(ctx context.Context, p *MobileB2BPayload) (res MobileB2BResponse, err error)
	}
)

// MobileCheckoutPayload is the payload type of the africastalking service
// MobileCheckout method.
type MobileCheckoutPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Africa’s Talking Payment product to initiate this transaction.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// Provider channel the payment will be initiated from.
	ProviderChannel string `form:"providerChannel,omitempty" json:"providerChannel,omitempty" xml:"providerChannel,omitempty"`
	// Phone number of the client that will complete this transaction.
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
	// 3-digit ISO format currency code.
	CurrencyCode string `form:"currencyCode" json:"currencyCode" xml:"currencyCode"`
	// Amount client is expected to confirm.
	Amount float64 `form:"amount" json:"amount" xml:"amount"`
	// Map of any metadata associates with the request
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// MobileCheckoutResponse is the result type of the africastalking service
// MobileCheckout method.
type MobileCheckoutResponse struct {
	// Status of the request
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Detailed description of the request status.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Unique ID that our API generates for successful requests.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// Provider channel used to initiate this transaction.
	ProviderChannel string `form:"providerChannel,omitempty" json:"providerChannel,omitempty" xml:"providerChannel,omitempty"`
}

// MobileB2CPayload is the payload type of the africastalking service MobileB2C
// method.
type MobileB2CPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Africa’s Talking Payment Product to initiate this transaction.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// A list of B2C Mobile Recipients
	Recipients []MobileRecipients `form:"recipients" json:"recipients" xml:"recipients"`
}

// MobileB2CResponse is the result type of the africastalking service MobileB2C
// method.
type MobileB2CResponse struct {
	// Number of B2C transactions that were successfully queued.
	NumQueued int `form:"numQueued,omitempty" json:"numQueued,omitempty" xml:"numQueued,omitempty"`
	// Total value of all the transactions that were successfully queued.
	TotalValue string `form:"totalValue,omitempty" json:"totalValue,omitempty" xml:"totalValue,omitempty"`
	// Total transaction fee charged for all the transactions that were
	// successfully queued.
	TotalTransactionFee string `form:"totalTransactionFee,omitempty" json:"totalTransactionFee,omitempty" xml:"totalTransactionFee,omitempty"`
	// A list of B2C entries
	Entries []B2CEntry `form:"entries,omitempty" json:"entries,omitempty" xml:"entries,omitempty"`
	// Error message if the ENTIRE request was rejected by the API
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

// MobileB2BPayload is the payload type of the africastalking service MobileB2B
// method.
type MobileB2BPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// Africa’s Talking Payment Product initiating transaction.
	ProductName string `form:"productName,omitempty" json:"productName,omitempty" xml:"productName,omitempty"`
	// Provider used to process the B2C request.
	Provider string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// Transfer type of the payment.
	TransferType string `form:"transferType,omitempty" json:"transferType,omitempty" xml:"transferType,omitempty"`
	// 3-digit ISO format currency code
	CurrencyCode string `form:"currencyCode,omitempty" json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	// Amount client is expected to confirm.
	Amount float64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// Name or number of the channel receiving payment by the provider.
	DestinationChannel string `form:"destinationChannel,omitempty" json:"destinationChannel,omitempty" xml:"destinationChannel,omitempty"`
	// Account name used by the business to receive money on the provided
	// destinationChannel.
	DestinationAccount string `form:"destinationAccount,omitempty" json:"destinationAccount,omitempty" xml:"destinationAccount,omitempty"`
	// A map of any metadata associated with the request.
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// MobileB2BResponse is the result type of the africastalking service MobileB2B
// method.
type MobileB2BResponse struct {
	// The status of the B2B transaction.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A unique id that our API generates for successful requests.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// Transaction fee charged by Africa’s Talking for this transaction.
	TransactionFee string `form:"transactionFee,omitempty" json:"transactionFee,omitempty" xml:"transactionFee,omitempty"`
	// Provider channel which facilitated the payment.
	ProviderChannel string `form:"providerChannel,omitempty" json:"providerChannel,omitempty" xml:"providerChannel,omitempty"`
	// A more descriptive error message for the status of this transaction.
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

// List of Recipient elements of a B2C transaction request.
type MobileRecipients struct {
	// Name of the B2C transaction recipient.
	Name string
	// Phone number of the B2C transaction recipient.
	PhoneNumber string
	// 3-digit ISO format currency code.
	CurrencyCode string
	// Amount that the client is expected to confirm.
	Amount float64
	// Channel payment will be made from.
	ProviderChannel string
	// Purpose of the payment.
	Reason string
	// Map of metadata associated with the request.
	Metadata map[string]string
}

type B2CEntry struct {
	// The phone number of the B2C transaction recipient.
	PhoneNumber string
	// The status of the B2C transaction.
	Status string
	// Unique ID that our API generates for successful requests.
	TransactionID string
	// Provider used to process the B2C request.
	Provider string
	// Channel used to process the B2C request.
	ProviderChannel string
	// Value sent to the mobile subscriber.
	Value string
	// Transaction fee charged by Africa’s Talking for this transaction.
	TransactionFee string
	// A more descriptive error message for the status of this transaction.
	ErrorMessage string
}

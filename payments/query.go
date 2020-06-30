package atgo

import (
	"context"
)

type (
	PaymentQueries interface {

		// Fetch transactions of a particular payment product.
		FindTransaction(ctx context.Context, p *FindTransactionPayload) (res *FindTransactionResponse, err error)

		// Fetch transactions of a particular payment product.
		FetchProductTransactions(ctx context.Context, p *ProductTransactionsPayload) (res *ProductTransactionsResponse, err error)

		// Fetch your wallet transactions
		FetchWalletTransactions(ctx context.Context, p *WalletTransactionsPayload) (res *WalletTransactionsResponse, err error)

		// Fetch your wallet balance
		FetchWalletBalance(ctx context.Context, p *WalletBalancePayload) (res *WalletBalanceResponse, err error)
	}
)

// FindTransactionPayload is the payload type of the africastalking service
// FindTransaction method.
type FindTransactionPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// ID of the transaction you would like to find.
	TransactionID string `form:"transactionId" json:"transactionId" xml:"transactionId"`
}

// FindTransactionResponse is the result type of the africastalking service
// FindTransaction method.
type FindTransactionResponse struct {
	// Status of the request
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Details of the transaction.
	Data TransactionResponse `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// A message detailing what happened with a failed request.
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

// ProductTransactionsPayload is the payload type of the africastalking service
// FetchProductTransactions method.
type ProductTransactionsPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Name of the payment product to fetch.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// Number of the page you’d like to read results from.
	PageNumber int `form:"pageNumber" json:"pageNumber" xml:"pageNumber"`
	// Number of transaction results you would like for this query.
	Count int `form:"count" json:"count" xml:"count"`
	// Transaction start date; in the format YYYY-MM-DD
	StartDate string `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Transaction end date; in the format YYYY-MM-DD
	EndDate string `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Transaction category you would like to consider.
	Category string `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Transaction provider you would like to consider.
	Provider string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// Transaction status you would like to consider
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Transaction source you would like to consider.
	Source string `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// Transaction destination you would like to consider.
	Destination string `form:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Transaction provider channel you would like to consider.
	ProviderChannel string `form:"providerChannel,omitempty" json:"providerChannel,omitempty" xml:"providerChannel,omitempty"`
}

// ProductTransactionsResponse is the result type of the africastalking service
// FetchProductTransactions method.
type ProductTransactionsResponse struct {
	// The status of the request.
	Status    string                `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	Responses []TransactionResponse `form:"responses,omitempty" json:"responses,omitempty" xml:"responses,omitempty"`
}

// WalletTransactionsPayload is the payload type of the africastalking service
// FetchWalletTransactions method.
type WalletTransactionsPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Number of the page you’d like to read results from.
	PageNumber int `form:"pageNumber" json:"pageNumber" xml:"pageNumber"`
	// Number of transaction results you would like for this query.
	Count int `form:"count" json:"count" xml:"count"`
	// Transaction start date; in the format YYYY-MM-DD
	StartDate string `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Transaction end date; in the format YYYY-MM-DD
	EndDate string `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// List of transaction categories you would like to consider.
	Categories string `form:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty"`
}

// WalletTransactionsResponse is the result type of the africastalking service
// FetchWalletTransactions method.
type WalletTransactionsResponse struct {
	// The status of the request
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// List of response Entry corresponding to a transaction result.
	Responses []WalletEntry `form:"responses,omitempty" json:"responses,omitempty" xml:"responses,omitempty"`
	// A message detailing what happened with a failed request.
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

// WalletBalancePayload is the payload type of the africastalking service
// FetchWalletBalance method.
type WalletBalancePayload struct {
	// Africa’s Talking application username
	Username string `form:"username" json:"username" xml:"username"`
}

// WalletBalanceResponse is the result type of the africastalking service
// FetchWalletBalance method.
type WalletBalanceResponse struct {
	// The status of the request.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Balance of the payment wallet.
	Balance string `form:"balance,omitempty" json:"balance,omitempty" xml:"balance,omitempty"`
	// A message detailing what happened with a failed request.
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

type TransactionResponse struct {
	// Metadata sent by your application when it initiated this transaction.
	RequestMetadata map[string]string `form:"requestMetadata,omitempty" json:"requestMetadata,omitempty" xml:"requestMetadata,omitempty"`
	// Type of party providing funds for this transaction (the Debit Party).
	SourceType string `form:"sourceType,omitempty" json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Uniquely identifies the party providing the funds for this transaction
	Source string `form:"source,omitempty" json:"source,omitempty" xml:"source,omitempty"`
	// Payment provider that facilitated this transaction
	Provider string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// Identifies party receiving funds in this transaction (the Credit Party)
	DestinationType string `form:"destinationType,omitempty" json:"destinationType,omitempty" xml:"destinationType,omitempty"`
	// Contains a detailed description of this transaction .
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Name or number of the channel used to facilitate this payment by the
	// provider.
	ProviderChannel string `form:"providerChannel,omitempty" json:"providerChannel,omitempty" xml:"providerChannel,omitempty"`
	// Transaction fee charged by Africa’s Talking for this transaction.
	TransactionFee string `form:"transactionFee,omitempty" json:"transactionFee,omitempty" xml:"transactionFee,omitempty"`
	// Unique ID generated by the payment provider for this transaction.
	ProviderRefID string `form:"providerRefId,omitempty" json:"providerRefId,omitempty" xml:"providerRefId,omitempty"`
	// Map of any additional data received from a payment provider.
	ProviderMetadata map[string]string `form:"providerMetadata,omitempty" json:"providerMetadata,omitempty" xml:"providerMetadata,omitempty"`
	// Final status of this transaction.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Identifies the Africa’s Talking Payment Product used.
	ProductName string `form:"productName,omitempty" json:"productName,omitempty" xml:"productName,omitempty"`
	// Category of the payment
	Category string `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Date and time when a successful transaction was completed.
	TransactionDate string `form:"transactionDate,omitempty" json:"transactionDate,omitempty" xml:"transactionDate,omitempty"`
	// Uniquely identifies the party receiving the funds for this transaction.
	Destination string `form:"destination,omitempty" json:"destination,omitempty" xml:"destination,omitempty"`
	// Value being exchanged in this transaction.
	Value string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// Unique transactionId generated for every payment sent and received.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// Date and time when a transaction was accepted by our APIs.
	CreationTime string `form:"creationTime,omitempty" json:"creationTime,omitempty" xml:"creationTime,omitempty"`
}

type WalletEntry struct {
	// Detailed description of this transaction
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The remaining wallet balance after the transaction was processed.
	Balance string `form:"balance,omitempty" json:"balance,omitempty" xml:"balance,omitempty"`
	// Category of the payment
	Category string `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Contains details of the specific transaction
	TransactionData FindTransactionResponse `form:"transactionData,omitempty" json:"transactionData,omitempty" xml:"transactionData,omitempty"`
	// Value being exchanged in this transaction.
	Value string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// A unique transactionId generated for every payment sent and received
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

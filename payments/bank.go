package atgo

import (
	"context"
)

type (
	Bank interface {

		// Collect money into your payment wallet.
		BankCheckout(ctx context.Context, p *BankCheckoutPayload) (res *BankCheckoutResponse, err error)

		// Validate a bank checkout charge request
		BankCheckoutValidate(ctx context.Context, p *BankCheckoutValidatePayload) (res *BankCheckoutValidateResponse, err error)

		// Initiate a bank transfer request.
		BankTransfer(ctx context.Context, p *BankTransferPayload) (res *BankTransferResponse, err error)
	}
)

// BankCheckoutPayload is the payload type of the africastalking service Bank
// Checkout method.
type BankCheckoutPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Africa’s Talking Payment Product to initiate this transaction.
	ProductName string      `form:"productName" json:"productName" xml:"productName"`
	BankAccount BankAccount `form:"bankAccount" json:"bankAccount" xml:"bankAccount"`
	// 3-digit ISO format currency code.
	CurrencyCode string `form:"currencyCode" json:"currencyCode" xml:"currencyCode"`
	// Amount client is expected to confirm.
	Amount float64 `form:"amount" json:"amount" xml:"amount"`
	// Short description of the transaction displayed on the clients statement.
	Narration string `form:"narration" json:"narration" xml:"narration"`
	// A map of any metadata that you would like us to associate with the request.
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// BankCheckoutResponse is the result type of the africastalking service Bank
// Checkout method.
type BankCheckoutResponse struct {
	// This corresponds to the status of the request.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A detailed description of the request status.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Unique ID that our API generates for successful requests.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

// BankCheckoutValidatePayload is the payload type of the africastalking
// service BankCheckoutValidate method.
type BankCheckoutValidatePayload struct {
	// Africa’s Talking application Username.
	Username string `form:"username" json:"username" xml:"username"`
	// The ID of the transaction to be validated.
	TransactionID string `form:"transactionId" json:"transactionId" xml:"transactionId"`
	// One Time Password bank sent to the client.
	Otp string `form:"otp" json:"otp" xml:"otp"`
}

// BankCheckoutValidateResponse is the result type of the africastalking
// service BankCheckoutValidate method.
type BankCheckoutValidateResponse struct {
	// The final status of this request.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A detailed description of the request status.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// BankTransferPayload is the payload type of the africastalking service
// BankTransfer method.
type BankTransferPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Africa’s Talking Payment product to initiate this transaction.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// Transfer Recipients
	Recipients []TransferRecipients `form:"recipients" json:"recipients" xml:"recipients"`
}

// BankTransferResponse is the result type of the africastalking service
// BankTransfer method.
type BankTransferResponse struct {
	// Transfer Entries
	Entries []TransferEntries `form:"entries,omitempty" json:"entries,omitempty" xml:"entries,omitempty"`
	// Error message if the ENTIRE request was rejected by the API.
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

type BankAccount struct {
	// Bank account name.
	AccountName string `form:"accountName" json:"accountName" xml:"accountName"`
	// Bank account number.
	AccountNumber string `form:"accountNumber" json:"accountNumber" xml:"accountNumber"`
	// 6-Digit Integer code for the bank that we allocate.
	BankCode int `form:"bankCode" json:"bankCode" xml:"bankCode"`
	// Date of birth of the account owner.
	DateOfBirth string `form:"dateOfBirth" json:"dateOfBirth" xml:"dateOfBirth"`
}

// A list of Recipient elements each corresponding to a bank transfer
// transaction request.
type TransferRecipients struct {
	// Details of a bank account to receive the bank transfer payment.
	BankAccount string `form:"bankAccount" json:"bankAccount" xml:"bankAccount"`
	// Bank account name.
	AccountName string `form:"accountName,omitempty" json:"accountName,omitempty" xml:"accountName,omitempty"`
	// Bank account number.
	AccountNumber string `form:"accountNumber" json:"accountNumber" xml:"accountNumber"`
	// 6-Digit Integer code for the bank that we allocate.
	BankCode string `form:"bankCode" json:"bankCode" xml:"bankCode"`
	// Date of birth of the account owner.
	DateOfBirth string `form:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty" xml:"dateOfBirth,omitempty"`
	// 3-digit ISO format currency code
	CurrencyCode string `form:"currencyCode" json:"currencyCode" xml:"currencyCode"`
	// Amount client is expected to receive.
	Amount string `form:"amount" json:"amount" xml:"amount"`
	// Short description of the transaction displayed on the clients statement.
	Narration string `form:"narration,omitempty" json:"narration,omitempty" xml:"narration,omitempty"`
	// A map of any metadata associated with the request.
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// A list of bank transfer entries.
type TransferEntries struct {
	// The account number of the bank transfer recipient.
	AccountNumber string `form:"accountNumber,omitempty" json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// The transaction has been accepted and queued for processing by the payment
	// provider.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A unique ID that our API generates for successful requests.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// Transaction fee charged by Africa’s Talking for this transaction.
	TransactionFee string `form:"transactionFee,omitempty" json:"transactionFee,omitempty" xml:"transactionFee,omitempty"`
	// A more descriptive error message for the status of this transaction.
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

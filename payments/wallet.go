package atgo

import (
	"context"
)

type (
	Wallet interface {

		// Transfer money from one Payment Product to another Payment Product hosted on Africa’s Talking.
		WalletTransfer(ctx context.Context, p *WalletTransferPayload) (res *WalletTransferResponse, err error)
	}
)

// WalletTransferPayload is the payload type of the africastalking service
// WalletTransfer method.
type WalletTransferPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Africa’s Talking Payment product to initiate this transaction.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// Unique product code to transfer the funds to.
	TargetProductCode int `form:"targetProductCode" json:"targetProductCode" xml:"targetProductCode"`
	// 3-digit ISO format currency code
	CurrencyCode string `form:"currencyCode" json:"currencyCode" xml:"currencyCode"`
	// Amount application will be topped up with.
	Amount float64 `form:"amount" json:"amount" xml:"amount"`
	// Metadata associated with the request.
	Metadata map[string]string `form:"metadata" json:"metadata" xml:"metadata"`
}

// WalletTransferResponse is the result type of the africastalking service
// WalletTransfer method.
type WalletTransferResponse struct {
	// Corresponds to the status of the request.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A detailed description of the request status.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// A unique id that our API generates for successful requests.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

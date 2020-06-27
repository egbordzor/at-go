// Code generated by goa v3.1.3, DO NOT EDIT.
//
// africastalking endpoints
//
// Command:
// $ goa gen github.com/wondenge/at-go/internal/design -o internal/pkg

package africastalking

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "africastalking" service endpoints.
type Endpoints struct {
	SendBulkSMS              endpoint.Endpoint
	SendPremiumSMS           endpoint.Endpoint
	FetchSMS                 endpoint.Endpoint
	NewCheckoutToken         endpoint.Endpoint
	NewPremiumSubscription   endpoint.Endpoint
	FetchPremiumSubscription endpoint.Endpoint
	PurgePremiumSubscription endpoint.Endpoint
	MakeCall                 endpoint.Endpoint
	TransferCall             endpoint.Endpoint
	UploadMedia              endpoint.Endpoint
	MobileCheckout           endpoint.Endpoint
	MobileB2C                endpoint.Endpoint
	MobileB2B                endpoint.Endpoint
	BankCheckout             endpoint.Endpoint
	BankCheckoutValidate     endpoint.Endpoint
	BankTransfer             endpoint.Endpoint
	CardCheckout             endpoint.Endpoint
	CardCheckoutValidate     endpoint.Endpoint
	WalletTransfer           endpoint.Endpoint
	TopupStash               endpoint.Endpoint
	FindTransaction          endpoint.Endpoint
	FetchProductTransactions endpoint.Endpoint
	FetchWalletTransactions  endpoint.Endpoint
	FetchWalletBalance       endpoint.Endpoint
	SendAirtime              endpoint.Endpoint
	PublishIoT               endpoint.Endpoint
	InitiateAppData          endpoint.Endpoint
	Generate                 endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "africastalking" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		SendBulkSMS:              NewSendBulkSMSEndpoint(s),
		SendPremiumSMS:           NewSendPremiumSMSEndpoint(s),
		FetchSMS:                 NewFetchSMSEndpoint(s),
		NewCheckoutToken:         NewNewCheckoutTokenEndpoint(s),
		NewPremiumSubscription:   NewNewPremiumSubscriptionEndpoint(s),
		FetchPremiumSubscription: NewFetchPremiumSubscriptionEndpoint(s),
		PurgePremiumSubscription: NewPurgePremiumSubscriptionEndpoint(s),
		MakeCall:                 NewMakeCallEndpoint(s),
		TransferCall:             NewTransferCallEndpoint(s),
		UploadMedia:              NewUploadMediaEndpoint(s),
		MobileCheckout:           NewMobileCheckoutEndpoint(s),
		MobileB2C:                NewMobileB2CEndpoint(s),
		MobileB2B:                NewMobileB2BEndpoint(s),
		BankCheckout:             NewBankCheckoutEndpoint(s),
		BankCheckoutValidate:     NewBankCheckoutValidateEndpoint(s),
		BankTransfer:             NewBankTransferEndpoint(s),
		CardCheckout:             NewCardCheckoutEndpoint(s),
		CardCheckoutValidate:     NewCardCheckoutValidateEndpoint(s),
		WalletTransfer:           NewWalletTransferEndpoint(s),
		TopupStash:               NewTopupStashEndpoint(s),
		FindTransaction:          NewFindTransactionEndpoint(s),
		FetchProductTransactions: NewFetchProductTransactionsEndpoint(s),
		FetchWalletTransactions:  NewFetchWalletTransactionsEndpoint(s),
		FetchWalletBalance:       NewFetchWalletBalanceEndpoint(s),
		SendAirtime:              NewSendAirtimeEndpoint(s),
		PublishIoT:               NewPublishIoTEndpoint(s),
		InitiateAppData:          NewInitiateAppDataEndpoint(s),
		Generate:                 NewGenerateEndpoint(s),
	}
}

// Use applies the given middleware to all the "africastalking" service
// endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.SendBulkSMS = m(e.SendBulkSMS)
	e.SendPremiumSMS = m(e.SendPremiumSMS)
	e.FetchSMS = m(e.FetchSMS)
	e.NewCheckoutToken = m(e.NewCheckoutToken)
	e.NewPremiumSubscription = m(e.NewPremiumSubscription)
	e.FetchPremiumSubscription = m(e.FetchPremiumSubscription)
	e.PurgePremiumSubscription = m(e.PurgePremiumSubscription)
	e.MakeCall = m(e.MakeCall)
	e.TransferCall = m(e.TransferCall)
	e.UploadMedia = m(e.UploadMedia)
	e.MobileCheckout = m(e.MobileCheckout)
	e.MobileB2C = m(e.MobileB2C)
	e.MobileB2B = m(e.MobileB2B)
	e.BankCheckout = m(e.BankCheckout)
	e.BankCheckoutValidate = m(e.BankCheckoutValidate)
	e.BankTransfer = m(e.BankTransfer)
	e.CardCheckout = m(e.CardCheckout)
	e.CardCheckoutValidate = m(e.CardCheckoutValidate)
	e.WalletTransfer = m(e.WalletTransfer)
	e.TopupStash = m(e.TopupStash)
	e.FindTransaction = m(e.FindTransaction)
	e.FetchProductTransactions = m(e.FetchProductTransactions)
	e.FetchWalletTransactions = m(e.FetchWalletTransactions)
	e.FetchWalletBalance = m(e.FetchWalletBalance)
	e.SendAirtime = m(e.SendAirtime)
	e.PublishIoT = m(e.PublishIoT)
	e.InitiateAppData = m(e.InitiateAppData)
	e.Generate = m(e.Generate)
}

// NewSendBulkSMSEndpoint returns an endpoint function that calls the method
// "SendBulkSMS" of service "africastalking".
func NewSendBulkSMSEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BulkPayload)
		res, err := s.SendBulkSMS(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedBulkResponse(res, "default")
		return vres, nil
	}
}

// NewSendPremiumSMSEndpoint returns an endpoint function that calls the method
// "SendPremiumSMS" of service "africastalking".
func NewSendPremiumSMSEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PremiumPayload)
		res, err := s.SendPremiumSMS(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPremiumSMSResponse(res, "default")
		return vres, nil
	}
}

// NewFetchSMSEndpoint returns an endpoint function that calls the method
// "FetchSMS" of service "africastalking".
func NewFetchSMSEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FetchMsgPayload)
		res, err := s.FetchSMS(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedFetchMsgResponse(res, "default")
		return vres, nil
	}
}

// NewNewCheckoutTokenEndpoint returns an endpoint function that calls the
// method "NewCheckoutToken" of service "africastalking".
func NewNewCheckoutTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CheckoutTokenPayload)
		res, err := s.NewCheckoutToken(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCheckoutTokenResponse(res, "default")
		return vres, nil
	}
}

// NewNewPremiumSubscriptionEndpoint returns an endpoint function that calls
// the method "NewPremiumSubscription" of service "africastalking".
func NewNewPremiumSubscriptionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*NewSubPayload)
		res, err := s.NewPremiumSubscription(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedNewSubResponse(res, "default")
		return vres, nil
	}
}

// NewFetchPremiumSubscriptionEndpoint returns an endpoint function that calls
// the method "FetchPremiumSubscription" of service "africastalking".
func NewFetchPremiumSubscriptionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FetchSubPayload)
		res, err := s.FetchPremiumSubscription(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedFetchSubResponse(res, "default")
		return vres, nil
	}
}

// NewPurgePremiumSubscriptionEndpoint returns an endpoint function that calls
// the method "PurgePremiumSubscription" of service "africastalking".
func NewPurgePremiumSubscriptionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PurgeSubPayload)
		res, err := s.PurgePremiumSubscription(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPurgeSubResponse(res, "default")
		return vres, nil
	}
}

// NewMakeCallEndpoint returns an endpoint function that calls the method
// "MakeCall" of service "africastalking".
func NewMakeCallEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MakeCallPayload)
		res, err := s.MakeCall(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMakeCallResponse(res, "default")
		return vres, nil
	}
}

// NewTransferCallEndpoint returns an endpoint function that calls the method
// "TransferCall" of service "africastalking".
func NewTransferCallEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CallTransferPayload)
		res, err := s.TransferCall(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCallTransferResponse(res, "default")
		return vres, nil
	}
}

// NewUploadMediaEndpoint returns an endpoint function that calls the method
// "UploadMedia" of service "africastalking".
func NewUploadMediaEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UploadMediaFile)
		return s.UploadMedia(ctx, p)
	}
}

// NewMobileCheckoutEndpoint returns an endpoint function that calls the method
// "MobileCheckout" of service "africastalking".
func NewMobileCheckoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MobileCheckoutPayload)
		res, err := s.MobileCheckout(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMobileCheckoutResponse(res, "default")
		return vres, nil
	}
}

// NewMobileB2CEndpoint returns an endpoint function that calls the method
// "MobileB2C" of service "africastalking".
func NewMobileB2CEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MobileB2CPayload)
		res, err := s.MobileB2C(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMobileB2CResponse(res, "default")
		return vres, nil
	}
}

// NewMobileB2BEndpoint returns an endpoint function that calls the method
// "MobileB2B" of service "africastalking".
func NewMobileB2BEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MobileB2BPayload)
		res, err := s.MobileB2B(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMobileB2BResponse(res, "default")
		return vres, nil
	}
}

// NewBankCheckoutEndpoint returns an endpoint function that calls the method
// "Bank Checkout" of service "africastalking".
func NewBankCheckoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BankCheckoutPayload)
		res, err := s.BankCheckout(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedBankCheckoutResponse(res, "default")
		return vres, nil
	}
}

// NewBankCheckoutValidateEndpoint returns an endpoint function that calls the
// method "BankCheckoutValidate" of service "africastalking".
func NewBankCheckoutValidateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BankCheckoutValidatePayload)
		res, err := s.BankCheckoutValidate(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedBankCheckoutValidateResponse(res, "default")
		return vres, nil
	}
}

// NewBankTransferEndpoint returns an endpoint function that calls the method
// "BankTransfer" of service "africastalking".
func NewBankTransferEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BankTransferPayload)
		res, err := s.BankTransfer(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedBankTransferResponse(res, "default")
		return vres, nil
	}
}

// NewCardCheckoutEndpoint returns an endpoint function that calls the method
// "CardCheckout" of service "africastalking".
func NewCardCheckoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CardCheckoutPayload)
		res, err := s.CardCheckout(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCardCheckoutResponse(res, "default")
		return vres, nil
	}
}

// NewCardCheckoutValidateEndpoint returns an endpoint function that calls the
// method "CardCheckoutValidate" of service "africastalking".
func NewCardCheckoutValidateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CardCheckoutValidatePayload)
		res, err := s.CardCheckoutValidate(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCardCheckoutValidateResponse(res, "default")
		return vres, nil
	}
}

// NewWalletTransferEndpoint returns an endpoint function that calls the method
// "WalletTransfer" of service "africastalking".
func NewWalletTransferEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*WalletTransferPayload)
		res, err := s.WalletTransfer(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedWalletTransferResponse(res, "default")
		return vres, nil
	}
}

// NewTopupStashEndpoint returns an endpoint function that calls the method
// "TopupStash" of service "africastalking".
func NewTopupStashEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*TopupStashPayload)
		res, err := s.TopupStash(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTopupStashResponse(res, "default")
		return vres, nil
	}
}

// NewFindTransactionEndpoint returns an endpoint function that calls the
// method "FindTransaction" of service "africastalking".
func NewFindTransactionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FindTransactionPayload)
		res, err := s.FindTransaction(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedFindTransactionResponse(res, "default")
		return vres, nil
	}
}

// NewFetchProductTransactionsEndpoint returns an endpoint function that calls
// the method "FetchProductTransactions" of service "africastalking".
func NewFetchProductTransactionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ProductTransactionsPayload)
		res, err := s.FetchProductTransactions(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedProductTransactionsResponse(res, "default")
		return vres, nil
	}
}

// NewFetchWalletTransactionsEndpoint returns an endpoint function that calls
// the method "FetchWalletTransactions" of service "africastalking".
func NewFetchWalletTransactionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*WalletTransactionsPayload)
		res, err := s.FetchWalletTransactions(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedWalletTransactionsResponse(res, "default")
		return vres, nil
	}
}

// NewFetchWalletBalanceEndpoint returns an endpoint function that calls the
// method "FetchWalletBalance" of service "africastalking".
func NewFetchWalletBalanceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*WalletBalancePayload)
		res, err := s.FetchWalletBalance(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedWalletBalanceResponse(res, "default")
		return vres, nil
	}
}

// NewSendAirtimeEndpoint returns an endpoint function that calls the method
// "SendAirtime" of service "africastalking".
func NewSendAirtimeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AirtimePayload)
		res, err := s.SendAirtime(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAirtimeResponse(res, "default")
		return vres, nil
	}
}

// NewPublishIoTEndpoint returns an endpoint function that calls the method
// "PublishIoT" of service "africastalking".
func NewPublishIoTEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IoTPayload)
		res, err := s.PublishIoT(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedIoTResponse(res, "default")
		return vres, nil
	}
}

// NewInitiateAppDataEndpoint returns an endpoint function that calls the
// method "InitiateAppData" of service "africastalking".
func NewInitiateAppDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(string)
		res, err := s.InitiateAppData(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUserResponse(res, "default")
		return vres, nil
	}
}

// NewGenerateEndpoint returns an endpoint function that calls the method
// "Generate" of service "africastalking".
func NewGenerateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GeneratePayload)
		res, err := s.Generate(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAccessTokenResponse(res, "default")
		return vres, nil
	}
}

package atgo

import (
	"errors"
	"fmt"
	"github.com/hashicorp/errwrap"
	"go.uber.org/zap"
	"net/http"
)

// Sandbox Endpoints
const (
	SMSTestURL     = "https://api.sandbox.africastalking.com"
	VoiceTestURL   = "https://calls.sandbox.africastalking.com"
	PaymentTestURL = "https://payments.sandbox.africastalking.com"
	AirtimeTestURL = "https://api.sandbox.africastalking.com"
	UserTestURL    = "https://api.sandbox.africastalking.com"
)

// Production Endpoints
const (
	SMSBaseURL     = "https://api.africastalking.com"
	VoiceBaseURL   = "https://calls.africastalking.com"
	PaymentBaseURL = "https://payments.africastalking.com"
	AirtimeBaseURL = "https://api.africastalking.com"
	IoTBaseURL     = "https://iot.africastalking.com"
	UserBaseURL    = "https://api.africastalking.com"
	AuthBaseURL    = "https://api.africastalking.com"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	ClientID     string `json:"client_id"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	UserID       string `json:"user_id"`
}

type (
	Client struct {
		Username        string
		BaseURL         string
		AccessToken     string
		UserID          string
		SMSEndpoint     string
		VoiceEndpoint   string
		PaymentEndpoint string
		AirtimeEndpoint string
		IoTEndpoint     string
		UserEndpoint    string
		AuthEndpoint    string
		APIKey          string
		HTTPClient      *http.Client
		Log             *zap.Logger // *zap.Logger // log the requests.
	}

	ErrorResponseDetail struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
	}

	ErrorResponse struct {
		Response        *http.Response        `json:"-"`
		Name            string                `json:"name"`
		DebugID         string                `json:"debug_id"`
		Message         string                `json:"message"`
		InformationLink string                `json:"information_link"`
		Details         []ErrorResponseDetail `json:"details"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

// NewAfricasTalkingClient returns new Client struct
// Use "test" for Sandbox Environment and "prod" for Production Environment
func NewAfricasTalkingClient(username, apiKey string) (*Client, error) {

	logger, err := zap.NewProduction()
	if err != nil {
		err := errwrap.Wrapf("could not initiate zap logger on client: {{err}}", err)
		logger.Info("error", zap.Error(err))
	}
	defer logger.Sync() // flushes buffer, if any

	if username == "" || apiKey == "" {
		return nil, errors.New("username, apiKey are required to create a Client")
	} else if username == "sandbox" {
		return &Client{
			Username:        "sandbox",
			SMSEndpoint:     SMSTestURL,
			VoiceEndpoint:   VoiceTestURL,
			PaymentEndpoint: PaymentTestURL,
			AirtimeEndpoint: AirtimeTestURL,
			IoTEndpoint:     "", // No Testing Environment for IoT
			UserEndpoint:    UserTestURL,
			AuthEndpoint:    "", // No Testing Environment for Auth
			APIKey:          apiKey,
			HTTPClient:      &http.Client{},
			Log:             logger,
		}, nil
	} else {
		return &Client{
			Username:        username,
			SMSEndpoint:     SMSBaseURL,
			VoiceEndpoint:   VoiceBaseURL,
			PaymentEndpoint: PaymentBaseURL,
			AirtimeEndpoint: AirtimeBaseURL,
			IoTEndpoint:     IoTBaseURL,
			UserEndpoint:    UserBaseURL,
			AuthEndpoint:    AuthBaseURL,
			APIKey:          apiKey,
			HTTPClient:      &http.Client{},
			Log:             logger,
		}, nil
	}
}

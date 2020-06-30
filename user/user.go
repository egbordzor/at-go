package user

import (
	"context"
)

type (
	User interface {

		// Initiate an application data request.
		InitiateAppData(ctx context.Context, p string) (res *UserResponse, err error)
	}
)

type (
	ATAuth interface {

		// Generates a valid auth token
		GenerateToken(ctx context.Context, p *GeneratePayload) (res *AccessTokenResponse, err error)
	}
)

// UserResponse is the result type of the InitiateAppData method.
type UserResponse struct {
	UserData UserData
}

type UserData struct {

	// Your Africa’s Talking application balance.
	Balance string `form:"balance,omitempty" json:"balance,omitempty" xml:"balance,omitempty"`
}

// InitiateAppDataResponseBody is the type of the "africastalking" service
// "InitiateAppData" endpoint HTTP response body.
type InitiateAppDataResponse struct {
	UserData *UserData `form:"UserData,omitempty" json:"UserData,omitempty" xml:"UserData,omitempty"`
}

// GeneratePayload is the payload type of the Generate method.
type GeneratePayload struct {

	// Africa's Talking Username.
	Username string `form:"username" json:"username" xml:"username"`

	// Africa's Talking API Key.
	APIKey string `form:"apiKey" json:"apiKey" xml:"apiKey"`
}

// AccessTokenResponse is the result type of the Generate method.
type AccessTokenResponse struct {

	// Generated Auth Token.
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`

	// Token Lifetime
	LifetimeInSeconds *int `form:"lifetimeInSeconds,omitempty" json:"lifetimeInSeconds,omitempty" xml:"lifetimeInSeconds,omitempty"`
}

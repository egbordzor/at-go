package payments

import (
	"github.com/go-kit/kit/log"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
	"net/http"
	"time"
)

const (

	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = ""

	// APIBaseLive points to the live version of the API
	APIBaseLive = ""
)

type (
	Client struct {
		Client         *http.Client
		APIKey         string
		Username       string
		Sandbox        string
		Live           string
		Token          *africastalking.AccessTokenResponse
		tokenExpiresAt time.Time
		logger         log.Logger // log the requests.
	}
)







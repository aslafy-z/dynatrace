package credentials

import (
	"net/http"
)

// Credentials are able to modify a given HTTP Request in order to authenticate properly.
// One use case is e.g. adding the Authorization Header
type Credentials interface {
	// Authenticate modifies a given HTTP Request in order to ensure proper authentication on the server side
	Authenticate(request *http.Request) error
}

type credentials struct {
	APIToken string `json:"api-token,omitempty"`
}

// New creates a new instance that is able to modify a given HTTP Reqest in orer to ensure proper authentication on the server side
func New(apiToken string) Credentials {
	return &credentials{APIToken: apiToken}
}

// Authenticate modifies a given HTTP Request in order to ensure proper authentication on the server side
func (credentials *credentials) Authenticate(request *http.Request) error {
	request.Header.Set("Authorization", "Api-Token "+credentials.APIToken)
	return nil
}

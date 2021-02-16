package autotags

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
	"github.com/dtcookie/opt"
)

// ServiceClient TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/config/v1"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(config *AutoTag) (*EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if len(opt.String(config.ID)) > 0 {
		return nil, errors.New("You MUST NOT provide an ID within the Notification payload upon creation")
	}

	if bytes, err = cs.client.POST("/autoTags", config, 201); err != nil {
		log.Fatal(err)
		return nil, err
	}
	var stub EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(config *AutoTag) error {
	if len(opt.String(config.ID)) == 0 {
		return errors.New("The Notification doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/autoTags/%s", opt.String(config.ID)), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("Empty ID provided for the Notification to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/autoTags/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*AutoTag, error) {
	if len(id) == 0 {
		return nil, errors.New("Empty ID provided for the Notification to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/autoTags/%s", id), 200); err != nil {
		return nil, err
	}
	var autoTag AutoTag
	if err = json.Unmarshal(bytes, &autoTag); err != nil {
		return nil, err
	}
	return &autoTag, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/autoTags", 200); err != nil {
		return nil, err
	}
	var stubList StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}

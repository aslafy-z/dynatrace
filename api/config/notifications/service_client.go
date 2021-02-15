package notifications

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
func (cs *ServiceClient) Create(config NotificationConfig) (*NotificationConfigStub, error) {
	var err error
	var bytes []byte

	if len(opt.String(config.GetID())) > 0 {
		return nil, errors.New("You MUST NOT provide an ID within the Notification payload upon creation")
	}

	if bytes, err = cs.client.POST("/notifications", config, 201); err != nil {
		log.Fatal(err)
		return nil, err
	}
	var stub NotificationConfigStub
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(config NotificationConfig) error {
	if len(opt.String(config.GetID())) == 0 {
		return errors.New("The Notification doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/notifications/%s", opt.String(config.GetID())), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("Empty ID provided for the Notification to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/notifications/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (NotificationConfig, error) {
	if len(id) == 0 {
		return nil, errors.New("Empty ID provided for the Notification to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/notifications/%s", id), 200); err != nil {
		return nil, err
	}
	var baseConfig BaseNotificationConfig
	if err = json.Unmarshal(bytes, &baseConfig); err != nil {
		return nil, err
	}
	switch baseConfig.Type {
	case Types.Ansibletower:
		var config AnsibleTowerNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Email:
		var config EmailNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Hipchat:
		var config HipChatNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Jira:
		var config JiraNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.OpsGenie:
		var config OpsGenieNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.PagerDuty:
		var config PagerDutyNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.ServiceNow:
		var config ServiceNowNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Slack:
		var config SlackNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Trello:
		var config TrelloNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Victorops:
		var config VictorOpsNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Webhook:
		var config WebHookNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	case Types.Xmatters:
		var config XMattersNotificationConfig
		if err = json.Unmarshal(bytes, &config); err != nil {
			return nil, err
		}
		return &config, nil
	default:
		return nil, fmt.Errorf("Unsupported notification config type %v", baseConfig.Type)
	}
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*NotificationConfigStubListDto, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/notifications", 200); err != nil {
		return nil, err
	}
	var stubList NotificationConfigStubListDto
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}

package customservices_test

import (
	"os"
	"strings"
	"testing"

	api "github.com/dtcookie/dynatrace/api/config"
	cs "github.com/dtcookie/dynatrace/api/config/customservices"
)

func TestInvalidToken(t *testing.T) {
	baseURL := os.Getenv("DT_BASE_URL")
	apiToken := "asdfasdfasd"
	apiClient := cs.NewService(baseURL, apiToken)
	_, err := apiClient.List(cs.Technologies.Java)
	if err == nil {
		t.Error("expected an error when querying with invalid API Token")
	} else if !strings.HasPrefix(err.Error(), "Unauthorized") {
		t.Errorf("expected: Unauthorized ..., actual: %s", err.Error())
	}
}

func TestSum(t *testing.T) {
	var err error
	var stubList *api.StubList
	baseURL := os.Getenv("DT_BASE_URL")
	apiToken := os.Getenv("DT_API_TOKEN")
	apiClient := cs.NewService(baseURL, apiToken)
	if stubList, err = apiClient.List(cs.Technologies.Java); err != nil {
		t.Error(err)
	}
	if stubList == nil {
		t.Error("expected a non nil result value")
	}
}

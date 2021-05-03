package managementzones_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dtcookie/dynatrace/api/config/managementzones"
)

func TestFoo(t *testing.T) {
	c := managementzones.DimensionalRuleCondition{
		Type:  managementzones.ConditionTypes.Dimension,
		Match: managementzones.RuleMatchers.BeginsWith,
		Key:   "asdf",
		// Value: opt.NewString("reini"),
		Value: nil,
	}
	data, err := json.MarshalIndent(&c, "", "  ")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
	var c2 managementzones.DimensionalRuleCondition
	err = json.Unmarshal(data, &c2)
	if err != nil {
		t.Error(err)
	}
	data, err = json.MarshalIndent(&c2, "", "  ")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
}

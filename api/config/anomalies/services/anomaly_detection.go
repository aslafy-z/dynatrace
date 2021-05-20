package services

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/anomalydetection/services/failurerateincrease"
	"github.com/dtcookie/dynatrace/api/config/anomalydetection/services/load"
	"github.com/dtcookie/dynatrace/api/config/anomalydetection/services/responsetimedegredation"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// AnomalyDetection Dynatrace automatically detects service-related performance anomalies such as response time degradations and failure rate increases. Use these settings to configure detection sensitivity, set alert thresholds, or disable alerting for certain services.
type AnomalyDetection struct {
	LoadSpike               *load.SpikeDetection               `json:"loadSpike,omitempty"`     // The configuration of load spikes detection.
	ResponseTimeDegradation *responsetimedegredation.Detection `json:"responseTimeDegradation"` // Configuration of response time degradation detection.
	FailureRateIncrease     *failurerateincrease.Detection     `json:"failureRateIncrease"`     // Configuration of failure rate increase detection.
	LoadDrop                *load.DropDetection                `json:"loadDrop,omitempty"`      // The configuration of load drops detection.
	Metadata                *api.ConfigMetadata                `json:"metadata,omitempty"`      // Metadata useful for debugging
	Unknowns                map[string]json.RawMessage         `json:"-"`
}

func (me *AnomalyDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"load_spikes": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration of load spikes detection",
			Elem:        &hcl.Resource{Schema: new(load.SpikeDetection).Schema()},
		},
		"response_times": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of response time degradation detection",
			Elem:        &hcl.Resource{Schema: new(responsetimedegredation.Detection).Schema()},
		},
		"failure_rates": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of response time degradation detection",
			Elem:        &hcl.Resource{Schema: new(failurerateincrease.Detection).Schema()},
		},
		"load_drops": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration of load drops detection",
			Elem:        &hcl.Resource{Schema: new(load.DropDetection).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *AnomalyDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.LoadSpike != nil {
		if marshalled, err := me.LoadSpike.MarshalHCL(hcl.NewDecoder(decoder, "load_spikes", 0)); err == nil {
			result["load_spikes"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.ResponseTimeDegradation != nil {
		if marshalled, err := me.ResponseTimeDegradation.MarshalHCL(hcl.NewDecoder(decoder, "response_times", 0)); err == nil {
			result["response_times"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.FailureRateIncrease != nil {
		if marshalled, err := me.FailureRateIncrease.MarshalHCL(hcl.NewDecoder(decoder, "failure_rates", 0)); err == nil {
			result["failure_rates"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.LoadDrop != nil {
		if marshalled, err := me.LoadDrop.MarshalHCL(hcl.NewDecoder(decoder, "load_drops", 0)); err == nil {
			result["load_drops"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *AnomalyDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "loadSpike")
		delete(me.Unknowns, "responseTimeDegradation")
		delete(me.Unknowns, "failureRateIncrease")
		delete(me.Unknowns, "loadDrop")
		delete(me.Unknowns, "metadata")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}

	if _, ok := decoder.GetOk("load_spikes.#"); ok {
		me.LoadSpike = new(load.SpikeDetection)
		if err := me.LoadSpike.UnmarshalHCL(hcl.NewDecoder(decoder, "load_spikes", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("response_times.#"); ok {
		me.ResponseTimeDegradation = new(responsetimedegredation.Detection)
		if err := me.ResponseTimeDegradation.UnmarshalHCL(hcl.NewDecoder(decoder, "response_times", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("failure_rates.#"); ok {
		me.FailureRateIncrease = new(failurerateincrease.Detection)
		if err := me.FailureRateIncrease.UnmarshalHCL(hcl.NewDecoder(decoder, "failure_rates", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("load_drops.#"); ok {
		me.LoadDrop = new(load.DropDetection)
		if err := me.LoadDrop.UnmarshalHCL(hcl.NewDecoder(decoder, "load_drops", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *AnomalyDetection) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"loadSpike":               me.LoadSpike,
		"responseTimeDegradation": me.ResponseTimeDegradation,
		"failureRateIncrease":     me.FailureRateIncrease,
		"loadDrop":                me.LoadDrop,
		"metadata":                me.Metadata,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *AnomalyDetection) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"loadSpike":               &me.LoadSpike,
		"responseTimeDegradation": &me.ResponseTimeDegradation,
		"failureRateIncrease":     &me.FailureRateIncrease,
		"loadDrop":                &me.LoadDrop,
		"metadata":                &me.Metadata,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}

package maintenancewindows

import "github.com/dtcookie/gojson"

// Schedule The schedule of the maintenance window.
type Schedule struct {
	Start          string         `json:"start"`                // The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	ZoneID         string         `json:"zoneId"`               // The time zone of the start and end time. Default time zone is UTC.  You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`).
	End            string         `json:"end"`                  // The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	Recurrence     *Recurrence    `json:"recurrence,omitempty"` // The recurrence of the maintenance window.
	RecurrenceType RecurrenceType `json:"recurrenceType"`       // The type of the schedule recurrence.
}

// UnmarshalJSON provides custom JSON deserialization
func (s *Schedule) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, s)
}

// MarshalJSON provides custom JSON serialization
func (s *Schedule) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(s)
}

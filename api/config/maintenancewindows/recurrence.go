package maintenancewindows

import "github.com/dtcookie/gojson"

// Recurrence The recurrence of the maintenance window.
type Recurrence struct {
	DayOfMonth      *int32     `json:"dayOfMonth,omitempty"` // The day of the month for monthly maintenance.  The value of `31` is treated as the last day of the month for months that don't have a 31st day. The value of `30` is also treated as the last day of the month for February.
	DayOfWeek       *DayOfWeek `json:"dayOfWeek,omitempty"`  // The day of the week for weekly maintenance.  The format is the full name of the day in upper case, for example `THURSDAY`.
	DurationMinutes int32      `json:"durationMinutes"`      // The duration of the maintenance window in minutes.
	StartTime       string     `json:"startTime"`            // The start time of the maintenance window in HH:mm format.
}

// UnmarshalJSON provides custom JSON deserialization
func (r *Recurrence) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, r)
}

// MarshalJSON provides custom JSON serialization
func (r *Recurrence) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(r)
}

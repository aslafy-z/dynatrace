package maintenancewindows

// RecurrenceType The type of the schedule recurrence.
type RecurrenceType string

// RecurrenceTypes offers the known enum values
var RecurrenceTypes = struct {
	Daily   RecurrenceType
	Monthly RecurrenceType
	Once    RecurrenceType
	Weekly  RecurrenceType
}{
	"DAILY",
	"MONTHLY",
	"ONCE",
	"WEEKLY",
}

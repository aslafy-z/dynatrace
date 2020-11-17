package maintenancewindows

// DayOfWeek The day of the week for weekly maintenance.
// The format is the full name of the day in upper case, for example `THURSDAY`.
type DayOfWeek string

// DayOfWeeks offers the known enum values
var DayOfWeeks = struct {
	Friday    DayOfWeek
	Monday    DayOfWeek
	Saturday  DayOfWeek
	Sunday    DayOfWeek
	Thursday  DayOfWeek
	Tuesday   DayOfWeek
	Wednesday DayOfWeek
}{
	"FRIDAY",
	"MONDAY",
	"SATURDAY",
	"SUNDAY",
	"THURSDAY",
	"TUESDAY",
	"WEDNESDAY",
}

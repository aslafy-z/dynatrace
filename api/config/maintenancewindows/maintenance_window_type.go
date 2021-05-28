package maintenancewindows

// MaintenanceWindowType The type of the maintenance: planned or unplanned.
type MaintenanceWindowType string

// MaintenanceWindowTypes offers the known enum values
var MaintenanceWindowTypes = struct {
	Planned   MaintenanceWindowType
	Unplanned MaintenanceWindowType
}{
	"PLANNED",
	"UNPLANNED",
}

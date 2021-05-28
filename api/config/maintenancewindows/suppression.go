package maintenancewindows

// Suppression The type of suppression of alerting and problem detection during the maintenance.
type Suppression string

// Suppressions offers the known enum values
var Suppressions = struct {
	DetectProblemsAndAlert  Suppression
	DetectProblemsDontAlert Suppression
	DontDetectProblems      Suppression
}{
	"DETECT_PROBLEMS_AND_ALERT",
	"DETECT_PROBLEMS_DONT_ALERT",
	"DONT_DETECT_PROBLEMS",
}

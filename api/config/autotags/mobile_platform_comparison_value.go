package autotags

// MobilePlatformComparisonValue The value to compare to.
type MobilePlatformComparisonValue string

// MobilePlatformComparisonValues offers the known enum values
var MobilePlatformComparisonValues = struct {
	Android MobilePlatformComparisonValue
	Ios     MobilePlatformComparisonValue
	Linux   MobilePlatformComparisonValue
	MacOS   MobilePlatformComparisonValue
	Other   MobilePlatformComparisonValue
	Tvos    MobilePlatformComparisonValue
	Windows MobilePlatformComparisonValue
}{
	"ANDROID",
	"IOS",
	"LINUX",
	"MAC_OS",
	"OTHER",
	"TVOS",
	"WINDOWS",
}

package requestattributes

// CapturingAndStorageLocation Specifies the location where the values are captured and stored.
//  Required if the **source** is one of the following: `GET_PARAMETER`, `URI`, `REQUEST_HEADER`, `RESPONSE_HEADER`.
//  Not applicable in other cases.
//  If the **source** value is `REQUEST_HEADER` or `RESPONSE_HEADER`, the `CAPTURE_AND_STORE_ON_BOTH` location is not allowed.
type CapturingAndStorageLocation string

// CapturingAndStorageLocations offers the known enum values
var CapturingAndStorageLocations = struct {
	CaptureAndStoreOnBoth        CapturingAndStorageLocation
	CaptureAndStoreOnClient      CapturingAndStorageLocation
	CaptureAndStoreOnServer      CapturingAndStorageLocation
	CaptureOnClientStoreOnServer CapturingAndStorageLocation
}{
	"CAPTURE_AND_STORE_ON_BOTH",
	"CAPTURE_AND_STORE_ON_CLIENT",
	"CAPTURE_AND_STORE_ON_SERVER",
	"CAPTURE_ON_CLIENT_STORE_ON_SERVER",
}

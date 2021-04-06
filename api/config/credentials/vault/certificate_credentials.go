package vault

import "github.com/dtcookie/gojson"

// CertificateCredentials A credentials set of the `CERTIFICATE` type.
type CertificateCredentials struct {
	BaseCredentials   `json:",type=CERTIFICATE"`
	Certificate       string            `json:"certificate"`       // The certificate in the string format.
	CertificateFormat CertificateFormat `json:"certificateFormat"` // The certificate format.
	Password          string            `json:"password"`          // The password of the credential.
}

// UnmarshalJSON provides custom JSON deserialization
func (cc *CertificateCredentials) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cc)
}

// MarshalJSON provides custom JSON serialization
func (cc *CertificateCredentials) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cc)
}

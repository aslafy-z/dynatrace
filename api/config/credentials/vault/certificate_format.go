package vault

// CertificateFormat The certificate format.
type CertificateFormat string

// CertificateFormats offers the known enum values
var CertificateFormats = struct {
	Pem     CertificateFormat
	Pkcs12  CertificateFormat
	Unknown CertificateFormat
}{
	"PEM",
	"PKCS12",
	"UNKNOWN",
}

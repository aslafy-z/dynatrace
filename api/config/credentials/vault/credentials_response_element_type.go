package vault

// CredentialsResponseElementType The type of the credentials set.
type CredentialsResponseElementType string

// CredentialsResponseElementTypes offers the known enum values
var CredentialsResponseElementTypes = struct {
	Certificate      CredentialsResponseElementType
	Token            CredentialsResponseElementType
	Unknown          CredentialsResponseElementType
	UsernamePassword CredentialsResponseElementType
}{
	"CERTIFICATE",
	"TOKEN",
	"UNKNOWN",
	"USERNAME_PASSWORD",
}

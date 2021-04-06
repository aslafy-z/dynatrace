package vault

// CredentialsType Defines the actual set of fields depending on the value. See one of the following objects:
// * `CERTIFICATE` -> CertificateCredentials
// * `USERNAME_PASSWORD` -> UserPasswordCredentials
// * `TOKEN` -> TokenCredentials
type CredentialsType string

// CredentialsTypes offers the known enum values
var CredentialsTypes = struct {
	Certificate      CredentialsType
	Token            CredentialsType
	UsernamePassword CredentialsType
}{
	"CERTIFICATE",
	"TOKEN",
	"USERNAME_PASSWORD",
}

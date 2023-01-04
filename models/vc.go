package models

const (
	FormatLDP = "ldp_vc"
	FormatJWT = "jwt_vc"
)

type VerifiableCredential interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(jsonData []byte) error

	ToString() string

	GetCredentialSubject() map[string]interface{}
	GetId() string
	GetType() []string
	GetFormat() string
}

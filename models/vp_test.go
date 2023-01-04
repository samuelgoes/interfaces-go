package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVerifiablePresentation_Serialization(t *testing.T) {

	ldpString := `
	{
		"@context": ["https://www.w3.org/2018/credentials/v1", "https://www.w3.org/2018/credentials/examples/v1"],
		"credentialStatus": {
			"id": "https://backbone.gataca.io/api/v1/group/otp/status",
			"type": "CredentialStatusList2017"
		},
		"credentialSchema": {
			"id": "https://api.pilot.ebsi.eu/trusted-schemas-registry/v1/schemas/0xc5e0edc8b72616383c71469c374bb246ff8d3077cbaf7f3e68dd50cd950d7846#",
			"type": "FullJsonSchemaValidator2021"
		},
		"credentialSubject": {
			"email": "josevallejera@gmail.com",
			"id": "did:gatc:YzQxNjRjM2U4YTUzZGVkNjhmNjAxYzk5"
		},
		"id": "cred:gatc:ODUyZjE0OTM3NWZmODgzOGZjMzcwOGQw",
		"issuanceDate": "2021-12-24T12:05:25Z",
		"issuer": "did:gatc:24gsRbsURij3edoveHv81jt9EnhggrnR",
		"proof": [{
			"created": "2021-12-24T12:04:39Z",
			"creator": "did:gatc:24gsRbsURij3edoveHv81jt9EnhggrnR#keys-1",
			"domain": "gataca.io",
			"nonce": "i2Qq0nQZFZ1cXLXONnbnF7Owc4A9zFpMiSx328i1ZEQ=",
			"proofPurpose": "assertionMethod",
			"type": "JcsEd25519Signature2020",
			"verificationMethod": "did:gatc:24gsRbsURij3edoveHv81jt9EnhggrnR#keys-1"
		}],
		"type": ["VerifiableCredential", "emailCredential"]
	}`

	jwtString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NzI3NTgwMjAsImlzcyI6InNhbXVlbCJ9.edzG4Fg_PB_OnZYfW6kA6WS_fUpILerEot7Wf_oh_LE <nil>Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NzI3NTgwMjAsImlzcyI6InNhbXVlbCJ9.edzG4Fg_PB_OnZYfW6kA6WS_fUpILerEot7Wf_oh_LEToken: &{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NzI3NTgwMjAsImlzcyI6InNhbXVlbCJ9.edzG4Fg_PB_OnZYfW6kA6WS_fUpILerEot7Wf_oh_LE"

	vp := `
	{
		"type": [VerifiablePresentation],
		"verifiableCredential" : [{
		"@context": ["https://www.w3.org/2018/credentials/v1", "https://www.w3.org/2018/credentials/examples/v1"],
		"credentialStatus": {
			"id": "https://backbone.gataca.io/api/v1/group/otp/status",
			"type": "CredentialStatusList2017"
		},
		"credentialSchema": {
			"id": "https://api.pilot.ebsi.eu/trusted-schemas-registry/v1/schemas/0xc5e0edc8b72616383c71469c374bb246ff8d3077cbaf7f3e68dd50cd950d7846#",
			"type": "FullJsonSchemaValidator2021"
		},
		"credentialSubject": {
			"email": "josevallejera@gmail.com",
			"id": "did:gatc:YzQxNjRjM2U4YTUzZGVkNjhmNjAxYzk5"
		},
		"id": "cred:gatc:ODUyZjE0OTM3NWZmODgzOGZjMzcwOGQw",
		"issuanceDate": "2021-12-24T12:05:25Z",
		"issuer": "did:gatc:24gsRbsURij3edoveHv81jt9EnhggrnR",
		"proof": [{
			"created": "2021-12-24T12:04:39Z",
			"creator": "did:gatc:24gsRbsURij3edoveHv81jt9EnhggrnR#keys-1",
			"domain": "gataca.io",
			"nonce": "i2Qq0nQZFZ1cXLXONnbnF7Owc4A9zFpMiSx328i1ZEQ=",
			"proofPurpose": "assertionMethod",
			"type": "JcsEd25519Signature2020",
			"verificationMethod": "did:gatc:24gsRbsURij3edoveHv81jt9EnhggrnR#keys-1"
		}],
		"type": ["VerifiableCredential", "emailCredential"]
	},
	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NzI3NTgwMjAsImlzcyI6InNhbXVlbCJ9.edzG4Fg_PB_OnZYfW6kA6WS_fUpILerEot7Wf_oh_LE <nil>Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NzI3NTgwMjAsImlzcyI6InNhbXVlbCJ9.edzG4Fg_PB_OnZYfW6kA6WS_fUpILerEot7Wf_oh_LEToken: &{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NzI3NTgwMjAsImlzcyI6InNhbXVlbCJ9.edzG4Fg_PB_OnZYfW6kA6WS_fUpILerEot7Wf_oh_LE
	]
	}`

	var vps VerifiablePresentation
	err := json.Unmarshal([]byte(vp), &vps)
	if err != nil {
		panic(nil)
	}

	fmt.Printf("" + jwtString + ldpString)
}

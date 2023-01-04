package models

import (
	"encoding/json"
	"fmt"
)

type LDPVerifiableCredential struct {
	//Context           *SSIContext            `json:"@context,omitempty" example:"https://www.w3.org/2018/credentials/v1" description:"Context for JSON-LD"`
	//CredentialSchema  *CredentialSchema      `json:"credentialSchema,omitempty" swaggertype:"object,string" description:"Definition to retrieve the schema of the credential"`
	//CredentialStatus  *CredentialStatus      `json:"credentialStatus,omitempty" swaggertype:"object,string" description:"Definition to retrieve the current status of the credential"`
	CredentialSubject map[string]interface{} `json:"credentialSubject,omitempty" swaggerignore:"true" example:"id:did:gatc:OWM3MzQ1NzYzZilliM2Q2MTg5NGYwNTQ3,email:email@gataca.com" description:"Claims in free format stated about the subject. Linked to the credential type."` //swagger ignore because of unexpected problems with current library
	//Evidence          *Evidence              `json:"evidence,omitempty" swaggertype:"object,string" description:"Definition of the evidence. Required by eIDas"`
	//ExpirationDate    *TimeWithFormat        `json:"expirationDate,omitempty" swaggertype:"string" example:"2019-10-01T12:12:15.999Z" description:"Timestamp of expiration of the credential"`
	Id string `json:"id" example:"cred:gatc:2798d86b-21ce-4794-a06a-275d12b43e48" description:"Unique identifier of the Verifiable Credential"`
	//IssuanceDate      *TimeWithFormat        `json:"issuanceDate,omitempty"  swaggertype:"string" example:"2019-10-01T12:12:05.999Z" description:"Timestamp of issuance of the credential"`
	//Issued            *TimeWithFormat        `json:"issued,omitempty"  swaggertype:"string" example:"2019-10-01T12:12:05.999Z" description:"Timestamp of issuance of the credential"`
	//Issuer            string                 `json:"issuer,omitempty" example:"did:gatc:OWM3MzQ1NzYzZilliM2Q2MTg5NGYwNTQ3" description:"Issuer of the credential"`
	//Proof             *SSIProof              `json:"proof,omitempty"  description:"Proofs to verify the presentation"`
	Type []string `json:"type,omitempty" example:"emailCredential" description:"Type definition of this verifiable credential establishing a specific json schema."`
	//ValidFrom         *TimeWithFormat        `json:"validFrom,omitempty" swaggertype:"string" example:"2019-10-01T12:12:05.999Z" description:"Timestamp from which the credential its valid"`
}

func NewLDPVerifiableCredential() LDPVerifiableCredential {
	return LDPVerifiableCredential{}
}

func (ldp *LDPVerifiableCredential) MarshalJSON() ([]byte, error) {
	type Alias LDPVerifiableCredential
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(ldp),
	})
}
func (ldp *LDPVerifiableCredential) UnmarshalJSON(data []byte) error {

	println("Data: ", string(data))

	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	aInterface := v["type"].([]interface{})
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}

	ldp.CredentialSubject, _ = v["credentialSubject"].(map[string]interface{})
	ldp.Id, _ = v["id"].(string)
	ldp.Type = aString

	return nil
}

func (ldp *LDPVerifiableCredential) ToString() string {
	return fmt.Sprintf("Verifiable Credential: Type --> %s", ldp.GetFormat())
}

func (ldp *LDPVerifiableCredential) GetCredentialSubject() map[string]interface{} {
	return nil
}
func (ldp *LDPVerifiableCredential) GetId() string {
	return ldp.Id
}
func (ldp *LDPVerifiableCredential) GetType() []string {
	return ldp.Type
}
func (ldp *LDPVerifiableCredential) GetFormat() string {
	return FormatLDP
}

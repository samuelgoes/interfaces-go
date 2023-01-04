package models

import (
	"fmt"
)

type JWTVerifiableCredential struct {
	//ExpirationDate       *NumericDate             `json:"exp,omitempty" swaggertype:"int" example:"1672313881" description:"Timestamp of expiration of the credential"`
	Id string `json:"jti" swaggertype:"string" example:"vc:ebsi:authentication#2798d86b-21ce-4794-a06a-275d12b43e48" description:"Unique identifier of the Verifiable Credential"`
	//Issued               *NumericDate             `json:"iat,omitempty" swaggertype:"int" example:"1672313881" description:"Timestamp of expiration of the credential"`
	Issuer string `json:"iss,omitempty" swaggertype:"string" example:"did:ebsi:zr2rWDHHrUCdZAW7wsSb5nQ" description:"Issuer of the credential"`
	//Proof                *SSIProof                `json:"-"`
	Subject string `json:"sub,omitempty" swaggertype:"string" example:"did:ebsi:zx2lowdbKPAugWPS2UT9TT" description:"DID of the VC owner"`
	//ValidFrom            *NumericDate             `json:"nbf,omitempty" swaggertype:"int" example:"1672313881" description:"Timestamp from which the credential its valid"`
	VerifiableCredential *LDPVerifiableCredential `json:"vc,omitempty" swaggertype:"object,string" description:"VC in W3C format"`
}

func NewJWTVerifiableCredential() JWTVerifiableCredential {
	return JWTVerifiableCredential{}
}

func (jvc *JWTVerifiableCredential) MarshalJSON() ([]byte, error) {
	return nil, nil
}
func (jvc *JWTVerifiableCredential) UnmarshalJSON(jsonData []byte) error {
	return nil
}

func (jvc JWTVerifiableCredential) ToString() string {
	return fmt.Sprintf("Verifiable Credential: Type --> %s", jvc.GetFormat())
}

func (jvc JWTVerifiableCredential) GetCredentialSubject() map[string]interface{} {
	return jvc.VerifiableCredential.GetCredentialSubject()
}
func (jvc JWTVerifiableCredential) GetId() string {
	return jvc.VerifiableCredential.GetId()
}
func (jvc JWTVerifiableCredential) GetType() []string {
	return jvc.VerifiableCredential.GetType()
}
func (jvc JWTVerifiableCredential) GetFormat() string {
	return FormatJWT
}

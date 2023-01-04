package models

type VerifiablePresentation struct {
	//Context                *SSIContext             `json:"@context,omitempty" example:"https://www.w3.org/2018/credentials/v1" description:"Context for JSON-LD"`
	//DataAgreementId string `json:"data_agreement_id,omitempty" example:"da:gatc:e3iu4wg39487wq9gf7a47af37" description:"Id of the data agreement supporting this exchange"`
	//Holder          string `json:"holder,omitempty" example:"did:gatc:example1234567" description:"DID of the Holder of the credentials"`
	//PresentationSubmission *PresentationSubmission `json:"presentation_submission,omitempty" description:"Presentation submission according to DIF PE"`
	//Proof                  *SSIProof               `json:"proof,omitempty" description:"Proofs to verify the presentation"`
	Type                 []string      `json:"type,omitempty" example:"VerifiablePresentation" description:"Definition of the format of the presentation"`
	VerifiableCredential []interface{} `json:"verifiableCredential,omitempty" description:"List of Verifiable Credentials included in the presentation"`
}

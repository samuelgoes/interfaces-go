package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/samuelgoes/interfaces-go/models"
	"reflect"
	"time"
)

const secret = "AllYourBase"

func main() {
	token, _ := generateJWT()

	fmt.Printf("Token: %s", token)

	verifyJWT(token)

	serialization(token)
}

func generateJWT() (string, error) {

	type MyCustomClaims struct {
		Credentials []models.VerifiableCredential `json:"vc"`
		jwt.StandardClaims
	}

	vc := []models.VerifiableCredential{
		&models.LDPVerifiableCredential{
			CredentialSubject: map[string]interface{}{
				"id":    "did:example:ebfeb1f712ebc6f1c276e12ec21",
				"email": "samuelgoes@gmail.com",
			},
			Id:   "cred:gatc:jfkdlsajfdlkajfdlas",
			Type: []string{"VerifiableCredential", "EmailCredential"},
		},
	}

	// Create the Claims
	claims := MyCustomClaims{
		vc,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			NotBefore: time.Now().Add(time.Minute * 10).Unix(),
			Id:        "cred:gatc:jfkdlsajfdlkajfdlas",
			Issuer:    "did:gatc:jfkdlajfksafjdzlakd",
			Subject:   "did:example:ebfeb1f712ebc6f1c276e12ec21",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))

	fmt.Printf("%v %v", ss, err)

	return ss, nil
}

func verifyJWT(tokenS string) {
	token, err := jwt.Parse(tokenS, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		fmt.Printf("Error parsing token")
	}

	fmt.Printf("Token: %v", token)
}

func serialization(token string) {
	vp := `
	{
		"type": ["VerifiablePresentation"],
		"verifiableCredential" : [
			{
				"credentialSubject": {
					"email": "samuelgoes@gmail.com",
					"id": "did:gatc:YzQxNjRjM2U4YTUzZGVkNjhmNjAxYzk5"
				},
				"id": "cred:gatc:ODUyZjE0OTM3NWZmODgzOGZjMzcwOGQw",
				"type": ["VerifiableCredential", "emailCredential"]
			},
			"` + token +
		`"]
	}`

	var vps models.VerifiablePresentation
	err := json.Unmarshal([]byte(vp), &vps)
	if err != nil {
		panic(nil)
	}

	fmt.Print("\n############################################################\n\n")

	decodeVCs(&vps)

	fmt.Println("VP: ", vps)

	vpB, err := json.Marshal(vps)
	if err != nil {
		panic(-1)
	}

	fmt.Printf("\nResult: %s", string(vpB))

}

func decodeVCs(vp *models.VerifiablePresentation) []models.VerifiableCredential {
	var vcl []models.VerifiableCredential
	var vc models.VerifiableCredential
	for _, vci := range vp.VerifiableCredential {
		switch v := vci.(type) {
		case string:
			fmt.Println("type JWT: ", vci)
			vc = decodeString(v)
			break
		case map[string]interface{}:
			fmt.Println("type map: ", vci)
			vc = decodeMap(v)
			break
		default:
			fmt.Println("type: ", reflect.TypeOf(vci), "\n")
		}

		vcl = append(vcl, vc)
	}

	return vcl
}

func decodeString(s string) models.VerifiableCredential {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Printf("Error parsing token")
	}

	fmt.Printf("Token: %v\n", token)

	var vc models.VerifiableCredential
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c := claims["vc"].([]interface{})
		// Convert map to json string
		jsonStr, err := json.Marshal(c[0])
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("jsonString: ", string(jsonStr))

		// Convert json string to struct
		var ldp models.LDPVerifiableCredential
		if err := json.Unmarshal(jsonStr, &ldp); err != nil {
			fmt.Println(err)
		}
		vc = &models.JWTVerifiableCredential{
			Id:                   fmt.Sprint(claims["jti"]),
			Issuer:               fmt.Sprint(claims["iss"]),
			Subject:              fmt.Sprint(claims["sub"]),
			VerifiableCredential: &ldp,
		}

	}

	return vc
}

func decodeMap(m map[string]interface{}) models.VerifiableCredential {
	vc := models.LDPVerifiableCredential{}
	if m["id"] != nil {
		buf, err := json.Marshal(&m)
		fmt.Println("Bytes: ", string(buf), "\n")
		err = json.Unmarshal(buf, &vc)
		if err != nil {
			panic(nil)
		}
	}

	return &vc
}

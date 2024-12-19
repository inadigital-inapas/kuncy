package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

func main() {
	fmt.Println("+++++++++++++++++++++++++++++++++")
	fmt.Println("Kuncy - INApas key pair generator")
	fmt.Println("+++++++++++++++++++++++++++++++++")
	fmt.Print("\n")

	var err error

	newKidSign, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	kidSign := fmt.Sprintf("%X", sha256.Sum256([]byte(newKidSign.String())))

	skSignPem, pkSignPem := genSigningV1()
	testSignVerify(skSignPem, pkSignPem, kidSign)
	writeSignToFile(skSignPem, pkSignPem)

	// test1()
	fmt.Println("")

	newKidEnc, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	kidEnc := fmt.Sprintf("%X", sha256.Sum256([]byte(newKidEnc.String())))

	skEncPem, pkEncPem := genEncryptionV1()
	testEncVerify(skEncPem, pkEncPem, kidEnc)
	writeEncToFile(skEncPem, pkEncPem)

	// print JWKS
	jwkEnc := loadECDSAPemToJWK(skEncPem, kidEnc)
	jwkSign := loadEd25519PemToJWK(pkSignPem, kidSign)

	jwks := jwk.NewSet()
	jwks.AddKey(jwkEnc)
	jwks.AddKey(jwkSign)
	jwksInJSON, err := json.MarshalIndent(jwks, "", "  ")
	if err != nil {
		panic(err)
	}

	os.Remove("jwks.json")
	err = os.WriteFile("jwks.json", jwksInJSON, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Print("\n")
	fmt.Println("+++++++++++++ EXIT +++++++++++++")
}

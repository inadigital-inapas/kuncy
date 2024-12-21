package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
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

	skSignPem, pkSignPem := genECDSA()
	testSignECDSA(skSignPem, pkSignPem, kidSign)
	writeSignECDSAToFile(skSignPem, pkSignPem)

	// test1()
	fmt.Println("")

	newKidEnc, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	kidEnc := fmt.Sprintf("%X", sha256.Sum256([]byte(newKidEnc.String())))

	skEncPem, pkEncPem := genECDSA()
	testEncECDSA(skEncPem, pkEncPem, kidEnc)
	writeEncECDSAToFile(skEncPem, pkEncPem)

	// print JWKS
	jwkEnc := loadECDSAPemToJWK(pkEncPem)
	jwkEnc.Set(jwk.KeyIDKey, kidEnc)
	jwkEnc.Set(jwk.KeyUsageKey, jwk.ForEncryption)
	jwkEnc.Set(jwk.AlgorithmKey, jwa.ECDH_ES_A256KW.String())

	jwkSign := loadECDSAPemToJWK(pkSignPem)
	jwkSign.Set(jwk.KeyIDKey, kidSign)
	jwkSign.Set(jwk.KeyUsageKey, jwk.ForSignature)
	jwkSign.Set(jwk.AlgorithmKey, jwa.ES512.String())

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

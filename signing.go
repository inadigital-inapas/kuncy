package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func genSigningV1() (privateKey []byte, publicKey []byte) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	return toPem(PEMTYPE_PRIVATE_KEY, privKey), toPem(PEMTYPE_PUBLIC_KEY, pubKey)
}

func testSignVerify(skSignPem []byte, pkSignPem []byte, kid string) {
	signingOK := "FAILED"
	verifyOK := "FAILED"

	// Create a new JWT
	token, err := jwt.NewBuilder().
		IssuedAt(time.Now()).
		Subject("INAPAS99ID").
		Expiration(time.Now().Add(15*time.Minute)).
		JwtID(srg()).
		Claim("type", "assertion").
		Build()
	if err != nil {
		panic(err)
	}

	// Test Signing
	skSignJWK := loadEd25519PemToJWK(skSignPem, kid)
	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.EdDSA, skSignJWK))
	if err != nil {
		panic(err)
	}

	signingOK = "PASSED"
	fmt.Println("+ test signing in JWT format:", signingOK)
	if signingOK == "PASSED" {
		fmt.Println("+ + output:", string(signedToken))
	}
	fmt.Println("")

	// Test Verify
	pkSignJWK := loadEd25519PemToJWK(pkSignPem, kid)
	newKeySet := jwk.NewSet()
	err = newKeySet.AddKey(pkSignJWK)
	if err != nil {
		panic(err)
	}

	parsedJWT, err := jwt.Parse(signedToken, jwt.WithKeySet(newKeySet))
	if err != nil {
		panic(err)
	}

	verifyOK = "PASSED"
	fmt.Println("+ test parsing signed JWT:", verifyOK)
	if verifyOK == "PASSED" {
		if parsedJWT != nil {
			parsed2InMap, err := parsedJWT.AsMap(context.Background())
			fmt.Println("+ + output:", parsed2InMap, "err:", err)
		}
	}
}

func writeSignToFile(skSignPem, pkSignPem []byte) {
	// Save Private Key
	os.Remove("signing_privkey.pem")
	err := os.WriteFile("signing_privkey.pem", skSignPem, 0644)
	if err != nil {
		panic(err)
	}

	// Save Public Key
	os.Remove("signing_pubkey.pem")
	err = os.WriteFile("signing_pubkey.pem", pkSignPem, 0644)
	if err != nil {
		panic(err)
	}
}

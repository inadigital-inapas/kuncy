package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func genECDSA() (privateKey []byte, publicKey []byte) {
	sk, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}

	// Marshal the private key into PKCS#8 DER format
	privBytes, err := x509.MarshalPKCS8PrivateKey(sk)
	if err != nil {
		panic(err)
	}

	// Marshal the public key into PKIX DER format
	pubBytes, err := x509.MarshalPKIXPublicKey(&sk.PublicKey)
	if err != nil {
		panic(err)
	}

	return toPem(PEMTYPE_PRIVATE_KEY, privBytes), toPem(PEMTYPE_PUBLIC_KEY, pubBytes)
}

// Encrypt and Decrypt

func testEncECDSA(skEncPem []byte, pkEncPem []byte, kid string) {
	// plainText := "Lorem Impsum Dolor Sit Amet"
	encryptionOK := "FAILED"
	decryptionOK := "FAILED"

	type selfData struct {
		Name     string `json:"string"`
		NIK      string `json:"nik"`
		INApasID string `json:"inapasID"`
	}

	payload, err := json.Marshal(selfData{Name: "Fulan", NIK: "1000200030004000", INApasID: "INAPAS99ID"})
	if err != nil {
		panic(err)
	}

	protected := jwe.NewHeaders()
	protected.Set(`kid`, kid)

	jwkPubKey := loadECDSAPemToJWK(pkEncPem)
	jwkPubKey.Set(jwk.KeyUsageKey, jwk.ForEncryption)
	jwkPubKey.Set(jwk.AlgorithmKey, jwa.ECDH_ES_A256KW.String())
	jwkPubKey.Set(jwk.KeyIDKey, kid)

	encrypted, err := jwe.Encrypt(
		payload,
		jwe.WithKey(jwa.ECDH_ES_A256KW, jwkPubKey),
		jwe.WithProtectedHeaders(protected),
	)
	if err != nil {
		panic(err)
	}

	encryptionOK = "PASSED"
	fmt.Println("+ test encryption in JWT format:", encryptionOK)
	if encryptionOK == "PASSED" {
		fmt.Println("+ + output:", string(encrypted))
	}
	fmt.Println("")

	// Test Verify
	skEncJWK := loadECDSAPemToJWK(skEncPem)
	skEncJWK.Set(jwk.AlgorithmKey, jwa.ECDH_ES_A256KW.String())
	skEncJWK.Set(jwk.KeyIDKey, kid)

	newKeySet := jwk.NewSet()
	err = newKeySet.AddKey(skEncJWK)
	if err != nil {
		panic(err)
	}

	decrypted, err := jwe.Decrypt(encrypted, jwe.WithKeySet(newKeySet))
	if err != nil {
		panic(err)
	}

	decryptionOK = "PASSED"
	fmt.Println("+ test decrypt chipertex of JWE:", decryptionOK)
	if decryptionOK == "PASSED" {
		fmt.Println("+ + output:", string(decrypted), "err:", err)
	}
}

func writeEncECDSAToFile(skEncPem, pkEncPem []byte) {
	// Save Private Key
	os.Remove("enc_privkey.pem")
	err := os.WriteFile("enc_privkey.pem", skEncPem, 0644)
	if err != nil {
		panic(err)
	}

	// Save Public Key
	os.Remove("enc_pubkey.pem")
	err = os.WriteFile("enc_pubkey.pem", pkEncPem, 0644)
	if err != nil {
		panic(err)
	}
}

// Sign and Verify

func testSignECDSA(skSignPem []byte, pkSignPem []byte, kid string) {
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
	skSignJWK := loadECDSAPemToJWK(skSignPem)
	skSignJWK.Set(jwk.KeyUsageKey, jwk.ForSignature)
	skSignJWK.Set(jwk.AlgorithmKey, jwa.ES512.String())
	skSignJWK.Set(jwk.KeyIDKey, kid)

	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.ES512, skSignJWK))
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
	pkSignJWK := loadECDSAPemToJWK(pkSignPem)
	pkSignJWK.Set(jwk.KeyUsageKey, jwk.ForSignature)
	pkSignJWK.Set(jwk.AlgorithmKey, jwa.ES512.String())
	pkSignJWK.Set(jwk.KeyIDKey, kid)

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

func writeSignECDSAToFile(skSignPem, pkSignPem []byte) {
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

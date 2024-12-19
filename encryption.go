package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

func genEncryptionV1() (privateKey []byte, publicKey []byte) {
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

func testEncVerify(skEncPem []byte, pkEncPem []byte, kid string) {
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

	encrypted, err := jwe.Encrypt(
		payload,
		jwe.WithKey(jwa.ECDH_ES_A256KW, loadECDSAPemToJWK(pkEncPem, kid)),
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
	skEncJWK := loadECDSAPemToJWK(skEncPem, kid)
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

func writeEncToFile(skEncPem, pkEncPem []byte) {
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

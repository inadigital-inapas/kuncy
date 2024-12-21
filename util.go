package main

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

const PEMTYPE_PRIVATE_KEY = "PRIVATE KEY"
const PEMTYPE_PUBLIC_KEY = "PUBLIC KEY"

func srg() string {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%X", sha256.Sum256(key))
}

func toPem(c string, b []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  c,
		Bytes: b,
	})
}

func loadEd25519PemToJWK(pemBlock []byte, kid string) jwk.Key {
	block, _ := pem.Decode(pemBlock)
	if block == nil {
		panic("failed to decode PEM block")
	}

	switch block.Type {
	case PEMTYPE_PRIVATE_KEY:
		key := ed25519.PrivateKey(block.Bytes)
		jwkKey, err := jwk.FromRaw(key)
		if err != nil {
			panic(err)
		}

		jwkKey.Set(jwk.AlgorithmKey, jwa.EdDSA.String())
		jwkKey.Set(jwk.KeyUsageKey, jwk.ForSignature)
		if kid != "" {
			jwkKey.Set(jwk.KeyIDKey, kid)
		}

		return jwkKey
	case PEMTYPE_PUBLIC_KEY:
		key := ed25519.PublicKey(block.Bytes)
		jwkKey, err := jwk.FromRaw(key)
		if err != nil {
			panic(err)
		}

		jwkKey.Set(jwk.AlgorithmKey, jwa.EdDSA.String())
		jwkKey.Set(jwk.KeyUsageKey, jwk.ForSignature)
		if kid != "" {
			jwkKey.Set(jwk.KeyIDKey, kid)
		}

		return jwkKey
	default:
		panic("failed to load PEM to JWK: unsupported block type")
	}
}

func loadECDSAPemToJWK(pemBlock []byte) jwk.Key {
	block, _ := pem.Decode(pemBlock)
	if block == nil {
		panic("failed to decode PEM block")
	}

	switch block.Type {
	case PEMTYPE_PRIVATE_KEY:
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			panic(err)
		}

		ecdsaPriv, ok := key.(*ecdsa.PrivateKey)
		if !ok {
			panic("failed to load PEM to JWK: unsupported public key type")
		}

		jwkKey, err := jwk.FromRaw(ecdsaPriv)
		if err != nil {
			panic(err)
		}

		return jwkKey
	case PEMTYPE_PUBLIC_KEY:
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			panic(err)
		}

		ecdsaPub, ok := key.(*ecdsa.PublicKey)
		if !ok {
			panic("failed to load PEM to JWK: unsupported public key type")
		}

		jwkKey, err := jwk.FromRaw(ecdsaPub)
		if err != nil {
			panic(err)
		}

		return jwkKey
	default:
		panic("failed to load PEM to JWK: unsupported block type")
	}
}

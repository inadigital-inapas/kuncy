# Kuncy - INApas Key Pair Generator

Kuncy is a Go-based utility designed for generating, signing, and encrypting cryptographic key pairs and managing JSON Web Key Sets (JWKS). This project aligns with the INApas digital identity framework, offering support for ECDSA and EdDSA keys and utilizing modern cryptographic libraries.

---

## Features

- **Key Generation**: Generate secure ECDSA and EdDSA key pairs for signing and encryption.
- **Signing and Verification**: Create and verify signed JWTs with ES512 and EdDSA algorithms.
- **Encryption and Decryption**: Encrypt payloads using ECDH-ES with AES-256-KW and decrypt them using JWK key sets.
- **JWKS Management**: Automatically generate and export JSON Web Key Sets (JWKS) for use in secure applications.
- **Platform Compatibility**: Build and package the application for Linux, macOS, and Windows.

---

## Installation

### Download the Binary

You can download the prebuilt binary from the [GitHub Releases Page](https://github.com/inadigital-inapas/kuncy/releases).

#### Option 1: Using `curl`

For Linux:

```
curl -L -o kuncy.tar.gz https://github.com/inadigital-inapas/kuncy/releases/download/<version>/kuncy-<version>-linux.tar.gz
```

For macOS:

```
curl -L -o kuncy.tar.gz https://github.com/inadigital-inapas/kuncy/releases/download/<version>/kuncy-<version>-darwin.tar.gz
```

For Windows:

```
curl -L -o kuncy.zip https://github.com/inadigital-inapas/kuncy/releases/download/<version>/kuncy-<version>-windows.zip
```

#### Option 2: Using `wget`

For Linux:

```
wget -O kuncy.tar.gz https://github.com/inadigital-inapas/kuncy/releases/download/<version>/kuncy-<version>-linux.tar.gz
```

For macOS:

```
wget -O kuncy.tar.gz https://github.com/inadigital-inapas/kuncy/releases/download/<version>/kuncy-<version>-darwin.tar.gz
```

For Windows:

```
wget -O kuncy.zip https://github.com/inadigital-inapas/kuncy/releases/download/<version>/kuncy-<version>-windows.zip
```

#### Option 3: Manual Download

1. Visit the [GitHub Releases Page](https://github.com/inadigital-inapas/kuncy/releases).
2. Download the appropriate binary for your operating system.
3. Save it to your desired directory.

---

### Extract the Binary

For Linux and macOS (tar.gz files):

```
tar -xvzf kuncy.tar.gz
chmod +x kuncy
```

For Windows (zip files):

Use your preferred archive manager (e.g., WinRAR, 7-Zip) to extract the `kuncy.exe` file.

---

## Usage

Run the application to generate key pairs and manage JWKS:

```
./kuncy
```

By default, the application performs the following:

- Generates ECDSA and EdDSA key pairs.
- Tests signing and encryption functionalities.
- Exports the generated JWKS to `jwks.json`.

You can customize the behavior by modifying the source files if you build from source.

### example output
```
+++++++++++++++++++++++++++++++++
Kuncy - INApas key pair generator
+++++++++++++++++++++++++++++++++

+ test signing in JWT format: PASSED
+ + output: eyJhbGciOiJFUzUxMiIsImtpZCI6IkFGRTlFRDIxQkVERUQ0NENGMUYxNTg5RjlEOTY3RUM4RTY4MTczNUU3NzQyNEUyRDkyMTg3QzU0OThDOTZCQzIiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE3MzYyMTkzMzQsImlhdCI6MTczNjIxODQzNCwianRpIjoiNDgzQTU5MjVCMEVBQjNEMUJDRUM5NTU1MzM2NTM1M0FBOUEwNjlBNzAzNUZFQjczOEJFMjMzMTEzNTc5QjBEMSIsInN1YiI6IklOQVBBUzk5SUQiLCJ0eXBlIjoiYXNzZXJ0aW9uIn0.ALbQ982JQPYs2ylfs8CJb9R_W9tw-e4PJWngfKGZEjGA6hc0-P2u_MTBTCaeFxTZOSvrKR6S0Kh7fTbynU49otGYAaW71foc3pgbHKVJSaEiUGGNCbXwoNWN2BJfp2IXQ564-nptK0rvJtwgL5oG12AIsPNYesB5DXF1YbaQgrxGIip6

+ test parsing signed JWT: PASSED
+ + output: map[exp:2025-01-07 03:08:54 +0000 UTC iat:2025-01-07 02:53:54 +0000 UTC jti:483A5925B0EAB3D1BCEC95553365353AA9A069A7035FEB738BE233113579B0D1 sub:INAPAS99ID type:assertion] err: <nil>

+ test encryption in JWT format: PASSED
+ + output: eyJhbGciOiJFQ0RILUVTK0EyNTZLVyIsImVuYyI6IkEyNTZHQ00iLCJlcGsiOnsiY3J2IjoiUC01MjEiLCJrdHkiOiJFQyIsIngiOiJBY3U1YXcxcjBIa2xQaEl2QTlxT1ZtQ0c5VHRKd1FGT25KR0tfSTdxWllVZFdoU0hkTTdNMDhkc19ZSGV6WVFzMGZ3SGhnZHFXWkxYSWQ4bmZrUklNSElZIiwieSI6IkFlVktOS2ZKaDhoWnVMU0M1dEd3eHJJSGx2TnJsaVFZSmtyWlk2NmN0cWhLX196cVVqYV9WYVZHbUctZkllTF9ZN0Qza1lYdDJrampMR3FMVDNrQTRGRlcifSwia2lkIjoiNDQzNTRFRTk5NzBDN0QzNTNBMEQ5MjlFNjAwMTlGQTM3MUFDRjc3OUZBNUQ4OTBERTJGMUU4QzAwQkE3NDgzMiJ9.JS43pnFL4UGSw89esMyBRE3cwLK_etz05I7RGbzNOey23M_FiX0Kdg.nSvSyJ1mGFpMPzeh.GDjKOAsLbXbF7qnSeaLXBLqAPwn8Q6nTgu4-DsVJxeybVyqAwmOweuITqHa82dNINKHXJX7vgoo02VwRJr04kmPEpA.Zu6Luxhrig7hwlLkp1LKfA

+ test decrypt chipertex of JWE: PASSED
+ + output: {"string":"Fulan","nik":"1000200030004000","inapasID":"INAPAS99ID"} err: <nil>

+++++++++++++ EXIT +++++++++++++
```

### file generated
```
- enc_privkey.pem
- enc_pubkey.pem
- jwks.json
- signing_privkey.pem
- signing_pubkey.pem
```

The file `jwks.json` contains both of signing and encryption public key in JWKS format. And it should be registered in INApas.

You have to keep the `enc_privkey.pem` and `signing_privkey.pem` securely. As the encryption privkey is used to decrypt the id_token payload. and signing privkey is to sign a jwt assertion for oidc exchange token.   

---

## Dependencies

- [lestrrat-go/jwx](https://github.com/lestrrat-go/jwx): For working with JWTs, JWKs, and cryptographic operations.
- [google/uuid](https://github.com/google/uuid): UUID generation.

For a full list of dependencies, see [`go.mod`](go.mod).

---

## Contribution

Feel free to fork this repository and open a pull request to contribute. Ensure that all new features are well-documented and tested.

---

## License

This project is licensed under the Apache License 2.0. See the [`LICENSE`](LICENSE) file for more details.

---

## Contact

For inquiries or support, please contact **@ghiyastfarisi** or open an issue in this repository.

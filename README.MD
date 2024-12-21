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

You can download the prebuilt binary from the [GitHub Releases Page](https://github.com/inadigita-inapas/kuncy/releases).

#### Option 1: Using `curl`

For Linux:

```
curl -L -o kuncy.tar.gz https://github.com/inadigita-inapas/kuncy/releases/download/<version>/kuncy-<version>-linux.tar.gz
```

For macOS:

```
curl -L -o kuncy.tar.gz https://github.com/inadigita-inapas/kuncy/releases/download/<version>/kuncy-<version>-darwin.tar.gz
```

For Windows:

```
curl -L -o kuncy.zip https://github.com/inadigita-inapas/kuncy/releases/download/<version>/kuncy-<version>-windows.zip
```

#### Option 2: Using `wget`

For Linux:

```
wget -O kuncy.tar.gz https://github.com/inadigita-inapas/kuncy/releases/download/<version>/kuncy-<version>-linux.tar.gz
```

For macOS:

```
wget -O kuncy.tar.gz https://github.com/inadigita-inapas/kuncy/releases/download/<version>/kuncy-<version>-darwin.tar.gz
```

For Windows:

```
wget -O kuncy.zip https://github.com/inadigita-inapas/kuncy/releases/download/<version>/kuncy-<version>-windows.zip
```

#### Option 3: Manual Download

1. Visit the [GitHub Releases Page](https://github.com/inadigita-inapas/kuncy/releases).
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

### Move the Binary (Optional)

For Linux and macOS:

```
sudo mv kuncy /usr/local/bin/
```

This allows you to run `kuncy` from anywhere on your system.

---

## Usage

Run the application to generate key pairs and manage JWKS:

```
kuncy
```

By default, the application performs the following:

- Generates ECDSA and EdDSA key pairs.
- Tests signing and encryption functionalities.
- Exports the generated JWKS to `jwks.json`.

You can customize the behavior by modifying the source files if you build from source.

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
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

1. **Download the prebuilt binary:**

   Visit the [Releases Page](https://github.com/inadigita-inapas/kuncy/releases) to download the latest version of Kuncy for your operating system.

   Example for Linux:

   ```
   wget https://github.com/inadigita-inapas/kuncy/releases/download/<version>/kuncy-<version>-linux-amd64.tar.gz
   tar -xvzf kuncy-<version>-linux-amd64.tar.gz
   ```

   Replace `<version>` with the version number of the release you want to download.

2. **Make the binary executable:**

   ```
   chmod +x kuncy
   ```

3. **Move the binary to your PATH (optional):**

   ```
   sudo mv kuncy /usr/local/bin/
   ```

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

For inquiries or support, please contact **Rizki Auliya Rachmawati** or open an issue in this repository.

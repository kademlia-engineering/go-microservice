/*
eccCrypto.go
v0.1.0
1/2/24

This file defines an interface for eliptic curve operations
*/
package crypto

// Ed25519Crypto implements ECCCrypto using the Ed25519 curve.
type Ed25519Crypto struct{}

// ECCCrypto is an interface for elliptic curve cryptography operations.
type ECCCrypto interface {
	GenerateKeyPair() (publicKey, privateKey []byte, err error)
	Sign(privateKey, publicKey, message []byte) (signature []byte, err error)
	Verify(publicKey, message, signature []byte) bool
}

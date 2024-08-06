/*
ed25519.go
v0.1.0
1/2/24

This file implements the eccCrypto interface for the eliptic curve ed25519
*/
package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

// GenerateKeyPair generates an Ed25519 public/private key pair.
func (e *Ed25519Crypto) GenerateKeyPair() (publicKey, privateKey []byte, err error) {
	public, private, err := ed25519.GenerateKey(rand.Reader)
	return public, private, err
}

// Sign signs the message with the private key.
func (e *Ed25519Crypto) Sign(privateKey, message []byte) (signature []byte, err error) {
	private := ed25519.PrivateKey(privateKey)
	return ed25519.Sign(private, message), nil
}

// Verify verifies the signature of the message against the public key.
func (e *Ed25519Crypto) Verify(publicKey, message, signature []byte) bool {
	public := ed25519.PublicKey(publicKey)
	return ed25519.Verify(public, message, signature)
}

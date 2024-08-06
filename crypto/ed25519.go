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
	if err != nil {
		return nil, nil, err
	}
	// The actual private key is the first 32 bytes of the generated private key.
	scalar := private[:32]
	return public, scalar, nil
}

// Sign signs the message with the private key.
func (e *Ed25519Crypto) Sign(privateKey, publicKey, message []byte) (signature []byte, err error) {
	pk := append(privateKey, publicKey...)
	private := ed25519.PrivateKey(pk)
	return ed25519.Sign(private, message), nil
}

// Verify verifies the signature of the message against the public key.
func (e *Ed25519Crypto) Verify(publicKey, message, signature []byte) bool {
	public := ed25519.PublicKey(publicKey)
	return ed25519.Verify(public, message, signature)
}

/*
aesCrypto.go
v0.1.0
1/2/24

This file exposes an aes encrypt and decrypt methods
*/
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/scrypt"
)

// EncryptData encrypts the given data using scrypt and AES-GCM.
func EncryptData(data []byte, password string) ([]byte, error) {
	// Derive key using scrypt
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	// Create AES-GCM cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}

	// Creating the nonce (nonce length should be 16 bytes)
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the data
	ciphertextSlice := aesgcm.Seal(nil, nonce, data, nil)
	ciphertext := hex.EncodeToString(ciphertextSlice)
	tag, err := splitEncodedString(ciphertext)
	if err != nil {
		return nil, err
	}

	// Convert to hexadecimal and format as <salt>.<iv>.<tag>.<ciphertext>
	result := fmt.Sprintf("%s.%s.%s.%s",
		hex.EncodeToString(salt),
		hex.EncodeToString(nonce),
		tag[1],
		tag[0],
	)
	return []byte(result), nil
}

func DecryptData(serializedData string, password string) ([]byte, error) {
	// Extracting salt, iv, tag, and ciphertext from the serializedData
	parts := splitSerializedData(serializedData)
	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	iv, err := hex.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	tag, err := hex.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}
	ciphertext, err := hex.DecodeString(parts[3])
	if err != nil {
		return nil, err
	}
	ciphertext = []byte(fmt.Sprintf("%s%s", ciphertext, tag))

	// Derive key using scrypt
	key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	// Create AES-GCM cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}

	// Decrypting the data
	dataBuffer, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return dataBuffer, nil
}

// splitSerializedData splits the serialized data string into parts
func splitSerializedData(serializedData string) []string {
	return strings.Split(serializedData, ".")
}

func splitEncodedString(encodedString string) ([]string, error) {
	// Decode the hex string
	decoded, err := hex.DecodeString(encodedString)
	if err != nil {
		return nil, err
	}

	// Check if the decoded string is at least 16 bytes long
	if len(decoded) < 16 {
		return nil, fmt.Errorf("decoded string is too short")
	}

	// Extract the last 16 bytes
	last16Bytes := decoded[len(decoded)-16:]

	// Create two strings: one with the original length - 16, and one with the last 16 bytes
	firstPart := encodedString[:len(encodedString)-32] // original length - 32
	secondPart := hex.EncodeToString(last16Bytes)

	return []string{firstPart, secondPart}, nil
}

/*
elipticCurveCryptography.go
v0.1.0
1/2/24

This file defines the handler functions for AES based operations
*/
package handlers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/btcsuite/btcutil/base58"
	"github.com/sirupsen/logrus"
	"kademlia.io/crypto"
	"kademlia.io/models"
	"kademlia.io/server"
)

const ED_CURVE string = "ed23319"

// Ed25519 Keypair. This api returns an ed25519 keypair.
//
//	@param  s - server.Server
//	@return http.HandlerFunc
func Ed25519Keypair(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Ed25519Keypair")

		ecc := &crypto.Ed25519Crypto{}
		publicKey, privateKey, err := ecc.GenerateKeyPair()
		if err != nil {
			logrus.Errorf("error: %s", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.HttpResponse{
				Payload: fmt.Sprintf("%s", err),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Keypair{
			Curve:      ED_CURVE,
			PublicKey:  base58.Encode(publicKey),
			PrivateKey: base58.Encode(privateKey),
		})
	}
}

// Ed25519 Signature. Given a private key and a message,
// this method will sign the message and return the signature.
//
//	@param  s - server.Server
//	@return http.HandlerFunc
func Ed25519Signature(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Ed25519Signature")

		// Extract Keypair
		publicKeyBase58 := r.Header.Get("public-key")
		privateKeyBase58 := r.Header.Get("private-key")

		// Decode to a slice of bytes
		publicKey := base58.Decode(publicKeyBase58)
		privateKey := base58.Decode(privateKeyBase58)

		// Decode the JSON body into the struct
		var data models.HttpRequest
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			logrus.Errorf("error: %s", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.HttpResponse{
				Payload: fmt.Sprintf("%s", err),
			})
			return
		}
		defer r.Body.Close()

		// Perform Signature
		ecc := &crypto.Ed25519Crypto{}
		signature, err := ecc.Sign(privateKey, publicKey, []byte(data.Data))
		if err != nil {
			logrus.Errorf("error: %s", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.HttpResponse{
				Payload: fmt.Sprintf("%s", err),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: hex.EncodeToString(signature),
		})
	}
}

// Ed25519 Verification. Given a public key and a signature,
// this method will verify the signature and return the verification.
//
//	@param  s - server.Server
//	@return http.HandlerFunc
func Ed25519Verification(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Ed25519Verification")

		// Extract Keypair & Signature
		publicKeyBase58 := r.Header.Get("public-key")
		signatureHex := r.Header.Get("signature")

		// Decode to a slice of bytes
		publicKey := base58.Decode(publicKeyBase58)

		// Decode to a slice of bytes
		signature, err := hex.DecodeString(signatureHex)
		if err != nil {
			logrus.Errorf("error: %s", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.HttpResponse{
				Payload: fmt.Sprintf("%s", err),
			})
			return
		}

		// Decode the JSON body into the struct
		var data models.HttpRequest
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			logrus.Errorf("error: %s", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.HttpResponse{
				Payload: fmt.Sprintf("%s", err),
			})
			return
		}
		defer r.Body.Close()

		// Perform Signature
		ecc := &crypto.Ed25519Crypto{}
		verification := ecc.Verify(publicKey, signature, []byte(data.Data))
		if err != nil {
			logrus.Errorf("error: %s", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.HttpResponse{
				Payload: fmt.Sprintf("%s", err),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: fmt.Sprintf("%t", verification),
		})
	}
}

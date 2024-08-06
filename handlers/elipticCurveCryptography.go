/*
elipticCurveCryptography.go
v0.1.0
1/2/24

This file defines the handler functions for AES based operations
*/
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/btcsuite/btcutil/base58"
	"github.com/sirupsen/logrus"
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
		ecc := &Ed25519Crypto{}

		publicKey, privateKey, err := ecc.GenerateKeyPair()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
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
			PrivateKey: base58.Encode(publicKey),
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: "config.Version",
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: "config.Version",
		})
	}
}

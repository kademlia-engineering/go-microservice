/*
aesEncryption.go
v0.1.0
1/2/24

This file defines the handler functions for AES based operations
*/
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"kademlia.io/crypto"
	"kademlia.io/models"
	"kademlia.io/server"
)

// Aes Encrypt, Given a cipher and a message,
// this api will return the encrypted message.
//
//	@param  s - server.Server
//	@return http.HandlerFunc
func AesEncrypt(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("AesEncrypt")

		// Extract Cipher
		cipher := r.Header.Get("cipher")

		// Decode the JSON body into the struct
		var reqData models.HttpRequest
		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		encrypted, err := crypto.EncryptData([]byte(reqData.Data), cipher)
		if err != nil {
			logrus.Info("Failed to encrypt data")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: string(encrypted),
		})
	}
}

// Aes Decrypt. Given a cipher and an encrypted message,
// this api will return the decrypted message.
//
//	@param  s - server.Server
//	@return http.HandlerFunc
func AesDecrypt(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("AesDecrypt")

		// Extract Cipher
		cipher := r.Header.Get("cipher")

		// Decode the JSON body into the struct
		var reqData models.HttpRequest
		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		decrypted, err := crypto.DecryptData(reqData.Data, cipher)
		if err != nil {
			logrus.Info("Failed to encrypt data")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: string(decrypted),
		})
	}
}

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: "config.Version",
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HttpResponse{
			Payload: "config.Version",
		})
	}
}

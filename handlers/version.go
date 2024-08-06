/*
version.go
v0.1.0
1/2/24

This file defines the handler function for the version endpoint
*/
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"kademlia.io/config"
	"kademlia.io/models"
	"kademlia.io/server"
)

// Version handler
//
//	@param  s - server.Server
//	@return http.HandlerFunc
func GetVersionRoute(s server.Server) http.HandlerFunc {
	config := config.Get()
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Version Request")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.VersionResponse{
			Version: config.Version,
		})
	}
}

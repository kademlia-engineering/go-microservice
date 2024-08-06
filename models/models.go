/*
models.go
v0.1.0
1/2/24

This file defines the data structures used by the web server
*/
package models

type VersionResponse struct {
	Version string `json:"version"`
}

type Keypair struct {
	Curve      string `json:"curve"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type HttpRequest struct {
	Data string `json:"data"`
}

type HttpResponse struct {
	Payload string `json:"payload"`
}

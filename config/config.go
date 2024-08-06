/*
config.go
v0.1.0
1/2/24

This file exposes the system versioning
*/
package config

const VERSION = "0.1.0"

type Config struct {
	Version string `json:"version"`
}

func Get() *Config {
	return &Config{
		Version: VERSION,
	}
}

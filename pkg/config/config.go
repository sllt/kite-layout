package config

import (
	"os"
)

// SetupKiteEnv ensures Kite can find the configs/.env file.
// Call this before kite.New() to set the CONFIGS_DIR environment variable.
func SetupKiteEnv() {
	if os.Getenv("CONFIGS_DIR") == "" {
		os.Setenv("CONFIGS_DIR", "./configs")
	}
}

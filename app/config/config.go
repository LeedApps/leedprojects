package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LPConfigField is a field in the configuration
type LPConfigField struct {
	Env     string
	Default string
}

// LPConfig to be embedded in all config structs
type LPConfig struct {
	Fields map[string]LPConfigField
	Values map[string]string
}

// DBConfig stores database related information
var DBConfig LPConfig = LPConfig{
	Fields: map[string]LPConfigField{
		"Host": {
			"DB_HOST",
			"localhost",
		},
		"Port": {
			"DB_PORT",
			"5432",
		},
		"User": {
			"DB_USER",
			"postgres",
		},
		"Password": {
			Env: "DB_PASSWORD",
		},
		"SslMode": {
			"DB_SSLMODE",
			"disable",
		},
		"Database": {
			"DB_DATABASE",
			"lpbackend_development",
		},
	},
	Values: map[string]string{},
}

// Load loads the configuration
func (lpconfig LPConfig) Load() {
	LoadEnv()
	for i, v := range lpconfig.Fields {
		lpconfig.Values[i] = v.Default
		if v.Env != "" {
			if env := os.Getenv(v.Env); env != "" {
				lpconfig.Values[i] = env
			}
		}
	}
}

// LoadEnv loads the correct environment variables using the application environment
func LoadEnv() {
	environment := os.Getenv("LP_ENV")
	if environment == "" {
		environment = "development"
	}
	godotenv.Load(environment + ".env")
}

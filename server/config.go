package server

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	BindAddress      string `toml:"bind_address"`
	LogLevel         string `toml:"log_level"`
	NodeMultiaddress string `toml:"node_multiaddress"`
	Rendezvous       string `toml:"rendezvous"`
	SessionKey       string `toml:"session_key"`
}

// viperEnvVariable loads db information from .env file
func viperEnvVariable(key, default_value string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		return default_value
	}

	return value
}

// NewConfig creates a new config based on default values or provided .env file
func NewConfig() *Config {
	BindAddress := viperEnvVariable("BIND_ADDRESS", ":8000")
	LogLevel := viperEnvVariable("LOG_LEVEL", "debug")
	NodeMultiaddress := viperEnvVariable("NODE_MULTIADDRESS", "/ip4/127.0.0.1/tcp/0")
	Rendezvous := viperEnvVariable("NODE_MULTIADDRESS", "filecoin-p2p-oracle")
	SessionKey := viperEnvVariable("SESSION_KEY", "go")

	return &Config{
		BindAddress:      BindAddress,
		LogLevel:         LogLevel,
		NodeMultiaddress: NodeMultiaddress,
		Rendezvous:       Rendezvous,
		SessionKey:       SessionKey,
	}
}

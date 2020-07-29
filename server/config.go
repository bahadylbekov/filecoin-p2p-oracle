package server

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       string `toml:"server_port"`
	LogLevel         string `toml:"log_level"`
	NodeMultiaddress string `toml:"node_multiaddress"`
	Rendezvous       string `toml:"rendezvous"`
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
	ServerPort := viperEnvVariable("SERVER_PORT", ":8000")
	LogLevel := viperEnvVariable("LOG_LEVEL", "debug")
	NodeMultiaddress := viperEnvVariable("NODE_MULTIADDRESS", "/ip4/127.0.0.1/tcp/0")
	Rendezvous := viperEnvVariable("NODE_MULTIADDRESS", "filecoin-p2p-oracle")

	return &Config{
		ServerPort:       ServerPort,
		LogLevel:         LogLevel,
		NodeMultiaddress: NodeMultiaddress,
		Rendezvous:       Rendezvous,
	}
}

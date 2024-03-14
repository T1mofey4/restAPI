package apiserver

import "github.com/T1mofey4/restAPI/internal/app/store"

// Config
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// New Config
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
package apiserver

import "github.com/T1mofey4/restAPI/internal/app/store/sqlstore"

// Config
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *sqlstore.Config
}

// New Config
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
		Store:    sqlstore.NewConfig(),
	}
}

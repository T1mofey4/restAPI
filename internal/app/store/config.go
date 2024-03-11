package store

// Config
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// New config
func NewConfig() *Config {
	return &Config{}
}

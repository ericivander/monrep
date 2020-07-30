package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

// Postgres contains postgreSQL connection options
type Postgres struct {
	Driver string `env:"DATABASE_DRIVER,default=postgres"`
	URL    string `env:"DATABASE_URL,required"`
}

// Config contains configuration from environment variables
type Config struct {
	Postgres Postgres
}

// NewConfig creates new Config instance from environment variables
func NewConfig() *Config {
	var cfg Config
	gotenv.Load(".env")
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

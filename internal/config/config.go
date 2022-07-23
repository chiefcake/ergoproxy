package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const prefix = "ergoproxy"

// Config contains the config values for the application.
type Config struct {
	ServerHost string `default:"" split_words:"true"`
	ServerPort string `default:"8080" split_words:"true"`
}

// New parses config from environment variables and returns an instance of config or an error.
func New() (*Config, error) {
	var cfg Config

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not read config from environment variables")
	}

	return &cfg, nil
}

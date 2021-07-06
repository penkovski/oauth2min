package config

import "time"

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Host         string
	Port         string        `default:"8080"`
	ReadTimeout  time.Duration `default:"5s"`
	WriteTimeout time.Duration `default:"5s"`
	IdleTimeout  time.Duration `default:"60s"`
}

package config

import "time"

type Config struct {
	HTTP    httpConfig
	OAuth   oauthConfig
	Storage storageConfig
}

type httpConfig struct {
	Host         string        `envconfig:"HTTP_HOST"`
	Port         string        `envconfig:"HTTP_PORT" default:"8080"`
	ReadTimeout  time.Duration `default:"10s"`
	WriteTimeout time.Duration `default:"10s"`
	IdleTimeout  time.Duration `default:"60s"`
}

type oauthConfig struct {
	DefaultTokenFormat string        `envconfig:"DEFAULT_TOKEN_FORMAT" default:"jwt"`
	JwtExpiration      time.Duration `envconfig:"JWT_EXPIRATION" default:"1h"`
	JwtKey             jwtKeyConfig
}

type jwtKeyConfig struct {
	ID      string `envconfig:"KEY_ID" required:"true"`      // key id
	Public  string `envconfig:"PUBLIC_KEY" required:"true"`  // public key
	Private string `envconfig:"PRIVATE_KEY" required:"true"` // private key
}

type storageConfig struct {
	URI  string `envconfig:"STORAGE_CONNECT_URI"`
	User string `envconfig:"STORAGE_USER"`
	Pass string `envconfig:"STORAGE_PASS"`
	Type string `envconfig:"STORAGE_TYPE" default:"memory"`
}

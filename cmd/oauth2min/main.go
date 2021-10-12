package main

import (
	"net/http"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/penkovski/graceful"
	"github.com/rs/zerolog"

	"github.com/penkovski/oauth2min/internal/config"
	"github.com/penkovski/oauth2min/internal/oauth2"
	"github.com/penkovski/oauth2min/internal/server"
)

var Version = "0.0.0+development"

func main() {
	// initialize structured logger
	logger := zerolog.New(os.Stderr)

	// load configuration from environment
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logger.Fatal().Err(err).Send()
	}

	// create oauth service
	oauth := oauth2.New()

	// create http handler
	server := server.New(oauth, logger)

	addr := cfg.HTTP.Host + ":" + cfg.HTTP.Port
	srv := &http.Server{
		Addr:         addr,
		Handler:      server,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
	}

	logger.Info().Str("addr", addr).Str("version", Version).Msg("starting oauth2 server")

	if err := graceful.Shutdown(srv, 10*time.Second); err != nil {
		logger.Error().Err(err)
	}

	logger.Info().Msg("bye bye")
}

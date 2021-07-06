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

func main() {
	// initialize structured logger
	logger := zerolog.New(os.Stderr)

	// load configuration from environment
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logger.Fatal().Err(err).Send()
	}

	oauth := oauth2.New()
	handler := server.New(oauth)

	addr := cfg.HTTP.Host + ":" + cfg.HTTP.Port
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
	}

	logger.Info().Str("addr", addr).Msg("starting oauth2 server")

	if err := graceful.Shutdown(srv, 10*time.Second); err != nil {
		logger.Error().Err(err)
	}

	logger.Info().Msg("bye bye")
}

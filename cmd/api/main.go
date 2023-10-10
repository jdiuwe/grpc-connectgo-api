package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-server-demo/wallet/pkg/jwt"

	"github.com/grpc-server-demo/wallet/db/postgresql"
	db "github.com/grpc-server-demo/wallet/db/sqlc"

	"github.com/grpc-server-demo/wallet/service/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

var (
	configPath = pflag.String("config", "./cmd/server/", "path to configuration .env file")
	debugMode  = pflag.Bool("debug", true, "debug mode")
)

func main() {
	pflag.Parse()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// read configuration
	cfg, err := LoadConfig(*configPath)
	if err != nil {
		logger.Panic().Err(err).Msg("failed to load configuration")
	}

	fmt.Println(cfg)

	if cfg.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	psql, err := postgresql.NewPostgresqlClient(&postgresql.Config{
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Name:     cfg.DBName,
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
	})
	if err != nil {
		logger.Panic().Err(err).Msg("could not connect to a database")
	}

	store := db.NewStore(psql)

	jwtManager := jwt.NewManager("jwtsecret", time.Hour)

	// start gRPC server
	srv, err := api.Run(&api.Config{ListenerPort: cfg.GRPCListenerPort}, &logger, store, jwtManager)
	if err != nil {
		log.Fatal().Msgf("could not start the server: %v", err)
	}

	defer func(srv *api.Server) {
		if err := srv.Shutdown(); err != nil {
			log.Fatal().Msgf("could not stop the server: %v", err)
		}
	}(srv)

	log.Info().Msg("Server OK")

	<-ctx.Done()

	log.Warn().Msg("shutting down")
}

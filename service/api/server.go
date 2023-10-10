package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	db "github.com/grpc-server-demo/wallet/db/sqlc"
	"github.com/grpc-server-demo/wallet/gen/go/api/user/v1/userv1connect"
	"github.com/grpc-server-demo/wallet/pkg/jwt"
	"github.com/grpc-server-demo/wallet/service/api/handlers"
	"github.com/grpc-server-demo/wallet/service/api/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	server *http.Server
}

type Config struct {
	ListenerPort string `yaml:"listenerPort"`
}

func Run(cfg *Config, logger *zerolog.Logger, store *db.Store, jwtManager *jwt.Manager) (*Server, error) {
	if cfg == nil {
		return nil, handlers.ErrNilServerConfig
	}

	srv := &Server{}

	interceptors := connect.WithInterceptors(
		middleware.NewLoggingInterceptor(logger),
		middleware.NewAuthInterceptor(jwtManager),
	)

	mux := http.NewServeMux()

	userHandlers := handlers.NewUserHandler(store, jwtManager)
	userPath, accountsHandler := userv1connect.NewUserServiceHandler(userHandlers, interceptors)
	mux.Handle(userPath, accountsHandler)

	reflector := grpcreflect.NewStaticReflector(
		"api.signing.v1.SigningService",
		"api.user.v1.UserService",
		"api.wallet.v1.WalletService",
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.ListenerPort),
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	srv.server = server

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			panic(err)
		}
	}()

	return srv, nil
}

func (srv *Server) Shutdown() error {
	return srv.server.Close()
}

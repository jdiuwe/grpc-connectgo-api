package middleware

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/grpc-connectgo-api-demo/wallet/pkg/jwt"

	"connectrpc.com/connect"
	"github.com/grpc-connectgo-api-demo/wallet/gen/go/api/user/v1/userv1connect"
)

var (
	ErrNoTokenProvided       = errors.New("no access token provided")
	ErrTokenValidationFailed = errors.New("access token validation failed")
	ErrTokenExpired          = errors.New("access token expired")
)

const (
	AuthorizationHeader = "Authorization"
	BearerHeader        = "Bearer"
)

func NewAuthInterceptor(jwtManager *jwt.Manager) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			// require authentication by default
			// only a small subset of requests does not require an access token
			switch request.Spec().Procedure {
			case userv1connect.UserServiceRegisterUserProcedure, userv1connect.UserServiceLoginUserProcedure:
				return next(ctx, request)
			default:
				// require authentication
			}

			authHeader := request.Header().Get(AuthorizationHeader)

			splitToken := strings.Split(authHeader, BearerHeader)
			if len(splitToken) != 2 {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					ErrNoTokenProvided,
				)
			}

			reqToken := strings.TrimSpace(splitToken[1])

			verify, err := jwtManager.Verify(reqToken)
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					ErrTokenValidationFailed,
				)
			}

			if !verify.VerifyExpiresAt(time.Now().Unix(), true) {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					ErrTokenExpired,
				)
			}

			return next(ctx, request)
		}
	}
}

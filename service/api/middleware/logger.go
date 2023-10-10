package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"connectrpc.com/connect"
	"github.com/rs/zerolog"
)

func NewLoggingInterceptor(l *zerolog.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			logger := l.Info()
			statusCode := codes.Unknown

			resp, err := next(ctx, req)

			if err != nil {
				logger = l.Error().Err(err)

				if st, ok := status.FromError(err); ok {
					statusCode = st.Code()

					switch st.Code() {
					case codes.PermissionDenied, codes.Unauthenticated, codes.NotFound, codes.AlreadyExists:
						logger = l.Warn().Err(err)
					case codes.Internal, codes.Aborted, codes.InvalidArgument:
						logger = l.Error().Err(err)
					case codes.Unimplemented:
						logger = l.Debug().Err(err)
					}
				}
			}

			logger.
				Str("procedure", req.Spec().Procedure).
				Str("protocol", req.Peer().Protocol).
				Str("addr", req.Peer().Addr).
				Int("status_code", int(statusCode)).
				Str("status", statusCode.String()).
				Dur("duration", time.Since(start)).
				Msg("request")

			return resp, err
		}
	}
}

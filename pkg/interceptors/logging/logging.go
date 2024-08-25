package logging

import (
	"context"
	"fmt"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func InterceptorLogger(l zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l := l.With().Fields(fields).Logger()

		switch lvl {
		case logging.LevelDebug:
			l.Debug().Msg(msg)
		case logging.LevelInfo:
			l.Info().Msg(msg)
		case logging.LevelWarn:
			l.Warn().Msg(msg)
		case logging.LevelError:
			l.Error().Msg(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func CreateServerLogInterceptor() grpc.ServerOption{
	logger := zerolog.New(os.Stdout)

	options := []logging.Option {
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	interceptor := grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogger(logger), options...),)

	return interceptor
}

func CreateClientLogInterceptor() grpc.DialOption{
	logger := zerolog.New(os.Stdout)

	options := []logging.Option {
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	interceptor := grpc.WithChainUnaryInterceptor(
		logging.UnaryClientInterceptor(InterceptorLogger(logger), options...),)

	return interceptor
}
package jwtvalidation

import (
	"context"

	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/pkg/jwt"
	"github.com/dsantaguida/idle-clicker/proto/authentication"
	"github.com/dsantaguida/idle-clicker/proto/bank"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"google.golang.org/grpc"
)

func CreateValidationInterceptor() grpc.ServerOption{
	return grpc.ChainUnaryInterceptor(
		selector.UnaryServerInterceptor(
			auth.UnaryServerInterceptor(ValidationInterceptor),
			selector.MatchFunc(ValidationMatcher),
		),
	)
}

func ValidationInterceptor(ctx context.Context) (context.Context, error) {
	token, err := jwt.GetTokenFromContext(ctx)
	if err != nil {
		return ctx, err
	}

	err = jwt.Validate(token)
	if err != nil {
		return ctx, idle_errors.ErrInvalidToken
	}

	return ctx, nil
}

//Run validation on the entire bank service, and the update password call of auth service
//TODO: Better way to do this, each service should have its own validation method that we call here
func ValidationMatcher(ctx context.Context, callMeta interceptors.CallMeta) bool {
	if callMeta.Service == bank.BankService_ServiceDesc.ServiceName {
		return true
	} else if callMeta.Service == authentication.AuthenticationService_ServiceDesc.ServiceName {
		return callMeta.Method == "UpdatePassword"
	}

	return true
}

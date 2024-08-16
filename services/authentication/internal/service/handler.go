package service

import (
	"context"

	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/pkg/jwt"
	authenticationService "github.com/dsantaguida/idle-clicker/proto/authentication"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/models"
)

func (b *AuthenticationServiceServer) Register(ctx context.Context, userRequest *authenticationService.UserRequest) (*authenticationService.RegisterResponse, error) {
	user := models.CreateUser(userRequest.User.Username, userRequest.User.Password)

	_, err := b.db.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &authenticationService.RegisterResponse{Result: true}, nil
}

func (b *AuthenticationServiceServer) Login(ctx context.Context, userRequest *authenticationService.UserRequest) (*authenticationService.LoginResponse, error) {
	user := models.CreateUser(userRequest.User.Username, userRequest.User.Password)

	user, err := b.db.FindUser(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	tokenString, err := jwt.CreateToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &authenticationService.LoginResponse{Token: tokenString}, nil
}

func (b *AuthenticationServiceServer) UpdatePassword(ctx context.Context, userRequest *authenticationService.UpdatePasswordRequest) (*authenticationService.UpdatePasswordResponse, error) {
	user := models.CreateUser(userRequest.User.Username, userRequest.User.Password)

	user, err := b.db.FindUser(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	user, err = b.db.UpdateUserPassword(ctx, user, userRequest.NewPassword)
	if err != nil {
		return nil, err
	}

	if user.Password != userRequest.NewPassword {
		return nil, idle_errors.ErrPasswordNotUpdated
	}

	return &authenticationService.UpdatePasswordResponse{User: b.userModelToProto(user)}, nil
}

func (*AuthenticationServiceServer) userModelToProto(user *models.User) (*authenticationService.User){
	return &authenticationService.User{Username: user.Username, Password: user.Password}
}
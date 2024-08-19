package client

import (
	"context"
	"errors"

	auth_proto "github.com/dsantaguida/idle-clicker/proto/authentication"
	bank_proto "github.com/dsantaguida/idle-clicker/proto/bank"
	"google.golang.org/grpc"
)

type IdleClient struct {
	authenticationClient auth_proto.AuthenticationServiceClient;
	bankClient bank_proto.BankServiceClient;
	authConn *grpc.ClientConn;
	bankConn *grpc.ClientConn;
}

func CreateIdleClient(authConn *grpc.ClientConn, bankConn *grpc.ClientConn) (*IdleClient) {
	authClient := auth_proto.NewAuthenticationServiceClient(authConn)
	bankClient := bank_proto.NewBankServiceClient(bankConn)

	return &IdleClient{
		authenticationClient: authClient, 
		bankClient: bankClient,
		bankConn: bankConn,
		authConn: authConn,
	}
}

func (client *IdleClient) Close() {
	client.authConn.Close();
	client.bankConn.Close();
}

func (client *IdleClient) Register(ctx context.Context, username string, password string) error {
	user := &auth_proto.User{Username: username, Password: password}
	userRequest := &auth_proto.UserRequest{User: user}

	registerResponse, err := client.authenticationClient.Register(ctx, userRequest)
	if err != nil {
		return err
	}
	if registerResponse.Token == "" {
		return errors.New("failed to create user")
	}

	bankRequest := &bank_proto.BankRequest{Token: registerResponse.Token}
	bankResponse, err := client.bankClient.CreateBank(ctx, bankRequest)
	if err != nil {
		return err
	}
	if len(bankResponse.Bank.Id) == 0 {
		return errors.New("failed to create bank")
	}

	return nil
}

func (client *IdleClient) Login(ctx context.Context, username string, password string) (string, int, error) {
	user := &auth_proto.User{Username: username, Password: password}
	userRequest := &auth_proto.UserRequest{User: user}

	loginResponse, err := client.authenticationClient.Login(ctx, userRequest)
	if err != nil {
		return "", -1, err
	}
	if len(loginResponse.Token) == 0 {
		return "", -1, errors.New("failed to get login token")
	}

	getBankDataRequest := &bank_proto.GetBankDataRequest{Token: loginResponse.Token}
	bankResponse, err := client.bankClient.GetBankData(ctx, getBankDataRequest)
	if err != nil {
		return "", -1, err
	}

	return loginResponse.Token, int(bankResponse.GetBank().Value), nil
}

func (client *IdleClient) UpdateBankValue(ctx context.Context, token string, value int64) error {
	setBankDataRequest := &bank_proto.SetBankDataRequest{Token: token, Value: value}

	bankResponse, err := client.bankClient.SetBankData(ctx, setBankDataRequest)
	if err != nil {
		return err
	}
	if bankResponse.Bank.Value != value {
		return errors.New("failed to set bank value")
	}

	return nil
}
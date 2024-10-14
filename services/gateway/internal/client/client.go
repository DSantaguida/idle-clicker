package client

import (
	"context"
	"errors"

	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/pkg/jwt"
	auth_proto "github.com/dsantaguida/idle-clicker/proto/authentication"
	bank_proto "github.com/dsantaguida/idle-clicker/proto/bank"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	var header metadata.MD

	_, err := client.authenticationClient.Register(ctx, userRequest, grpc.Header(&header))
	if err != nil {
		return err
	}

	token := header.Get(jwt.TOKEN_KEY)[0]
	if len(token) == 0 {
		return idle_errors.ErrTokenNotInHeader
	}

	if token == "" {
		return errors.New("failed to create user")
	}

	//TODO: Do I need to set the header of this request?
	header = metadata.Pairs(jwt.TOKEN_KEY, token)
	grpc.SetHeader(ctx, header)

	bankRequest := &bank_proto.BankRequest{}
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
	var header metadata.MD

	_, err := client.authenticationClient.Login(ctx, userRequest, grpc.Header(&header))
	if err != nil {
		return "", -1, err
	}

	token := header.Get(jwt.TOKEN_KEY)[0]
	if len(token) == 0 {
		return "", -1, idle_errors.ErrTokenNotInHeader
	}

	getBankDataRequest := &bank_proto.GetBankDataRequest{}
	ctx = metadata.AppendToOutgoingContext(ctx, jwt.TOKEN_KEY, token)
	bankResponse, err := client.bankClient.GetBankData(ctx, getBankDataRequest)
	if err != nil {
		return "", -1, err
	}

	return token, int(bankResponse.GetBank().Value), nil
}

func (client *IdleClient) UpdateBankValue(ctx context.Context, token string, value int64) error {
	setBankDataRequest := &bank_proto.SetBankDataRequest{Value: value}

	ctx = metadata.AppendToOutgoingContext(ctx, jwt.TOKEN_KEY, token)
	bankResponse, err := client.bankClient.SetBankData(ctx, setBankDataRequest)
	if err != nil {
		return err
	}
	if bankResponse.Bank.Value != value {
		return errors.New("failed to set bank value")
	}

	return nil
}
package client

import (
	"context"
	"fmt"

	"assesement-test-MicroServices/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AccountServiceClient struct {
	Client pb.AccountServiceClient
}

func InitAccountServiceClient(url string) AccountServiceClient {
	cc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := AccountServiceClient{
		Client: pb.NewAccountServiceClient(cc),
	}

	return c
}
func (c *AccountServiceClient) GetAccountById(productId string) (*pb.AccountResponse, error) {
	req := &pb.ById{
		Id: productId,
	}

	resp, err := c.Client.GetAccountById(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}

	return resp, nil
}

func (c *AccountServiceClient) UpdateAccount(req *pb.UpdateAccountRequest) (*pb.AccountResponse, error) {
	resp, err := c.Client.UpdateAccount(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to update account: %w", err)
	}

	return resp, nil
}
func (c *AccountServiceClient) CreateAccount(req *pb.CreateAccountRequest) (*pb.AccountResponse, error) {

	resp, err := c.Client.CreateAccount(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return resp, nil
}
func (c *AccountServiceClient) DeleteAccount(accountId string) (*pb.AccountResponse, error) {

	resp, err := c.Client.DeleteAccount(context.Background(), &pb.ById{Id: accountId})
	if err != nil {
		return nil, fmt.Errorf("failed to delete account: %w", err)
	}

	return resp, nil
}
func (c *AccountServiceClient) ListAccounts() (*pb.AccountResponseRepeated, error) {

	resp, err := c.Client.ListAccounts(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to update account: %w", err)
	}

	return resp, nil
}

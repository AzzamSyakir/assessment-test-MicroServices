package client

import (
	"context"
	"fmt"

	"assesement-test-MicroServices/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func InitUserServiceClient(url string) UserServiceClient {
	cc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := UserServiceClient{
		Client: pb.NewUserServiceClient(cc),
	}

	return c
}
func (c *UserServiceClient) GetUserById(id string) (*pb.UserResponse, error) {
	req := &pb.ById{
		Id: id,
	}

	resp, err := c.Client.GetUserById(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}

	return resp, nil
}
func (c *UserServiceClient) UpdateUser(req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	resp, err := c.Client.UpdateUser(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to update account: %w", err)
	}

	return resp, nil
}
func (c *UserServiceClient) CreateUser(req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	resp, err := c.Client.CreateUser(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return resp, nil
}
func (c *UserServiceClient) DeleteUser(accountId string) (*pb.UserResponse, error) {

	resp, err := c.Client.DeleteUser(context.Background(), &pb.ById{Id: accountId})
	if err != nil {
		return nil, fmt.Errorf("failed to delete account: %w", err)
	}

	return resp, nil
}
func (c *UserServiceClient) ListUsers() (*pb.UserResponseRepeated, error) {

	resp, err := c.Client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to update account: %w", err)
	}

	return resp, nil
}

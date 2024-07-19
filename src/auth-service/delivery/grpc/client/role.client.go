package client

import (
	"context"
	"fmt"

	"assesement-test-MicroServices/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RoleServiceClient struct {
	Client pb.RoleServiceClient
}

func InitRoleServiceClient(url string) RoleServiceClient {
	cc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := RoleServiceClient{
		Client: pb.NewRoleServiceClient(cc),
	}

	return c
}
func (c *RoleServiceClient) GetRoleById(productId string) (*pb.RoleResponse, error) {
	req := &pb.ById{
		Id: productId,
	}

	resp, err := c.Client.GetRoleById(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get role by id: %w", err)
	}

	return resp, nil
}

func (c *RoleServiceClient) UpdateRole(req *pb.UpdateRoleRequest) (*pb.RoleResponse, error) {
	resp, err := c.Client.UpdateRole(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	return resp, nil
}
func (c *RoleServiceClient) CreateRole(req *pb.CreateRoleRequest) (*pb.RoleResponse, error) {

	resp, err := c.Client.CreateRole(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	return resp, nil
}
func (c *RoleServiceClient) DeleteRole(roleId string) (*pb.RoleResponse, error) {

	resp, err := c.Client.DeleteRole(context.Background(), &pb.ById{Id: roleId})
	if err != nil {
		return nil, fmt.Errorf("failed to delete role: %w", err)
	}

	return resp, nil
}
func (c *RoleServiceClient) ListRoles() (*pb.RoleResponseRepeated, error) {

	resp, err := c.Client.ListRoles(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	return resp, nil
}

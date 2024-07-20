package client

import (
	"context"
	"fmt"

	"assesement-test-MicroServices/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OfficeServiceClient struct {
	Client pb.OfficeServiceClient
}

func InitOfficeServiceClient(url string) OfficeServiceClient {
	cc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := OfficeServiceClient{
		Client: pb.NewOfficeServiceClient(cc),
	}

	return c
}
func (c *OfficeServiceClient) GetOfficeById(productId string) (*pb.OfficeResponse, error) {
	req := &pb.ById{
		Id: productId,
	}

	resp, err := c.Client.GetOfficeById(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get office by id: %w", err)
	}

	return resp, nil
}

func (c *OfficeServiceClient) UpdateOffice(req *pb.UpdateOfficeRequest) (*pb.OfficeResponse, error) {
	resp, err := c.Client.UpdateOffice(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to update office: %w", err)
	}

	return resp, nil
}
func (c *OfficeServiceClient) CreateOffice(req *pb.CreateOfficeRequest) (*pb.OfficeResponse, error) {

	resp, err := c.Client.CreateOffice(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to create office: %w", err)
	}

	return resp, nil
}
func (c *OfficeServiceClient) DeleteOffice(officeId string) (*pb.OfficeResponse, error) {

	resp, err := c.Client.DeleteOffice(context.Background(), &pb.ById{Id: officeId})
	if err != nil {
		return nil, fmt.Errorf("failed to delete office: %w", err)
	}

	return resp, nil
}
func (c *OfficeServiceClient) ListOffices() (*pb.OfficeResponseRepeated, error) {

	resp, err := c.Client.ListOffices(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to update office: %w", err)
	}

	return resp, nil
}

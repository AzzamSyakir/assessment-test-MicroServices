package client

import (
	"context"
	"fmt"

	"assesement-test-MicroServices/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ScreenServiceClient struct {
	Client pb.ScreenServiceClient
}

func InitScreenServiceClient(url string) ScreenServiceClient {
	cc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ScreenServiceClient{
		Client: pb.NewScreenServiceClient(cc),
	}

	return c
}
func (c *ScreenServiceClient) GetScreenById(productId string) (*pb.ScreenResponse, error) {
	req := &pb.ById{
		Id: productId,
	}

	resp, err := c.Client.GetScreenById(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to get screen by id: %w", err)
	}

	return resp, nil
}

func (c *ScreenServiceClient) UpdateScreen(req *pb.UpdateScreenRequest) (*pb.ScreenResponse, error) {
	resp, err := c.Client.UpdateScreen(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to update screen: %w", err)
	}

	return resp, nil
}
func (c *ScreenServiceClient) CreateScreen(req *pb.CreateScreenRequest) (*pb.ScreenResponse, error) {

	resp, err := c.Client.CreateScreen(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to create screen: %w", err)
	}

	return resp, nil
}
func (c *ScreenServiceClient) DeleteScreen(screenId string) (*pb.ScreenResponse, error) {

	resp, err := c.Client.DeleteScreen(context.Background(), &pb.ById{Id: screenId})
	if err != nil {
		return nil, fmt.Errorf("failed to delete screen: %w", err)
	}

	return resp, nil
}
func (c *ScreenServiceClient) ListScreens() (*pb.ScreenResponseRepeated, error) {

	resp, err := c.Client.ListScreens(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to update screen: %w", err)
	}

	return resp, nil
}

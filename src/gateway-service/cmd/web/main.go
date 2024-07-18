package main

import (
	"assesement-test-MicroServices/src/gateway-service/container"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Gateway Services started.")

	webContainer := container.NewWebContainer()

	address := fmt.Sprintf(
		"%s:%s",
		webContainer.Env.App.Host,
		webContainer.Env.App.Port,
	)
	netListen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	fmt.Println("Gateway Services listen", address)

	if err := webContainer.Grpc.Serve(netListen); err != nil {
		log.Fatalf("failed to serve %v", err.Error())
	}
	fmt.Println("Gateway Services finished.")
}

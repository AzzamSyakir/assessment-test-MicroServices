package main

import (
	"assesement-test-MicroServices/src/auth-service/container"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Auth-Employee Services started.")

	webContainer := container.NewWebContainer()

	address := fmt.Sprintf(
		"%s:%s",
		webContainer.Env.App.Host,
		webContainer.Env.App.AuthPort,
	)
	netListen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	fmt.Println("Auth-Employee Services listen", address)

	if err := webContainer.Grpc.Serve(netListen); err != nil {
		log.Fatalf("failed to serve %v", err.Error())
	}
	fmt.Println("Auth-Employee Services finished.")
}

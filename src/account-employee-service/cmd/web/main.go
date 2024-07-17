package main

import (
	"assesement-test-MicroServices/src/account-employee-service/container"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Account-Employee Services started.")

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
	if err := webContainer.Grpc.Serve(netListen); err != nil {
		log.Fatalf("failed to serve %v", err.Error())
	}
	fmt.Println("Account-Employee Services finished.")
}

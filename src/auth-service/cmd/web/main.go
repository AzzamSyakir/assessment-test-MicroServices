package main

import (
	"assesement-test-MicroServices/src/auth-service/container"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Auth Services started.")

	webContainer := container.NewWebContainer()

	address := fmt.Sprintf(
		"%s:%s",
		"0.0.0.0",
		webContainer.Env.App.AuthPort,
	)
	fmt.Println("Auth Services listen", address)
	listenAndServeErr := http.ListenAndServe(address, webContainer.Route.Router)
	if listenAndServeErr != nil {
		fmt.Println("Failed to start Auth Services:", listenAndServeErr)
		panic(listenAndServeErr)
	}
	fmt.Println("Auth Services finished.")
}

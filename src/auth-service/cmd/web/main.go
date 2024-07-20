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
		webContainer.Env.App.Host,
		webContainer.Env.App.AuthPort,
	)
	listenAndServeErr := http.ListenAndServe(address, webContainer.Route.Router)
	if listenAndServeErr != nil {
		panic(listenAndServeErr)
	}
	fmt.Println("Auth Services listen ", address)
	fmt.Println("Auth Services finished.")
}

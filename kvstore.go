package main

import (
	"fmt"
	"kvstore/config"
	"kvstore/server"
	"net/http"
)

func main() {
	configuration := config.Setup()

	server.Setup()
	port := fmt.Sprintf(":%s", configuration.Port)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

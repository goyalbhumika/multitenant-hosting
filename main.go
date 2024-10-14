package main

import (
	"fmt"
	"multitenant-hosting/config"
	"net/http"
	"os"

	"multitenant-hosting/handlers"
	"multitenant-hosting/service"
	repository "multitenant-hosting/store"
)

func main() {
	config.SetConfig()
	store := repository.NewStore()
	registry := service.NewRegistry(store)
	http.HandleFunc("/v1/apps", handlers.CreateAppHandler(registry))

	// Start the main platform server
	fmt.Println("Platform server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start platform server: %s\n", err)
		os.Exit(1)
	}
}

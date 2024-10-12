package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type LocalAppInstance struct {
	mu   sync.Mutex
	port int
}

// Handle the "Hello World" for each app instance
func startAppInstance(ctx context.Context, appID string, port int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World %s\n", appID)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("App %s running on port %d\n", appID, port)

	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start app instance %s on port %d: %v", appID, port, err)
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-time.After(3 * time.Second):
		// wait for server to start
	}
	return nil
}

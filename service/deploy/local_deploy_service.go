package service

import (
	"context"
	"fmt"
	"log"
	"multitenant-hosting/domain"
	"net/http"
	"sync"
	"time"
)

type localDeploySvc struct {
	mu   sync.Mutex
	Port int
}

func NewlocalDeploySvc(port int) DeployInstance {
	return &localDeploySvc{Port: port}
}

// Handle the "Hello World" for each app instance
func (svc *localDeploySvc) DeployAppInstance(ctx context.Context, appID string) (*domain.DeployResponse, error) {
	svc.mu.Lock()
	svc.Port++
	port := svc.Port
	svc.mu.Unlock()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World %s\n", appID)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("App %s running on Port %d\n", appID, port)

	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start app instance %s on Port %d: %v", appID, svc.Port, err)
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case <-time.After(1 * time.Second):
		// wait for server to start
	}
	return &domain.DeployResponse{
		Port: port,
		DNS:  fmt.Sprintf("%s.gravityfalls42.hitchhiker", appID),
	}, nil
}

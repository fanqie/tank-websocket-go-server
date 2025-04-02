package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fanqie/tank-websocket-go-server/pkg"
)

// StartServerSingle starts the WebSocket server
func StartServerSingle(addr string) (*pkg.Manager, error) {
	manager := pkg.GetSingleInstance()
	go manager.Start()

	// Handle error events
	go handleErrors(manager)

	// Handle connection events
	go handleConnectionEvents(manager)

	// Create HTTP server
	server := &http.Server{
		Addr: addr,
	}

	// Set HTTP server reference for shutdown
	manager.SetHTTPServer(server)

	// Register handler
	http.HandleFunc("/ws", manager.HandleConnection)

	// Start HTTP server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server startup failed: %v", err)
		}
	}()

	log.Printf("WebSocket server started on %s", addr)

	return manager, nil
}

// StartNewServer starts a new WebSocket server instance
func StartNewServer(addr string) (*pkg.Manager, error) {
	manager := pkg.NewInstance()
	go manager.Start()

	// Handle error events
	go handleErrors(manager)

	// Handle connection events
	go handleConnectionEvents(manager)

	// Create HTTP server
	server := &http.Server{
		Addr: addr,
	}

	// Set HTTP server reference for shutdown
	manager.SetHTTPServer(server)

	// Register handler
	http.HandleFunc("/ws", manager.HandleConnection)

	// Start HTTP server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server startup failed: %v", err)
		}
	}()

	log.Printf("WebSocket server started on %s", addr)

	return manager, nil
}

// Handle error events
func handleErrors(manager *pkg.Manager) {
	for err := range manager.Errors {
		log.Printf("WebSocket error: [Code: %d] %s, Time: %v",
			err.Code, err.Message, err.Time.Format("2006-01-02 15:04:05"))
	}
}

// Handle connection events
func handleConnectionEvents(manager *pkg.Manager) {
	for event := range manager.ConnEvents {
		switch event.EventType {
		case "connect":
			log.Printf("New connection: UserID=%s, Time=%s",
				event.UserID, event.Time.Format("2006-01-02 15:04:05"))
		case "disconnect":
			log.Printf("Disconnection: UserID=%s, Time=%s",
				event.UserID, event.Time.Format("2006-01-02 15:04:05"))
		case "subscribe":
			log.Printf("Topic subscription: UserID=%s, Topic=%s, Time=%s",
				event.UserID, event.Topic, event.Time.Format("2006-01-02 15:04:05"))
		case "unsubscribe":
			log.Printf("Topic unsubscription: UserID=%s, Topic=%s, Time=%s",
				event.UserID, event.Topic, event.Time.Format("2006-01-02 15:04:05"))
		}
	}
}

func main() {
	// Start server
	manager, err := StartServerSingle(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Print status information
	log.Printf("Server status: Running=%v", manager.IsRunning())
	// manager.DisableDebug()
	// Broadcast messages
	// manager.BroadcastMessage([]byte("Hello, World!"), nil)
	manager.BroadcastTopicMessage("news", "Latest news")

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	// Receive syscall.SIGINT and syscall.SIGTERM signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a 5-second timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown WebSocket manager and HTTP server
	if err := manager.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server has been shut down successfully")
}

/* Multi-server example:
func multiServerExample() {
	// Start two different WebSocket server instances
	manager1, err1 := StartServerSingle(":8080")
	manager2, err2 := StartNewServer(":8081")

	if err1 != nil || err2 != nil {
		log.Fatal("Failed to start servers")
	}

	// Broadcast to different servers
	manager1.BroadcastMessage([]byte("Server 1 message"))
	manager2.BroadcastMessage([]byte("Server 2 message"))

	// Monitor connection counts
	go func() {
		for {
			fmt.Printf("Server 1 connections: %d, Server 2 connections: %d\n",
				manager1.GetClientCount(), manager2.GetClientCount())
			time.Sleep(5 * time.Second)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create a 5-second timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown both servers
	manager1.Shutdown(ctx)
	manager2.Shutdown(ctx)

	log.Println("All servers have been shut down")
}
*/

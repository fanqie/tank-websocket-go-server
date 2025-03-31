# Tank WebSocket Go Server

A lightweight, high-performance WebSocket server that supports both singleton and multi-instance modes, with topic subscription functionality.

## Features

- **Dual Mode Support**: Both singleton and multi-instance modes
- **Topic Subscription**: Clients can subscribe to specific topics to receive targeted messages
- **Error Handling**: Built-in error event channel for monitoring and handling errors
- **Connection Monitoring**: Provides connection status change events (connect, disconnect, subscribe, unsubscribe)
- **Graceful Shutdown**: Supports graceful server shutdown to ensure proper resource release
- **Status Statistics**: Provides client count, topic subscription count, and other statistical functions
- **Heartbeat Mechanism**: Keep-alive for long-running connections with configurable intervals
- **Authentication**: Customizable authentication function for connection security
- **Connection Management**: Ability to close specific client connections

## Installation and Usage

### Installation

```bash
go get github.com/fanqie/tank-websocket-go-server
```

### Usage Examples

#### Singleton Mode

```go
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

func main() {
	// Get singleton WebSocket manager
	tkws := pkg.GetSingleInstance()

	// Enable heartbeat (optional)
	tkws.EnableHeartbeat(30*time.Second, 60*time.Second)
	
	// Enable authentication (optional)
	tkws.EnableAuth(func(r *http.Request) bool {
		// Check token or other authentication logic
		token := r.URL.Query().Get("token")
		return token == "your-secret-token"
	})
	
	// Set up HTTP handler
	http.HandleFunc("/ws", tkws.HandleConnection)

	// Create HTTP server
	server := &http.Server{
		Addr: ":8080",
	}

	// Set HTTP server reference for shutdown
	tkws.SetHTTPServer(server)

	// Start HTTP server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server startup failed: %v", err)
		}
	}()

	log.Println("WebSocket server started on :8080")

	// Broadcast message to all clients
	tkws.BroadcastMessage([]byte("Welcome to Tank WebSocket Server"))

	// Broadcast message to specific topic subscribers
	tkws.BroadcastTopicMessage("news", "This is a news message")

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	tkws.Shutdown(ctx)
}
```

#### Multi-Instance Mode

```go
// Create multiple WebSocket server instances
tkws1 := pkg.NewInstance()
tkws2 := pkg.NewInstance()

go tkws1.Start()
go tkws2.Start()

// Handle WebSocket connections on different paths
http.HandleFunc("/ws1", tkws1.HandleConnection)
http.HandleFunc("/ws2", tkws2.HandleConnection)
```

## Client Connection

### Browser Client Example

```javascript
// Create WebSocket connection
const socket = new WebSocket('ws://localhost:8080/ws?user_id=user123&token=your-secret-token');

// When connection is open
socket.onopen = function(e) {
  console.log('WebSocket connection established');
  
  // Subscribe to topic
  socket.send('sub:news');
};

// Receive messages
socket.onmessage = function(event) {
  console.log('Message received:', event.data);
};

// Connection closed
socket.onclose = function(event) {
  console.log('WebSocket connection closed');
};

// Unsubscribe from topic
function unsubscribe() {
  socket.send('unsub:news');
}

// Send message
function sendMessage(message) {
  socket.send(message);
}
```

## API Reference

### Core Components

- **Manager**: WebSocket connection manager
- **Client**: Represents a single WebSocket connection
- **Subscription**: Represents a topic subscription

### Main Methods

| Method | Description |
|-----|------|
| `pkg.GetSingleInstance()` | Get singleton instance of WebSocket manager |
| `pkg.NewInstance()` | Create a new WebSocket manager instance (multi-instance mode) |
| `pkg.SetCustomUpgrader(upgrader)` | Set custom WebSocket upgrader |
| `tkws.Start()` | Start the WebSocket manager |
| `tkws.HandleConnection(w, r)` | Handle new WebSocket connection requests |
| `tkws.BroadcastMessage(message)` | Broadcast message to all connected clients |
| `tkws.BroadcastTopicMessage(topic, data)` | Broadcast message to subscribers of a specific topic |
| `tkws.Shutdown(ctx)` | Gracefully shut down WebSocket manager and HTTP server |
| `tkws.GetClientCount()` | Get number of currently connected clients |
| `tkws.GetTopicSubscriberCount(topic)` | Get number of subscribers for a specific topic |
| `tkws.GetAllTopics()` | Get all available topics |
| `tkws.IsRunning()` | Check if server is running |
| `tkws.EnableHeartbeat(interval, timeout)` | Enable heartbeat mechanism with specified interval and timeout |
| `tkws.DisableHeartbeat()` | Disable heartbeat mechanism |
| `tkws.EnableAuth(authFunc)` | Enable authentication with specified function |
| `tkws.DisableAuth()` | Disable authentication |
| `tkws.CloseClient(userID)` | Close connection to a specific client by user ID |

### Event Monitoring

```go
// Handle error events
go func() {
    for err := range tkws.Errors {
        log.Printf("WebSocket error: [Code: %d] %s", err.Code, err.Message)
    }
}()

// Handle connection events
go func() {
    for event := range tkws.ConnEvents {
        switch event.EventType {
        case "connect":
            log.Printf("New connection: UserID=%s", event.UserID)
        case "disconnect":
            log.Printf("Disconnection: UserID=%s", event.UserID)
        case "subscribe":
            log.Printf("Topic subscription: UserID=%s, Topic=%s", event.UserID, event.Topic)
        case "unsubscribe":
            log.Printf("Topic unsubscription: UserID=%s, Topic=%s", event.UserID, event.Topic)
        }
    }
}()
```

## Advanced Configuration

### Custom WebSocket Upgrader

```go
customUpgrader := websocket.Upgrader{
    ReadBufferSize:  4096,
    WriteBufferSize: 4096,
    CheckOrigin: func(r *http.Request) bool {
        // Custom origin check logic
        origin := r.Header.Get("Origin")
        return origin == "https://allowed-domain.com"
    },
}

// Set custom upgrader
pkg.SetCustomUpgrader(customUpgrader)
```

### Authentication

```go
// Enable token-based authentication
tkws.EnableAuth(func(r *http.Request) bool {
    token := r.URL.Query().Get("token")
    return isValidToken(token) // Your token validation logic
})

// Enable Basic Auth
tkws.EnableAuth(func(r *http.Request) bool {
    username, password, ok := r.BasicAuth()
    if !ok {
        return false
    }
    return username == "admin" && password == "secret"
})
```

### Heartbeat Configuration

```go
// Enable heartbeat with 30-second interval and 60-second timeout
tkws.EnableHeartbeat(30*time.Second, 60*time.Second)

// Disable heartbeat
tkws.DisableHeartbeat()
```

## Performance Considerations

- Server uses goroutines to handle each connection, suitable for high concurrency scenarios
- Error and connection event channels use buffered channels to avoid blocking
- For long idle connections, use the built-in heartbeat mechanism to keep connections active
- Consider using authentication for production environments
- For high-load scenarios, distribute connections across multiple instances

## Contributing

Issues and Pull Requests are welcome to help improve this project!

## License

MIT 
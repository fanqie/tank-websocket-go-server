# Quick Start

This guide will help you get started with Tank WebSocket server.

## Basic Server Setup

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    // Create a new WebSocket manager
    manager := tkws.NewManager()
    
    // Enable heartbeat (optional, enabled by default with 5s interval)
    manager.EnableHeartbeat(5 * time.Second)
    
    // Enable debug logging (optional, enabled by default)
    manager.EnableDebug()
    
    // Start the manager
    go manager.Start()
    
    // Handle WebSocket connections
    http.HandleFunc("/ws", manager.HandleConnection)
    
    // Start HTTP server
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Available Features

### Authentication

```go
// Enable authentication
manager.EnableAuth(func(r *http.Request) bool {
    // Your authentication logic here
    token := r.URL.Query().Get("token")
    return token == "your-auth-token"
})

// Disable authentication
manager.DisableAuth()
```

### Broadcasting Messages

```go
// Broadcast to all clients
manager.BroadcastMessage([]byte("Hello everyone!"), nil)

// Broadcast to specific topic
manager.BroadcastTopicMessage("news", "Breaking news!")
```

### Topic Management

The server supports topic-based message routing. Clients can subscribe to topics and receive messages sent to those topics.

### Connection Management

```go
// Get total connected clients count
count := manager.GetClientCount()

// Get subscribers count for a topic
topicCount := manager.GetTopicSubscriberCount("news")

// Get all active topics
topics := manager.GetAllTopics()

// Close a specific client connection
manager.CloseClient("user_123")
```

### Heartbeat Mechanism

```go
// Enable heartbeat with custom interval
manager.EnableHeartbeat(5 * time.Second)

// Disable heartbeat
manager.DisableHeartbeat()
```

### Debug Logging

```go
// Enable debug logging
manager.EnableDebug()

// Disable debug logging
manager.DisableDebug()
```

## Next Steps

- [Client Connection Guide](./client-connection.md)
- [Heartbeat Mechanism](./heartbeat.md)
- [Topic Subscription](./topic-subscription.md)
- [Authentication](./authentication.md)
- [Debug Logging](./debug-logging.md) 
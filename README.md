# Tank WebSocket Server

[中文文档](README_zh.md)

A lightweight, feature-rich WebSocket server implementation in Go.

## Features

- **Heartbeat Mechanism**: Automatically maintains long-lived connections
- **Topic Subscription**: Support for pub/sub messaging patterns
- **Authentication**: Flexible authentication system
- **Connection Management**: Efficient handling of client connections
- **Debug Logging**: Built-in debug logging system
- **Error Handling**: Comprehensive error reporting system
- **Event Notifications**: Connection and subscription event tracking

## Installation

```bash
go get github.com/fanqie/tank-websocket-go-server
```

## Quick Start

```go
package main

import (
	"log"
	"net/http"
	tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
	// Create WebSocket manager
	manager := tkws.NewManager()

	// Enable heartbeat (5 seconds interval)
	manager.EnableHeartbeat(5 * time.Second)

	// Start the manager
	go manager.Start()

	// Handle WebSocket connections
	http.HandleFunc("/ws", manager.HandleConnection)

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Client Connection

### Using Native WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onmessage = function(event) {
	console.log('Received:', event.data);
};

ws.onclose = function() {
	console.log('Connection closed');
};

// Send message
ws.send('Hello server!');

// Subscribe to topic
ws.send('sub:mytopic');

// Unsubscribe from topic
ws.send('unsub:mytopic');
```

### Using Tank WebSocket Client (Recommended)

We provide a dedicated client library [tank-websocket.js](https://github.com/fanqie/tank-websocket.js) that offers a more convenient way to interact with the server:

```javascript


import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

twsc.onOpen((event) => {
    console.log("Connection opened", event);
    
    // Subscribe to topic with callback
    twsc.subTopic("mytopic", (data) => {
        console.log("Received topic message:", data);
    });
    
    // Send message
    twsc.send("Hello server!");
});

twsc.onError((event) => {
    console.log("Connection error:", event);
});

twsc.onClose((event) => {
    console.log("Connection closed:", event);
});

// Unsubscribe from topic
twsc.unsubTopic("mytopic");

// Unsubscribe from all topics
twsc.destroyTopics();
```

Install the client library:
```bash
npm install tank-websocket.js
# or
yarn add tank-websocket.js
```

## API Reference

### Manager Methods

- `NewManager()`: Creates a new WebSocket manager
- `Start()`: Starts the WebSocket manager
- `EnableHeartbeat(interval time.Duration)`: Enables heartbeat mechanism
- `DisableHeartbeat()`: Disables heartbeat mechanism
- `EnableAuth(authFunc func(r *http.Request) bool)`: Enables authentication
- `DisableAuth()`: Disables authentication
- `EnableDebug()`: Enables debug logging
- `DisableDebug()`: Disables debug logging
- `BroadcastMessage(message []byte, excludeClient *Client)`: Broadcasts message to all clients
- `BroadcastTopicMessage(topic string, data string)`: Broadcasts message to topic subscribers
- `GetClientCount()`: Gets the number of connected clients
- `GetTopicSubscriberCount(topic string)`: Gets the number of subscribers for a topic
- `GetAllTopics()`: Gets all available topics
- `CloseClient(userID string)`: Closes connection to a specific client
- `Shutdown(ctx context.Context)`: Gracefully shuts down the server

## Advanced Configuration

### Custom WebSocket Upgrader

```go
upgrader := websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Implement your CORS logic
	},
}
tkws.SetCustomUpgrader(upgrader)
```

### Authentication

```go
manager.EnableAuth(func(r *http.Request) bool {
	token := r.URL.Query().Get("token")
	return validateToken(token) // Implement your auth logic
})
```

## Error Handling

The server provides an error event channel that you can listen to:

```go
go func() {
	for err := range manager.Errors {
		log.Printf("Error: %v (Code: %d)", err.Message, err.Code)
	}
}()
```

## Connection Events

Monitor connection events:

```go
go func() {
	for event := range manager.ConnEvents {
		log.Printf("Event: %s, User: %s", event.EventType, event.UserID)
	}
}()
```

## License

MIT License - see the [LICENSE](LICENSE) file for details 
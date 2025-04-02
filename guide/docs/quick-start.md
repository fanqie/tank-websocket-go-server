# Quick Start

This guide will help you get started with Tank WebSocket server quickly.

## Basic Server Setup

Create a new file `main.go`:

```go
package main

import (
    "log"
    "net/http"
    "time"
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
    log.Println("WebSocket server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Running the Server

```bash
go run main.go
```

## Client Connection

### Using Native WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    console.log('Connected to WebSocket server');
    
    // Subscribe to a topic
    ws.send('sub:mytopic');
    
    // Send a message
    ws.send('Hello server!');
};

ws.onmessage = function(event) {
    console.log('Received:', event.data);
};

ws.onclose = function() {
    console.log('Disconnected from WebSocket server');
};
```

### Using Tank WebSocket Client (Recommended)

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

twsc.onOpen((event) => {
    console.log("Connection opened", event);
    
    // Subscribe to topic with callback
    twsc.subscribe("mytopic", (data) => {
        console.log("Received topic message:", data);
    });
    
    // Send message
    twsc.send("Hello server!");
});

// Handle errors
twsc.onError((event) => {
    console.log("Connection error:", event);
});

// Handle connection close
twsc.onClose((event) => {
    console.log("Connection closed:", event);
});
```

## Testing the Connection

1. Start the server:
```bash
go run main.go
```

2. Open your browser's developer tools (F12)
3. In the console, paste the client code
4. You should see connection messages and be able to send/receive messages

## Next Steps

- [Client Connection Guide](./client-connection.md)
- [Heartbeat Mechanism](./heartbeat.md)
- [Topic Subscription](./topic-subscription.md) 
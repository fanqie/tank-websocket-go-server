# Debug Logging

The debug logging system in Tank WebSocket helps you monitor and troubleshoot connection issues and system behavior.

## Overview

The debug logging system provides:

- Detailed connection logs
- Message tracking
- Error reporting
- Performance monitoring

## Server-Side Implementation

### Basic Logging Setup

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

    // Enable debug logging
    manager.EnableDebugLogging()

    // Set custom logger
    manager.SetLogger(log.New(os.Stdout, "[WebSocket] ", log.LstdFlags))

    // Enable heartbeat
    manager.EnableHeartbeat(5 * time.Second)

    // Start the manager
    go manager.Start()

    // Handle WebSocket connections
    http.HandleFunc("/ws", manager.HandleConnection)

    // Start HTTP server
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Custom Logging Handler

```go
// Custom logging handler
type CustomLogger struct {
    logger *log.Logger
}

func (l *CustomLogger) Log(level string, message string, args ...interface{}) {
    l.logger.Printf("[%s] %s: %v", level, message, args)
}

// Set custom logger
customLogger := &CustomLogger{
    logger: log.New(os.Stdout, "[Custom] ", log.LstdFlags),
}
manager.SetLogger(customLogger)
```

## Client-Side Implementation

### Using Native WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

// Enable debug logging
ws.debug = true;

ws.onopen = function() {
    console.log('Connection opened');
};

ws.onmessage = function(event) {
    console.log('Message received:', event.data);
};

ws.onerror = function(error) {
    console.error('Connection error:', error);
};
```

### Using Tank WebSocket Client (Recommended)

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    debug: true,
    logLevel: 'debug'
});

// Log events
twsc.onOpen((event) => {
    console.log("Connection opened", event);
});

twsc.onMessage((event) => {
    console.log("Message received:", event.data);
});

twsc.onError((event) => {
    console.log("Connection error:", event);
});
```

## Log Levels

### Available Log Levels

1. **Debug**
   - Detailed information for debugging
   - Connection state changes
   - Message flow

2. **Info**
   - General operational information
   - Connection events
   - System status

3. **Warning**
   - Potential issues
   - Performance concerns
   - Non-critical errors

4. **Error**
   - Critical errors
   - Connection failures
   - System malfunctions

### Setting Log Levels

```go
// Server-side
manager.SetLogLevel("debug") // or "info", "warning", "error"

// Client-side
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    logLevel: 'debug' // or 'info', 'warning', 'error'
});
```

## Logging Features

### Connection Logging

```go
// Log connection events
manager.OnConnect(func(client *tkws.Client) {
    log.Printf("Client connected: %s", client.ID)
})

manager.OnDisconnect(func(client *tkws.Client) {
    log.Printf("Client disconnected: %s", client.ID)
})
```

### Message Logging

```go
// Log message events
manager.OnMessage(func(client *tkws.Client, message []byte) {
    log.Printf("Message from %s: %s", client.ID, string(message))
})
```

### Error Logging

```go
// Log error events
manager.OnError(func(client *tkws.Client, err error) {
    log.Printf("Error for client %s: %v", client.ID, err)
})
```

## Best Practices

1. **Log Management**
   - Implement log rotation
   - Set appropriate log levels
   - Clean up old logs

2. **Performance**
   - Use async logging
   - Implement log buffering
   - Monitor log size

3. **Security**
   - Avoid logging sensitive data
   - Implement log access control
   - Use secure log storage

## Next Steps

- [Client Connection Guide](./client-connection.md)
- [Topic Subscription](./topic-subscription.md)
- [Authentication](./authentication.md) 
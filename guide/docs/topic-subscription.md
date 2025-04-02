# Topic Subscription

The topic subscription system in Tank WebSocket allows you to implement pub/sub messaging patterns efficiently.

## Overview

The topic subscription system provides:

- Topic-based message routing
- Multiple subtopicrs per topic
- Automatic message broadcasting
- Subscription management

## Server-Side Implementation

### Basic Topic Management

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

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

### Publishing Messages

```go
// Publish message to a topic
manager.Publish("mytopic", "Hello subtopicrs!")

// Publish binary data
data := []byte{1, 2, 3, 4}
manager.Publish("mytopic", data)
```

## Client-Side Implementation

### Using Native WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    // SubTopic to topic
    ws.send('sub:mytopic');
};

ws.onmessage = function(event) {
    console.log('Received:', event.data);
};

// Unsubtopic from topic
ws.send('unsub:mytopic');
```

### Using Tank WebSocket Client (Recommended)

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

twsc.onOpen((event) => {
    // SubTopic to topic with callback
    twsc.subTopic("mytopic", (data) => {
        console.log("Received topic message:", data);
    });
});

// Unsubtopic from topic
twsc.unsubTopic("mytopic");
```

## Topic Patterns

### Wildcard Topics

```javascript
// SubTopic to all topics
twsc.subTopic("*", (data) => {
    console.log("Received message from any topic:", data);
});

// SubTopic to topics matching pattern
twsc.subTopic("news.*", (data) => {
    console.log("Received news message:", data);
});
```

### Multiple Topics

```javascript
// SubTopic to multiple topics
twsc.subtopic(["topic1", "topic2"], (data) => {
    console.log("Received message from either topic:", data);
});
```

## Best Practices

1. **Topic Naming**
   - Use descriptive, hierarchical names
   - Follow a consistent naming convention
   - Avoid special characters

2. **Subscription Management**
   - SubTopic only to needed topics
   - Unsubtopic when no longer needed
   - Monitor subscription count

3. **Message Handling**
   - Validate message format
   - Handle different data types
   - Implement error handling

## Performance Considerations

1. **Message Size**
   - Keep messages concise
   - Use appropriate data formats
   - Consider compression for large payloads

2. **Subscription Count**
   - Monitor number of subscriptions
   - Implement subscription limits
   - Clean up unused subscriptions

3. **Broadcasting**
   - Use selective broadcasting
   - Implement message filtering
   - Consider message queuing

## Next Steps

- [Client Connection Guide](./client-connection.md)
- [Heartbeat Mechanism](./heartbeat.md)
- [Debug Logging](./debug-logging.md) 
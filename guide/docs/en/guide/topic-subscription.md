# Topic Subscription

Tank WebSocket server supports pub/sub messaging patterns through topic subscription.

## Basic Usage

### Using Native WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

// Subscribe to topic
ws.send('sub:mytopic');

// Send message
ws.send('Hello server!');

// Unsubscribe from topic
ws.send('unsub:mytopic');
```

### Using Tank WebSocket Client

We provide a dedicated client library [tank-websocket.js](https://github.com/fanqie/tank-websocket.js) that offers a more convenient way to handle topic subscriptions:

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

// Unsubscribe from topic
twsc.unsubTopic("mytopic");

// Unsubscribe from all topics
twsc.destroyTopics();
```

## Server-Side Topic Management

### Broadcasting Messages

```go
// Broadcast message to all clients
manager.BroadcastMessage([]byte("Hello everyone!"), nil)

// Broadcast message to specific topic
manager.BroadcastTopicMessage("mytopic", "Hello topic subscribers!")
```

### Topic Management Methods

```go
// Get number of subscribers for a topic
count := manager.GetTopicSubscriberCount("mytopic")

// Get all available topics
topics := manager.GetAllTopics()
```

### Monitoring Topic Events

The server provides connection events for topic subscriptions:

```go
go func() {
    for event := range manager.ConnEvents {
        switch event.EventType {
        case "subscribe":
            log.Printf("User %s subscribed to topic: %s", event.UserID, event.Topic)
        case "unsubscribe":
            log.Printf("User %s unsubscribed from topic: %s", event.UserID, event.Topic)
        }
    }
}()
```

## Next Steps

- [Client Connection Guide](./client-connection.md)
- [Heartbeat Mechanism](./heartbeat.md)
- [Authentication](./authentication.md)
- [Debug Logging](./debug-logging.md) 
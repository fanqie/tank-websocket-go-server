# Client Connection Guide

This guide explains how to connect to the Tank WebSocket server using different methods.

## Connection Methods

### 1. Using Native WebSocket

The basic way to connect using the browser's native WebSocket API:

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    console.log('Connected to WebSocket server');
};

ws.onmessage = function(event) {
    console.log('Received:', event.data);
};

ws.onclose = function() {
    console.log('Disconnected from WebSocket server');
};

ws.onerror = function(error) {
    console.error('WebSocket error:', error);
};
```

### 2. Using Tank WebSocket Client (Recommended)

Our dedicated client library provides a more convenient way to connect:

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

// Connection events
twsc.onOpen((event) => {
    console.log("Connection opened", event);
});

twsc.onMessage((event) => {
    console.log("Message received:", event.data);
});

twsc.onClose((event) => {
    console.log("Connection closed:", event);
});

twsc.onError((event) => {
    console.log("Connection error:", event);
});
```

## Connection Options

### Authentication

You can add authentication tokens to the connection URL:

```javascript
// Using native WebSocket
const ws = new WebSocket('ws://localhost:8080/ws?token=your-auth-token');

// Using Tank WebSocket Client
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws?token=your-auth-token');
```

### Reconnection

The Tank WebSocket Client automatically handles reconnection:

```javascript
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    reconnect: true,
    reconnectInterval: 1000 // milliseconds
});
```

## Message Types

### Text Messages

```javascript
// Send text message
ws.send('Hello server!');
// or
twsc.send('Hello server!');
```

### Binary Messages

```javascript
// Send binary message
const data = new Uint8Array([1, 2, 3, 4]);
ws.send(data);
// or
twsc.send(data);
```

### Topic Messages

```javascript
// Subscribe to topic
ws.send('sub:mytopic');
// or
twsc.subscribe('mytopic', (data) => {
    console.log('Received topic message:', data);
});

// Unsubscribe from topic
ws.send('unsub:mytopic');
// or
twsc.unsubscribe('mytopic');
```

## Best Practices

1. **Error Handling**
   - Always implement error handlers
   - Log connection errors for debugging
   - Implement retry logic for failed connections

2. **Connection Management**
   - Close connections when they're no longer needed
   - Handle reconnection scenarios
   - Monitor connection state

3. **Message Handling**
   - Validate message format
   - Handle different message types appropriately
   - Implement timeout for message responses

## Next Steps

- [Heartbeat Mechanism](./heartbeat.md)
- [Topic Subscription](./topic-subscription.md)
- [Authentication](./authentication.md) 
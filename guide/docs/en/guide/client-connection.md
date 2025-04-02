# Client Connection Guide

This guide explains how to connect to the Tank WebSocket server and handle various message types.

## Connection URL

The default WebSocket endpoint is `ws://your-server:8080/ws`. You can optionally include a `user_id` parameter for client identification:

```
ws://your-server:8080/ws?user_id=client123
```

If no `user_id` is provided, the server will generate one automatically.

## Message Format

### Topic Messages

When receiving messages from a subscribed topic, the message will be in JSON format:

```json
{
    "topic": "topic_name",
    "data": "message_content"
}
```

### Connection Events

The server sends connection event notifications with the following format:

```json
sub:topic_name
```

### Error Events

Error messages from the server have this format:

```json
{
    "message": "Error description",
    "code": 1001,
    "time": "2024-04-02T10:00:00Z"
}
```

## Using Native WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws?user_id=client123');

ws.onopen = function() {
    console.log('Connected to WebSocket server');
};

ws.onmessage = function(event) {
    const data = JSON.parse(event.data);
    
    // Handle different message types
    if (data.topic) {
        console.log('Topic message:', data);
    } else if (data.event_type) {
        console.log('Connection event:', data);
    } else if (data.code) {
        console.log('Error event:', data);
    }
};

ws.onerror = function(error) {
    console.error('WebSocket error:', error);
};

ws.onclose = function() {
    console.log('Disconnected from WebSocket server');
};
```

## Authentication

If authentication is enabled on the server, you need to include an authentication token:

```javascript
const ws = new WebSocket('ws://localhost:8080/ws?token=your-auth-token');
```

## Error Codes

- 1001: Message serialization failed
- 1002: Connection upgrade failed
- 1007: Authentication failed

## Next Steps

- [Heartbeat Mechanism](./heartbeat.md)
- [Topic Subscription](./topic-subscription.md)
- [Authentication](./authentication.md)
- [Debug Logging](./debug-logging.md) 
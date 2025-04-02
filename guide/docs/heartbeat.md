# Heartbeat Mechanism

The heartbeat mechanism in Tank WebSocket helps maintain long-lived connections by automatically detecting and handling connection issues.

## Overview

The heartbeat mechanism works by:

1. Sending periodic ping messages to clients
2. Expecting pong responses within a timeout period
3. Closing connections that don't respond in time

## Server Configuration

### Enabling Heartbeat

```go
package main

import (
    "time"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

    // Enable heartbeat with 5-second interval and 15-second timeout
    manager.EnableHeartbeat(5 * time.Second)

    // ... rest of your server setup
}
```

### Customizing Heartbeat Settings

```go
// Set custom heartbeat interval and timeout
manager.EnableHeartbeat(10 * time.Second) // 10-second interval
```

## Client Implementation

### Using Native WebSocket

The browser's native WebSocket implementation automatically handles ping/pong messages. You don't need to implement anything special.

### Using Tank WebSocket Client

The Tank WebSocket Client automatically handles heartbeat messages:

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    heartbeat: true,
    heartbeatInterval: 5000, // 5 seconds
    heartbeatTimeout: 15000  // 15 seconds
});

// Handle heartbeat events
twsc.onHeartbeat(() => {
    console.log("Heartbeat received");
});

twsc.onHeartbeatTimeout(() => {
    console.log("Heartbeat timeout");
});
```

## How It Works

1. **Server Side**:
   - Sends ping messages at configured intervals
   - Monitors for pong responses
   - Closes connections that don't respond within timeout

2. **Client Side**:
   - Automatically responds to ping messages with pong
   - Monitors connection health
   - Handles reconnection if needed

## Best Practices

1. **Interval Selection**
   - Choose intervals based on your network conditions
   - Consider server load and client requirements
   - Typical values: 5-30 seconds for interval, 15-60 seconds for timeout

2. **Error Handling**
   - Implement proper error handling for heartbeat failures
   - Log heartbeat events for debugging
   - Consider implementing reconnection logic

3. **Monitoring**
   - Monitor heartbeat success rates
   - Track connection stability
   - Set up alerts for frequent timeouts

## Troubleshooting

### Common Issues

1. **Frequent Timeouts**
   - Check network stability
   - Verify client is properly handling pings
   - Consider increasing timeout duration

2. **Missing Heartbeats**
   - Check server logs for ping failures
   - Verify client is receiving messages
   - Check for network issues

3. **Connection Drops**
   - Monitor server resources
   - Check for client-side issues
   - Review network configuration

## Next Steps

- [Client Connection Guide](./client-connection.md)
- [Topic Subscription](./topic-subscription.md)
- [Debug Logging](./debug-logging.md) 
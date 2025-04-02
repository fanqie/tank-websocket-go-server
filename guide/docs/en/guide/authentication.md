# Authentication

The authentication system in Tank WebSocket provides secure connection handling and user verification.

## Overview

The authentication system supports:

- Token-based authentication
- Custom authentication handlers
- Connection validation
- User identification

## Server-Side Implementation

### Basic Authentication Setup

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

    // Set authentication handler
    manager.SetAuthHandler(func(token string) bool {
        // Implement your authentication logic here
        return validateToken(token)
    })

    // Enable heartbeat
    manager.EnableHeartbeat(5 * time.Second)

    // Start the manager
    go manager.Start()

    // Handle WebSocket connections
    http.HandleFunc("/ws", manager.HandleConnection)

    // Start HTTP server
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Example token validation function
func validateToken(token string) bool {
    // Implement your token validation logic
    return token != ""
}
```

### Custom Authentication Handler

```go
// Custom authentication handler with user information
type UserInfo struct {
    ID    string
    Name  string
    Roles []string
}

func customAuthHandler(token string) (*UserInfo, error) {
    // Validate token and return user information
    if token == "" {
        return nil, errors.New("invalid token")
    }
    
    return &UserInfo{
        ID:    "user123",
        Name:  "John Doe",
        Roles: []string{"user"},
    }, nil
}

// Set custom authentication handler
manager.SetAuthHandler(customAuthHandler)
```

## Client-Side Implementation

### Using Native WebSocket

```javascript
// Connect with authentication token
const token = "your-auth-token";
const ws = new WebSocket(`ws://localhost:8080/ws?token=${token}`);

ws.onopen = function() {
    console.log('Authenticated connection established');
};

ws.onerror = function(error) {
    console.error('Authentication failed:', error);
};
```

### Using Tank WebSocket Client (Recommended)

```javascript
import TankWebSocket from "tank-websocket.js";

const token = "your-auth-token";
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    auth: {
        token: token
    }
});

twsc.onOpen((event) => {
    console.log("Authenticated connection opened", event);
});

twsc.onError((event) => {
    console.log("Authentication error:", event);
});
```

## Authentication Methods

### Token-Based Authentication

```javascript
// Generate and use JWT token
const token = generateJWTToken();
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    auth: {
        token: token
    }
});
```

### Custom Authentication Headers

```javascript
// Use custom authentication headers
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    auth: {
        headers: {
            'Authorization': 'Bearer your-token',
            'X-Custom-Auth': 'custom-value'
        }
    }
});
```

## Best Practices

1. **Token Management**
   - Use secure token generation
   - Implement token expiration
   - Handle token refresh

2. **Security**
   - Use HTTPS for WebSocket connections
   - Implement rate limiting
   - Validate all user input

3. **Error Handling**
   - Handle authentication failures
   - Implement retry logic
   - Log security events

## Security Considerations

1. **Token Storage**
   - Store tokens securely
   - Use appropriate token formats
   - Implement token rotation

2. **Connection Security**
   - Use secure WebSocket (WSS)
   - Implement connection encryption
   - Monitor for suspicious activity

3. **Access Control**
   - Implement role-based access
   - Set up connection limits
   - Monitor user sessions

## Next Steps

- [Client Connection Guide](./client-connection.md)
- [Topic Subscription](./topic-subscription.md)
- [Debug Logging](./debug-logging.md) 
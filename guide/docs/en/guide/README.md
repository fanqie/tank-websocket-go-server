---
home: true
heroImage: /logo.png
actionText: Get Started →
actionLink: /guide/
features:
- title: Lightweight
  details: Minimal setup with small memory footprint and fast performance
- title: Feature-rich
  details: Includes authentication, heartbeat, topic subscription, and more
- title: Go-powered
  details: Built on Go's concurrency model for high performance WebSocket handling
footer: MIT Licensed | Copyright © 2023-present Tank WebSocket
---

# Tank WebSocket

Tank WebSocket is a lightweight, feature-rich WebSocket server implementation in Go. It provides a simple and efficient way to add real-time communication features to your applications.

## Features

- **Simple API**: Easy-to-use API for managing WebSocket connections
- **Heartbeat Mechanism**: Keep connections alive with automatic ping/pong
- **Topic Subscription**: Implement pub/sub messaging patterns
- **Authentication**: Secure your WebSocket connections
- **Debug Logging**: Comprehensive logging for troubleshooting

## Quick Start

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()
    
    // Start the manager
    go manager.Start()
    
    // Handle WebSocket connections
    http.HandleFunc("/ws", manager.HandleConnection)
    
    // Start HTTP server
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

[Get Started →](/guide/) 
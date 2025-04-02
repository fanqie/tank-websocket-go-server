---
home: true
heroImage: /tank-websocket-go-server/images/logo.png
heroText: Tank WebSocket
tagline: A lightweight, feature-rich WebSocket server implementation in Go
actionText: Get Started →
actionLink: /tank-websocket-go-server/en/guide/
features:
- title: Lightweight
  details: Minimal dependencies and small binary size
- title: Feature-rich
  details: Built-in support for authentication, heartbeat, topic subscription, and more
- title: Easy to use
  details: Simple API and comprehensive documentation
- title: Production ready
  details: Battle-tested in production environments
footer: MIT Licensed | Copyright © 2024-present Tank WebSocket
---

## Quick Start

```bash
# Install
go get github.com/fanqie/tank-websocket-go-server

# Use in your code
import "github.com/fanqie/tank-websocket-go-server/pkg"

// Create a new WebSocket manager
manager := pkg.NewManager()

// Start the server
manager.Start(":8080")
```

## Features

- 🔐 Authentication support
- 💓 Heartbeat mechanism
- 📢 Topic subscription
- 🔍 Debug logging
- 🚀 High performance
- 🔒 Secure by default

## Documentation

- [Installation Guide](/tank-websocket-go-server/en/guide/installation)
- [Quick Start Guide](/tank-websocket-go-server/en/guide/quick-start)
- [Client Connection](/tank-websocket-go-server/en/guide/client-connection)
- [Topic Subscription](/tank-websocket-go-server/en/guide/topic-subscription)
- [Heartbeat Mechanism](/tank-websocket-go-server/en/guide/heartbeat)
- [Authentication](/tank-websocket-go-server/en/guide/authentication)
- [Debug Logging](/tank-websocket-go-server/en/guide/debug-logging)

## Contributing

We welcome contributions! Please see our [Contributing Guide](https://github.com/fanqie/tank-websocket-go-server/blob/main/CONTRIBUTING.md) for details.

## License

[MIT](https://github.com/fanqie/tank-websocket-go-server/blob/main/LICENSE) 
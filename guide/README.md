---
home: true
heroImage: /logo.png
heroText: Tank WebSocket
tagline: A lightweight, feature-rich WebSocket server implementation in Go
actionText: Get Started →
actionLink: /guide/
features:
- title: Heartbeat Mechanism
  details: Automatically maintains long-lived connections with configurable intervals and timeouts
- title: Topic Subscription
  details: Support for pub/sub messaging patterns with flexible topic management
- title: Authentication
  details: Built-in authentication system with customizable validation logic
- title: Connection Management
  details: Efficient handling of client connections with comprehensive event tracking
- title: Debug Logging
  details: Built-in debug logging system for easy troubleshooting
- title: Error Handling
  details: Comprehensive error reporting system with detailed error codes and messages
footer: MIT Licensed | Copyright © 2024-present Fanqie
---

## Quick Start

```bash
# Install the package
go get github.com/fanqie/tank-websocket-go-server

# Create a new WebSocket server
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()
    manager.EnableHeartbeat(5 * time.Second)
    go manager.Start()
    http.HandleFunc("/ws", manager.HandleConnection)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Client Library

We also provide a dedicated client library [tank-websocket.js](https://github.com/fanqie/tank-websocket.js) for easier integration:

```bash
npm install tank-websocket.js
# or
yarn add tank-websocket.js
```

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

twsc.onOpen((event) => {
    console.log("Connection opened", event);
    twsc.subTopic("mytopic", (data) => {
        console.log("Received topic message:", data);
    });
});
```

# Tank WebSocket Documentation

This directory contains the documentation for the Tank WebSocket project. The documentation is built using VuePress.

## Development

To start the development server:

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

The documentation will be available at `http://localhost:8080`.

## Building

To build the documentation for production:

```bash
npm run build
```

The built files will be in `.vuepress/dist`.

## Deployment

To deploy to GitHub Pages:

```bash
npm run deploy
```

This will build the documentation and push it to the `gh-pages` branch.

## Directory Structure

```
docs/
├── .vuepress/          # VuePress configuration
│   ├── config/        # Configuration files
│   └── public/        # Static assets
├── guide/             # Guide documentation
├── api/               # API documentation
├── zh/                # Chinese documentation
└── README.md          # This file
```

## Contributing

To contribute to the documentation:

1. Fork the repository
2. Create a new branch for your changes
3. Make your changes
4. Submit a pull request

## License

MIT License - see the [LICENSE](../LICENSE) file for details 
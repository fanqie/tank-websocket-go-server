# Installation

This guide will help you install and set up Tank WebSocket server.

## Prerequisites

- Go 1.16 or later
- Git

## Installation Methods

### Using go get

```bash
go get github.com/fanqie/tank-websocket-go-server
```

### Manual Installation

1. Clone the repository:
```bash
git clone https://github.com/fanqie/tank-websocket-go-server.git
cd tank-websocket-go-server
```

2. Install dependencies:
```bash
go mod download
```

## Project Structure

After installation, you'll have access to the following package structure:

```
tank-websocket-go-server/
├── pkg/
│   ├── server.go    # Main server implementation
│   ├── types.go     # Type definitions
│   └── instance.go  # Instance management
└── examples/        # Example implementations
```

## Importing the Package

In your Go code, you can import the package as follows:

```go
import tkws "github.com/fanqie/tank-websocket-go-server/pkg"
```

## Next Steps

- [Quick Start Guide](./quick-start.md)
- [Client Connection Guide](./client-connection.md) 
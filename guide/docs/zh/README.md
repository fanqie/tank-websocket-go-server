---
home: true
heroImage: /logo.png
actionText: 开始使用 →
actionLink: /zh/guide/
features:
- title: 轻量级
  details: 最小化设置，内存占用小，性能高
- title: 功能丰富
  details: 包含身份验证、心跳机制、主题订阅等多种功能
- title: Go 语言驱动
  details: 基于 Go 的并发模型，提供高性能的 WebSocket 处理能力
footer: MIT 许可 | 版权所有 © 2023-present Tank WebSocket
---

# Tank WebSocket

Tank WebSocket 是一个用 Go 语言实现的轻量级、功能丰富的 WebSocket 服务器。它提供了一种简单高效的方式来为您的应用程序添加实时通信功能。

## 特性

- **简单的 API**：易于使用的 WebSocket 连接管理 API
- **心跳机制**：通过自动 ping/pong 保持连接活跃
- **主题订阅**：实现发布/订阅消息模式
- **身份验证**：保护您的 WebSocket 连接
- **调试日志**：全面的日志记录，方便故障排查

## 快速开始

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()
    
    // 启动管理器
    go manager.Start()
    
    // 处理 WebSocket 连接
    http.HandleFunc("/ws", manager.HandleConnection)
    
    // 启动 HTTP 服务器
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

[开始使用 →](/zh/guide/) 
# 指南

欢迎使用 Tank WebSocket 指南。本指南将帮助您理解并有效使用 Tank WebSocket 服务器。

## 什么是 Tank WebSocket？

Tank WebSocket 是一个用 Go 语言实现的轻量级、功能丰富的 WebSocket 服务器。它提供：

- 自动心跳机制，用于维护长连接
- 基于主题的发布/订阅消息系统
- 内置的身份验证系统
- 全面的连接管理
- 调试日志功能
- 详细的错误报告

## 开始使用

### 安装

```bash
go get github.com/fanqie/tank-websocket-go-server
```

### 基本用法

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    // 创建 WebSocket 管理器
    manager := tkws.NewManager()

    // 启用心跳（5秒间隔）
    manager.EnableHeartbeat(5 * time.Second)

    // 启动管理器
    go manager.Start()

    // 处理 WebSocket 连接
    http.HandleFunc("/ws", manager.HandleConnection)

    // 启动 HTTP 服务器
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 客户端连接

#### 使用原生 WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onmessage = function(event) {
    console.log('收到消息:', event.data);
};

// 订阅主题
ws.send('sub:mytopic');

// 发送消息
ws.send('你好，服务器！');

// 取消订阅主题
ws.send('unsub:mytopic');
```

#### 使用 Tank WebSocket 客户端（推荐）

我们提供了一个专门的客户端库 [tank-websocket.js](https://github.com/fanqie/tank-websocket.js) 以便于集成：

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

twsc.onOpen((event) => {
    console.log("连接已打开", event);
    
    // 订阅主题并设置回调
    twsc.subscribe("mytopic", (data) => {
        console.log("收到主题消息:", data);
    });
    
    // 发送消息
    twsc.send("你好，服务器！");
});

// 处理错误
twsc.onError((event) => {
    console.log("连接错误:", event);
});

// 处理连接关闭
twsc.onClose((event) => {
    console.log("连接已关闭:", event);
});
```

## 下一步

- [安装指南](./installation.md)
- [快速开始指南](./quick-start.md)
- [客户端连接指南](./client-connection.md)
- [心跳机制](./heartbeat.md)
- [身份验证](./authentication.md)
- [调试日志](./debug-logging.md) 
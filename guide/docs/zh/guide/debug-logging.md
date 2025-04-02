# 调试日志

Tank WebSocket 的调试日志系统帮助您监控和排查连接问题及系统行为。

## 概述

调试日志系统提供：

- 详细的连接日志
- 消息跟踪
- 错误报告
- 性能监控

## 服务器端实现

### 基本日志设置

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

    // 启用调试日志
    manager.EnableDebugLogging()

    // 设置自定义日志记录器
    manager.SetLogger(log.New(os.Stdout, "[WebSocket] ", log.LstdFlags))

    // 启用心跳
    manager.EnableHeartbeat(5 * time.Second)

    // 启动管理器
    go manager.Start()

    // 处理 WebSocket 连接
    http.HandleFunc("/ws", manager.HandleConnection)

    // 启动 HTTP 服务器
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 自定义日志处理器

```go
// 自定义日志处理器
type CustomLogger struct {
    logger *log.Logger
}

func (l *CustomLogger) Log(level string, message string, args ...interface{}) {
    l.logger.Printf("[%s] %s: %v", level, message, args)
}

// 设置自定义日志记录器
customLogger := &CustomLogger{
    logger: log.New(os.Stdout, "[自定义] ", log.LstdFlags),
}
manager.SetLogger(customLogger)
```

## 客户端实现

### 使用原生 WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

// 启用调试日志
ws.debug = true;

ws.onopen = function() {
    console.log('连接已打开');
};

ws.onmessage = function(event) {
    console.log('收到消息:', event.data);
};

ws.onerror = function(error) {
    console.error('连接错误:', error);
};
```

### 使用 Tank WebSocket 客户端（推荐）

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    debug: true,
    logLevel: 'debug'
});

// 日志事件
twsc.onOpen((event) => {
    console.log("连接已打开", event);
});

twsc.onMessage((event) => {
    console.log("收到消息:", event.data);
});

twsc.onError((event) => {
    console.log("连接错误:", event);
});
```

## 日志级别

### 可用的日志级别

1. **调试**
   - 详细的调试信息
   - 连接状态变化
   - 消息流程

2. **信息**
   - 一般操作信息
   - 连接事件
   - 系统状态

3. **警告**
   - 潜在问题
   - 性能问题
   - 非关键错误

4. **错误**
   - 严重错误
   - 连接失败
   - 系统故障

### 设置日志级别

```go
// 服务器端
manager.SetLogLevel("debug") // 或 "info", "warning", "error"

// 客户端
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    logLevel: 'debug' // 或 'info', 'warning', 'error'
});
```

## 日志功能

### 连接日志

```go
// 记录连接事件
manager.OnConnect(func(client *tkws.Client) {
    log.Printf("客户端已连接: %s", client.ID)
})

manager.OnDisconnect(func(client *tkws.Client) {
    log.Printf("客户端已断开: %s", client.ID)
})
```

### 消息日志

```go
// 记录消息事件
manager.OnMessage(func(client *tkws.Client, message []byte) {
    log.Printf("来自 %s 的消息: %s", client.ID, string(message))
})
```

### 错误日志

```go
// 记录错误事件
manager.OnError(func(client *tkws.Client, err error) {
    log.Printf("客户端 %s 的错误: %v", client.ID, err)
})
```

## 最佳实践

1. **日志管理**
   - 实现日志轮转
   - 设置适当的日志级别
   - 清理旧日志

2. **性能**
   - 使用异步日志
   - 实现日志缓冲
   - 监控日志大小

3. **安全性**
   - 避免记录敏感数据
   - 实现日志访问控制
   - 使用安全的日志存储

## 下一步

- [客户端连接指南](./client-connection.md)
- [主题订阅](./topic-subscription.md)
- [身份验证](./authentication.md) 
# 快速开始

本指南将帮助您快速上手 Tank WebSocket 服务器。

## 基本服务器设置

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
    
    // 启用心跳（可选，默认已启用，间隔 5 秒）
    manager.EnableHeartbeat(5 * time.Second)
    
    // 启用调试日志（可选，默认已启用）
    manager.EnableDebug()
    
    // 启动管理器
    go manager.Start()
    
    // 处理 WebSocket 连接
    http.HandleFunc("/ws", manager.HandleConnection)
    
    // 启动 HTTP 服务器
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 可用功能

### 身份验证

```go
// 启用身份验证
manager.EnableAuth(func(r *http.Request) bool {
    // 在此处添加您的身份验证逻辑
    token := r.URL.Query().Get("token")
    return token == "your-auth-token"
})

// 禁用身份验证
manager.DisableAuth()
```

### 广播消息

```go
// 向所有客户端广播
manager.BroadcastMessage([]byte("大家好！"), nil)

// 向特定主题广播
manager.BroadcastTopicMessage("news", "重要新闻！")
```

### 主题管理

服务器支持基于主题的消息路由。客户端可以订阅主题并接收发送到这些主题的消息。

### 连接管理

```go
// 获取已连接客户端总数
count := manager.GetClientCount()

// 获取特定主题的订阅者数量
topicCount := manager.GetTopicSubscriberCount("news")

// 获取所有活动主题
topics := manager.GetAllTopics()

// 关闭特定客户端连接
manager.CloseClient("user_123")
```

### 心跳机制

```go
// 启用心跳并设置自定义间隔
manager.EnableHeartbeat(5 * time.Second)

// 禁用心跳
manager.DisableHeartbeat()
```

### 调试日志

```go
// 启用调试日志
manager.EnableDebug()

// 禁用调试日志
manager.DisableDebug()
```

## 下一步

- [客户端连接指南](./client-connection.md)
- [心跳机制](./heartbeat.md)
- [主题订阅](./topic-subscription.md)
- [身份验证](./authentication.md)
- [调试日志](./debug-logging.md) 
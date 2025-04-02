# 主题订阅

Tank WebSocket 的主题订阅系统允许客户端订阅特定主题并接收相关消息。

## 概述

主题订阅系统提供以下功能：

- 基于主题的消息路由
- 每个主题支持多个订阅者
- 自动消息广播
- 订阅状态管理

## 服务器端实现

### 基本主题订阅设置

```go
package main

import (
    "log"
    "net/http"
    "time"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

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

### 向主题发布消息

```go
// 向主题发送文本消息
manager.PublishToTopic("news", []byte("重要新闻更新"))

// 向主题发送二进制数据
binaryData := []byte{0x01, 0x02, 0x03}
manager.PublishToTopic("binary-updates", binaryData)

// 向多个主题发布
manager.PublishToTopics([]string{"news", "updates"}, []byte("多主题更新"))
```

## 客户端实现

### 使用原生 WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

// 订阅主题
ws.onopen = function() {
    // 发送订阅消息
    ws.send(JSON.stringify({
        action: 'subscribe',
        topic: 'news'
    }));
};

// 接收消息
ws.onmessage = function(event) {
    const message = JSON.parse(event.data);
    
    if (message.topic === 'news') {
        console.log('收到新闻更新:', message.data);
    }
};

// 取消订阅
function unsubscribe() {
    ws.send(JSON.stringify({
        action: 'unsubscribe',
        topic: 'news'
    }));
}
```

### 使用 Tank WebSocket 客户端（推荐）

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

// 连接到服务器
twsc.connect();

// 订阅主题
twsc.subscribe('news', (message) => {
    console.log('收到新闻更新:', message);
});

// 订阅另一个主题
twsc.subscribe('updates', (message) => {
    console.log('收到更新:', message);
});

// 取消订阅
twsc.unsubscribe('news');
```

## 主题模式

### 通配符主题

Tank WebSocket 支持使用通配符进行主题匹配：

```javascript
// 订阅以 'news/' 开头的所有主题
twsc.subscribe('news/*', (message, topic) => {
    console.log(`收到来自 ${topic} 的消息:`, message);
});

// 这将接收到 news/sports, news/politics 等主题的消息
```

### 订阅多个主题

```javascript
// 一次订阅多个主题
twsc.subscribeMulti(['news', 'updates', 'alerts'], (message, topic) => {
    console.log(`收到来自 ${topic} 的消息:`, message);
});
```

## 最佳实践

### 主题命名

- 使用层次结构的主题名称，如 `app/users/notifications`
- 保持一致的命名约定
- 分隔相关主题的组

### 订阅管理

- 仅在需要时订阅主题
- 不再需要时取消订阅主题
- 监控活动订阅数量

### 消息处理

- 有效验证和处理接收到的消息
- 实现错误处理和重试机制
- 使用适当的序列化格式（如 JSON）

## 性能考虑

### 消息大小

- 保持消息紧凑
- 考虑使用二进制格式减少数据大小
- 针对大型消息实现分块传输

### 订阅数量

- 监控每个客户端的订阅数量
- 限制过多的订阅请求
- 实现自动清理过期订阅

### 广播策略

- 优化对大量订阅者的广播
- 考虑异步消息处理
- 实现消息队列以处理高流量情况

## 下一步

- [客户端连接指南](./client-connection.md)
- [心跳机制](./heartbeat.md)
- [调试日志](./debug-logging.md)
# 快速开始

本指南将帮助您快速开始使用 Tank WebSocket 服务器。

## 基本服务器设置

创建新文件 `main.go`：

```go
package main

import (
    "log"
    "net/http"
    "time"
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
    log.Println("WebSocket 服务器启动在 :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 运行服务器

```bash
go run main.go
```

## 客户端连接

### 使用原生 WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    console.log('已连接到 WebSocket 服务器');
    
    // 订阅主题
    ws.send('sub:mytopic');
    
    // 发送消息
    ws.send('你好，服务器！');
};

ws.onmessage = function(event) {
    console.log('收到消息:', event.data);
};

ws.onclose = function() {
    console.log('已断开与 WebSocket 服务器的连接');
};
```

### 使用 Tank WebSocket 客户端（推荐）

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

## 测试连接

1. 启动服务器：
```bash
go run main.go
```

2. 打开浏览器的开发者工具（F12）
3. 在控制台中粘贴客户端代码
4. 您应该能看到连接消息，并且能够发送/接收消息

## 下一步

- [客户端连接指南](./client-connection.md)
- [心跳机制](./heartbeat.md)
- [主题订阅](./topic-subscription.md) 
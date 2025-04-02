# Tank WebSocket 服务器

[English Documentation](README.md)

一个用 Go 语言实现的轻量级、功能丰富的 WebSocket 服务器。

## 特性

- **心跳机制**：自动维护长连接
- **主题订阅**：支持发布/订阅消息模式
- **身份验证**：灵活的认证系统
- **连接管理**：高效的客户端连接处理
- **调试日志**：内置的调试日志系统
- **错误处理**：全面的错误报告系统
- **事件通知**：连接和订阅事件跟踪

## 安装

```bash
go get github.com/fanqie/tank-websocket-go-server
```

## 快速开始

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

## 客户端连接

### 使用原生 WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onmessage = function(event) {
	console.log('收到消息:', event.data);
};

ws.onclose = function() {
	console.log('连接已关闭');
};

// 发送消息
ws.send('你好，服务器！');

// 订阅主题
ws.send('sub:mytopic');

// 取消订阅主题
ws.send('unsub:mytopic');
```

### 使用 Tank WebSocket 客户端（推荐）

我们提供了一个专门的客户端库 [tank-websocket.js](https://github.com/fanqie/tank-websocket.js)，它提供了更便捷的方式来与服务器交互：

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

twsc.onOpen((event) => {
    console.log("连接已打开", event);
    
    // 订阅主题并设置回调
    twsc.subTopic("mytopic", (data) => {
        console.log("收到主题消息:", data);
    });
    
    // 发送消息
    twsc.send("你好，服务器！");
});

twsc.onError((event) => {
    console.log("连接错误:", event);
});

twsc.onClose((event) => {
    console.log("连接已关闭:", event);
});

// 取消订阅主题
twsc.unsubTopic("mytopic");

// 取消订阅所有主题
twsc.destroyTopics();
```

安装客户端库：
```bash
npm install tank-websocket.js
# 或
yarn add tank-websocket.js
```

## API 参考

### 管理器方法

- `NewManager()`: 创建新的 WebSocket 管理器
- `Start()`: 启动 WebSocket 管理器
- `EnableHeartbeat(interval time.Duration)`: 启用心跳机制
- `DisableHeartbeat()`: 禁用心跳机制
- `EnableAuth(authFunc func(r *http.Request) bool)`: 启用身份验证
- `DisableAuth()`: 禁用身份验证
- `EnableDebug()`: 启用调试日志
- `DisableDebug()`: 禁用调试日志
- `BroadcastMessage(message []byte, excludeClient *Client)`: 向所有客户端广播消息
- `BroadcastTopicMessage(topic string, data string)`: 向主题订阅者广播消息
- `GetClientCount()`: 获取已连接客户端数量
- `GetTopicSubscriberCount(topic string)`: 获取主题订阅者数量
- `GetAllTopics()`: 获取所有可用主题
- `CloseClient(userID string)`: 关闭特定客户端的连接
- `Shutdown(ctx context.Context)`: 优雅关闭服务器

## 高级配置

### 自定义 WebSocket 升级器

```go
upgrader := websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 实现你的 CORS 逻辑
	},
}
tkws.SetCustomUpgrader(upgrader)
```

### 身份验证

```go
manager.EnableAuth(func(r *http.Request) bool {
	token := r.URL.Query().Get("token")
	return validateToken(token) // 实现你的认证逻辑
})
```

## 错误处理

服务器提供了一个错误事件通道，你可以监听它：

```go
go func() {
	for err := range manager.Errors {
		log.Printf("错误: %v (代码: %d)", err.Message, err.Code)
	}
}()
```

## 连接事件

监控连接事件：

```go
go func() {
	for event := range manager.ConnEvents {
		log.Printf("事件: %s, 用户: %s", event.EventType, event.UserID)
	}
}()
```

## 许可证

MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

# Tank WebSocket Go Server

一个轻量级、高性能的WebSocket服务器，支持单例模式和多例模式，提供主题订阅功能。

## 功能特点

- **双模式支持**：同时支持单例模式和多例模式
- **主题订阅**：支持客户端订阅特定主题，接收定向消息
- **错误处理**：内置错误事件通道，方便监控和处理错误
- **连接监控**：提供连接状态变化事件（连接、断开、订阅、取消订阅）
- **优雅关闭**：支持服务器优雅关闭，确保资源正确释放
- **状态统计**：提供客户端数量、主题订阅数等统计功能


## 安装使用

### 安装

```bash
go get github.com/fanqie/tank-websocket-go-server
```

### 使用示例

#### 单例模式

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
	// 获取单例WebSocket管理器
	tkws := pkg.GetSingleInstance()
	go tkws.Start()

	// 设置HTTP处理程序
	http.HandleFunc("/ws", tkws.HandleConnection)

	// 创建HTTP服务器
	server := &http.Server{
		Addr: ":8080",
	}

	// 设置HTTP服务器引用，用于关闭
	tkws.SetHTTPServer(server)

	// 启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP服务器启动失败: %v", err)
		}
	}()

	log.Println("WebSocket服务器已在 :8080 启动")

	// 广播消息给所有客户端
	tkws.BroadcastMessage([]byte("欢迎使用Tank WebSocket服务器"))

	// 广播消息给特定主题的订阅者
	tkws.BroadcastTopicMessage("news", "这是一条新闻消息")

	// 等待中断信号，优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 创建超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器
	tkws.Shutdown(ctx)
}
```

#### 多例模式

```go
// 创建多个WebSocket服务器实例
tkws1 := pkg.NewInstance()
tkws2 := pkg.NewInstance()

go tkws1.Start()
go tkws2.Start()

// 处理不同路径的WebSocket连接
http.HandleFunc("/ws1", tkws1.HandleConnection)
http.HandleFunc("/ws2", tkws2.HandleConnection)
```

## 客户端连接

### 浏览器客户端示例

```javascript
// 创建WebSocket连接
const socket = new WebSocket('ws://localhost:8080/ws?user_id=user123');

// 连接打开时
socket.onopen = function(e) {
  console.log('WebSocket连接已建立');
  
  // 订阅主题
  socket.send('sub:news');
};

// 接收消息
socket.onmessage = function(event) {
  console.log('收到消息:', event.data);
};

// 关闭连接
socket.onclose = function(event) {
  console.log('WebSocket连接已关闭');
};

// 取消订阅主题
function unsubscribe() {
  socket.send('unsub:news');
}

// 发送消息
function sendMessage(message) {
  socket.send(message);
}
```

## API参考

### 核心组件

- **Manager**: WebSocket连接管理器
- **Client**: 表示单个WebSocket连接
- **Subscription**: 表示主题订阅

### 主要方法

| 方法 | 描述 |
|-----|------|
| `pkg.GetSingleInstance()` | 获取WebSocket管理器的单例实例 |
| `pkg.NewInstance()` | 创建新的WebSocket管理器实例（多例模式） |
| `tkws.Start()` | 启动WebSocket管理器 |
| `tkws.HandleConnection(w, r)` | 处理新的WebSocket连接请求 |
| `tkws.BroadcastMessage(message)` | 广播消息给所有连接的客户端 |
| `tkws.BroadcastTopicMessage(topic, data)` | 广播消息给特定主题的订阅者 |
| `tkws.Shutdown(ctx)` | 优雅关闭WebSocket管理器和HTTP服务器 |
| `tkws.GetClientCount()` | 获取当前连接的客户端数量 |
| `tkws.GetTopicSubscriberCount(topic)` | 获取特定主题的订阅者数量 |
| `tkws.GetAllTopics()` | 获取所有可用的主题 |
| `tkws.IsRunning()` | 检查服务器是否正在运行 |

### 事件监控

```go
// 处理错误事件
go func() {
    for err := range tkws.Errors {
        log.Printf("WebSocket错误: [代码: %d] %s", err.Code, err.Message)
    }
}()

// 处理连接事件
go func() {
    for event := range tkws.ConnEvents {
        switch event.EventType {
        case "connect":
            log.Printf("新连接: 用户ID=%s", event.UserID)
        case "disconnect":
            log.Printf("断开连接: 用户ID=%s", event.UserID)
        case "subscribe":
            log.Printf("订阅主题: 用户ID=%s, 主题=%s", event.UserID, event.Topic)
        case "unsubscribe":
            log.Printf("取消订阅: 用户ID=%s, 主题=%s", event.UserID, event.Topic)
        }
    }
}()
```

## 高级配置

### 自定义WebSocket升级器

```go
customUpgrader := websocket.Upgrader{
    ReadBufferSize:  4096,
    WriteBufferSize: 4096,
    CheckOrigin: func(r *http.Request) bool {
        // 自定义来源检查逻辑
        origin := r.Header.Get("Origin")
        return origin == "https://allowed-domain.com"
    },
}

// 设置自定义升级器
pkg.SetCustomUpgrader(customUpgrader)
```

## 性能考虑

- 服务器使用goroutine处理每个连接，适合高并发场景
- 错误和连接事件通道使用缓冲通道，避免阻塞
- 对长时间空闲的连接，考虑实现心跳机制保持活动状态

## 贡献

欢迎提交问题和Pull Request，一起改进这个项目！

## 许可证

MIT

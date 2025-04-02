# 主题订阅

Tank WebSocket 服务器支持通过主题订阅实现发布/订阅消息模式。

## 基本用法

### 使用原生 WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

// 订阅主题
ws.send('sub:mytopic');

// 发送消息
ws.send('Hello server!');

// 取消订阅主题
ws.send('unsub:mytopic');
```

### 使用 Tank WebSocket 客户端

我们提供了专门的客户端库 [tank-websocket.js](https://github.com/fanqie/tank-websocket.js)，它提供了更便捷的方式来处理主题订阅：

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
    twsc.send("Hello server!");
});

// 取消订阅主题
twsc.unsubTopic("mytopic");

// 取消所有主题订阅
twsc.destroyTopics();
```

## 服务器端主题管理

### 广播消息

```go
// 向所有客户端广播消息
manager.BroadcastMessage([]byte("大家好！"), nil)

// 向特定主题广播消息
manager.BroadcastTopicMessage("mytopic", "向主题订阅者问好！")
```

### 主题管理方法

```go
// 获取主题的订阅者数量
count := manager.GetTopicSubscriberCount("mytopic")

// 获取所有可用主题
topics := manager.GetAllTopics()
```

### 监控主题事件

服务器提供了主题订阅的连接事件：

```go
go func() {
    for event := range manager.ConnEvents {
        switch event.EventType {
        case "subscribe":
            log.Printf("用户 %s 订阅了主题: %s", event.UserID, event.Topic)
        case "unsubscribe":
            log.Printf("用户 %s 取消订阅了主题: %s", event.UserID, event.Topic)
        }
    }
}()
```

## 下一步

- [客户端连接指南](./client-connection.md)
- [心跳机制](./heartbeat.md)
- [身份验证](./authentication.md)
- [调试日志](./debug-logging.md) 
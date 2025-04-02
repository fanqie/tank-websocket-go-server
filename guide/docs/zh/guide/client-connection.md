# 客户端连接指南

本指南说明如何连接到 Tank WebSocket 服务器并处理各种消息类型。

## 连接 URL

默认的 WebSocket 端点是 `ws://your-server:8080/ws`。您可以选择包含 `user_id` 参数来标识客户端：

```
ws://your-server:8080/ws?user_id=client123
```

如果未提供 `user_id`，服务器将自动生成一个。

## 消息格式

### 主题消息

从订阅的主题接收消息时，消息将采用 JSON 格式：

```json
{
    "topic": "topic_name",
    "data": "message_content"
}
```

### 连接事件

服务器发送的连接事件通知采用以下格式：

```json
{
    "event_type": "connect|disconnect|subscribe|unsubscribe",
    "user_id": "client_id",
    "topic": "topic_name",  // 仅用于订阅/取消订阅事件
    "time": "2024-04-02T10:00:00Z"
}
```

### 错误事件

来自服务器的错误消息采用此格式：

```json
{
    "message": "错误描述",
    "code": 1001,
    "time": "2024-04-02T10:00:00Z"
}
```

## 使用原生 WebSocket

```javascript
const ws = new WebSocket('ws://localhost:8080/ws?user_id=client123');

ws.onopen = function() {
    console.log('已连接到 WebSocket 服务器');
};

ws.onmessage = function(event) {
    const data = JSON.parse(event.data);
    
    // 处理不同类型的消息
    if (data.topic) {
        console.log('主题消息:', data);
    } else if (data.event_type) {
        console.log('连接事件:', data);
    } else if (data.code) {
        console.log('错误事件:', data);
    }
};

ws.onerror = function(error) {
    console.error('WebSocket 错误:', error);
};

ws.onclose = function() {
    console.log('已断开与 WebSocket 服务器的连接');
};
```

## 身份验证

如果服务器启用了身份验证，您需要包含身份验证令牌：

```javascript
const ws = new WebSocket('ws://localhost:8080/ws?token=your-auth-token');
```

## 错误代码

- 1001：消息序列化失败
- 1002：连接升级失败
- 1007：身份验证失败

## 下一步

- [心跳机制](./heartbeat.md)
- [主题订阅](./topic-subscription.md)
- [身份验证](./authentication.md)
- [调试日志](./debug-logging.md) 
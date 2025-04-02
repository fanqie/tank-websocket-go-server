# 客户端连接指南

本指南解释了如何使用不同的方法连接到 Tank WebSocket 服务器。

## 连接方法

### 1. 使用原生 WebSocket

使用浏览器原生 WebSocket API 的基本连接方式：

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    console.log('已连接到 WebSocket 服务器');
};

ws.onmessage = function(event) {
    console.log('收到消息:', event.data);
};

ws.onclose = function() {
    console.log('已断开与 WebSocket 服务器的连接');
};

ws.onerror = function(error) {
    console.error('WebSocket 错误:', error);
};
```

### 2. 使用 Tank WebSocket 客户端（推荐）

我们的专用客户端库提供了更便捷的连接方式：

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

// 连接事件
twsc.onOpen((event) => {
    console.log("连接已打开", event);
});

twsc.onMessage((event) => {
    console.log("收到消息:", event.data);
});

twsc.onClose((event) => {
    console.log("连接已关闭:", event);
});

twsc.onError((event) => {
    console.log("连接错误:", event);
});
```

## 连接选项

### 身份验证

您可以在连接 URL 中添加认证令牌：

```javascript
// 使用原生 WebSocket
const ws = new WebSocket('ws://localhost:8080/ws?token=your-auth-token');

// 使用 Tank WebSocket 客户端
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws?token=your-auth-token');
```

### 重连机制

Tank WebSocket 客户端自动处理重连：

```javascript
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    reconnect: true,
    reconnectInterval: 1000 // 毫秒
});
```

## 消息类型

### 文本消息

```javascript
// 发送文本消息
ws.send('你好，服务器！');
// 或
twsc.send('你好，服务器！');
```

### 二进制消息

```javascript
// 发送二进制消息
const data = new Uint8Array([1, 2, 3, 4]);
ws.send(data);
// 或
twsc.send(data);
```

### 主题消息

```javascript
// 订阅主题
ws.send('sub:mytopic');
// 或
twsc.subscribe('mytopic', (data) => {
    console.log('收到主题消息:', data);
});

// 取消订阅主题
ws.send('unsub:mytopic');
// 或
twsc.unsubscribe('mytopic');
```

## 最佳实践

1. **错误处理**
   - 始终实现错误处理器
   - 记录连接错误以便调试
   - 为失败的连接实现重试逻辑

2. **连接管理**
   - 在不再需要时关闭连接
   - 处理重连场景
   - 监控连接状态

3. **消息处理**
   - 验证消息格式
   - 适当处理不同类型的消息
   - 为消息响应实现超时机制

## 下一步

- [心跳机制](./heartbeat.md)
- [主题订阅](./topic-subscription.md)
- [身份验证](./authentication.md) 
# 心跳机制

Tank WebSocket 的心跳机制通过自动检测和处理连接问题来维持长连接。

## 概述

心跳机制的工作原理：

1. 定期向客户端发送 ping 消息
2. 在超时时间内等待 pong 响应
3. 关闭未及时响应的连接

## 服务器配置

### 启用心跳

```go
package main

import (
    "time"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

    // 启用心跳，设置5秒间隔和15秒超时
    manager.EnableHeartbeat(5 * time.Second)

    // ... 其余服务器设置
}
```

### 自定义心跳设置

```go
// 设置自定义心跳间隔和超时时间
manager.EnableHeartbeat(10 * time.Second) // 10秒间隔
```

## 客户端实现

### 使用原生 WebSocket

浏览器的原生 WebSocket 实现会自动处理 ping/pong 消息，您不需要特别实现任何功能。

### 使用 Tank WebSocket 客户端

Tank WebSocket 客户端自动处理心跳消息：

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    heartbeat: true,
    heartbeatInterval: 5000, // 5秒
    heartbeatTimeout: 15000  // 15秒
});

// 处理心跳事件
twsc.onHeartbeat(() => {
    console.log("收到心跳");
});

twsc.onHeartbeatTimeout(() => {
    console.log("心跳超时");
});
```

## 工作原理

1. **服务器端**：
   - 按配置的间隔发送 ping 消息
   - 监控 pong 响应
   - 关闭超时未响应的连接

2. **客户端**：
   - 自动响应 ping 消息
   - 监控连接健康状态
   - 必要时处理重连

## 最佳实践

1. **间隔选择**
   - 根据网络条件选择间隔
   - 考虑服务器负载和客户端需求
   - 典型值：间隔5-30秒，超时15-60秒

2. **错误处理**
   - 实现适当的心跳失败处理
   - 记录心跳事件以便调试
   - 考虑实现重连逻辑

3. **监控**
   - 监控心跳成功率
   - 跟踪连接稳定性
   - 设置频繁超时告警

## 故障排除

### 常见问题

1. **频繁超时**
   - 检查网络稳定性
   - 验证客户端是否正确处理 ping
   - 考虑增加超时时间

2. **心跳丢失**
   - 检查服务器日志中的 ping 失败
   - 验证客户端是否收到消息
   - 检查网络问题

3. **连接断开**
   - 监控服务器资源
   - 检查客户端问题
   - 检查网络配置

## 下一步

- [客户端连接指南](./client-connection.md)
- [主题订阅](./topic-subscription.md)
- [调试日志](./debug-logging.md) 
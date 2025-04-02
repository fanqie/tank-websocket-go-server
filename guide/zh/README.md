---
home: true
heroImage: /logo.png
heroText: Tank WebSocket
tagline: 一个用 Go 语言实现的轻量级、功能丰富的 WebSocket 服务器
actionText: 开始使用 →
actionLink: /zh/guide/
features:
- title: 心跳机制
  details: 自动维护长连接，支持可配置的心跳间隔和超时时间
- title: 主题订阅
  details: 支持发布/订阅消息模式，灵活的主题管理
- title: 身份验证
  details: 内置认证系统，支持自定义验证逻辑
- title: 连接管理
  details: 高效的客户端连接处理，全面的连接事件跟踪
- title: 调试日志
  details: 内置调试日志系统，方便问题排查
- title: 错误处理
  details: 全面的错误报告系统，包含详细的错误代码和消息
footer: MIT 许可证 | Copyright © 2024-present Fanqie
---

## 快速开始

```bash
# 安装包
go get github.com/fanqie/tank-websocket-go-server

# 创建新的 WebSocket 服务器
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()
    manager.EnableHeartbeat(5 * time.Second)
    go manager.Start()
    http.HandleFunc("/ws", manager.HandleConnection)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 客户端库

我们还提供了一个专门的客户端库 [tank-websocket.js](https://github.com/fanqie/tank-websocket.js) 以便于集成：

```bash
npm install tank-websocket.js
# 或
yarn add tank-websocket.js
```

```javascript
import TankWebSocket from "tank-websocket.js";

const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws');

twsc.onOpen((event) => {
    console.log("连接已打开", event);
    twsc.subTopic("mytopic", (data) => {
        console.log("收到主题消息:", data);
    });
});
``` 
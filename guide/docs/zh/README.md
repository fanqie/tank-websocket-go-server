---
home: true
heroImage: /tank-websocket-go-server/images/logo.png
heroText: Tank WebSocket
tagline: 一个用 Go 语言实现的轻量级、功能丰富的 WebSocket 服务器
actionText: 开始使用 →
actionLink: /tank-websocket-go-server/zh/guide/
features:
- title: 轻量级
  details: 最小化依赖，二进制文件体积小
- title: 功能丰富
  details: 内置身份验证、心跳检测、主题订阅等功能
- title: 易于使用
  details: 简单的 API 和全面的文档
- title: 生产就绪
  details: 在生产环境中经过验证
footer: MIT 许可 | 版权所有 © 2024-present Tank WebSocket
---

## 快速开始

```bash
# 安装
go get github.com/fanqie/tank-websocket-go-server

# 在代码中使用
import "github.com/fanqie/tank-websocket-go-server/pkg"

// 创建新的 WebSocket 管理器
manager := pkg.NewManager()

// 启动服务器
manager.Start(":8080")
```

## 特性

- 🔐 身份验证支持
- 💓 心跳机制
- 📢 主题订阅
- 🔍 调试日志
- 🚀 高性能
- 🔒 默认安全

## 文档

- [安装指南](/tank-websocket-go-server/zh/guide/installation)
- [快速开始](/tank-websocket-go-server/zh/guide/quick-start)
- [客户端连接](/tank-websocket-go-server/zh/guide/client-connection)
- [主题订阅](/tank-websocket-go-server/zh/guide/topic-subscription)
- [心跳机制](/tank-websocket-go-server/zh/guide/heartbeat)
- [身份验证](/tank-websocket-go-server/zh/guide/authentication)
- [调试日志](/tank-websocket-go-server/zh/guide/debug-logging)

## 贡献

我们欢迎贡献！请查看我们的[贡献指南](https://github.com/fanqie/tank-websocket-go-server/blob/main/CONTRIBUTING.md)了解详情。

## 许可证

[MIT](https://github.com/fanqie/tank-websocket-go-server/blob/main/LICENSE) 
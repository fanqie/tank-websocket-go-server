# 安装

本指南将帮助您安装和设置 Tank WebSocket 服务器。

## 前置条件

- Go 1.16 或更高版本
- Git

## 安装方法

### 使用 go get

```bash
go get github.com/fanqie/tank-websocket-go-server
```

### 手动安装

1. 克隆仓库：
```bash
git clone https://github.com/fanqie/tank-websocket-go-server.git
cd tank-websocket-go-server
```

2. 安装依赖：
```bash
go mod download
```

## 项目结构

安装后，您将获得以下包结构：

```
tank-websocket-go-server/
├── pkg/
│   ├── server.go    # 主要服务器实现
│   ├── types.go     # 类型定义
│   └── instance.go  # 实例管理
└── examples/        # 示例实现
```

## 导入包

在您的 Go 代码中，可以按以下方式导入包：

```go
import tkws "github.com/fanqie/tank-websocket-go-server/pkg"
```

## 下一步

- [快速开始指南](./quick-start.md)
- [客户端连接指南](./client-connection.md) 
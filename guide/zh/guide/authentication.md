# 身份验证

Tank WebSocket 的身份验证系统提供安全的连接处理和用户验证。

## 概述

身份验证系统支持：

- 基于令牌的身份验证
- 自定义身份验证处理器
- 连接验证
- 用户识别

## 服务器端实现

### 基本身份验证设置

```go
package main

import (
    "log"
    "net/http"
    tkws "github.com/fanqie/tank-websocket-go-server/pkg"
)

func main() {
    manager := tkws.NewManager()

    // 设置身份验证处理器
    manager.SetAuthHandler(func(token string) bool {
        // 在此实现您的身份验证逻辑
        return validateToken(token)
    })

    // 启用心跳
    manager.EnableHeartbeat(5 * time.Second)

    // 启动管理器
    go manager.Start()

    // 处理 WebSocket 连接
    http.HandleFunc("/ws", manager.HandleConnection)

    // 启动 HTTP 服务器
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// 示例令牌验证函数
func validateToken(token string) bool {
    // 实现您的令牌验证逻辑
    return token != ""
}
```

### 自定义身份验证处理器

```go
// 包含用户信息的自定义身份验证处理器
type UserInfo struct {
    ID    string
    Name  string
    Roles []string
}

func customAuthHandler(token string) (*UserInfo, error) {
    // 验证令牌并返回用户信息
    if token == "" {
        return nil, errors.New("无效的令牌")
    }
    
    return &UserInfo{
        ID:    "user123",
        Name:  "张三",
        Roles: []string{"user"},
    }, nil
}

// 设置自定义身份验证处理器
manager.SetAuthHandler(customAuthHandler)
```

## 客户端实现

### 使用原生 WebSocket

```javascript
// 使用身份验证令牌连接
const token = "your-auth-token";
const ws = new WebSocket(`ws://localhost:8080/ws?token=${token}`);

ws.onopen = function() {
    console.log('已建立身份验证连接');
};

ws.onerror = function(error) {
    console.error('身份验证失败:', error);
};
```

### 使用 Tank WebSocket 客户端（推荐）

```javascript
import TankWebSocket from "tank-websocket.js";

const token = "your-auth-token";
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    auth: {
        token: token
    }
});

twsc.onOpen((event) => {
    console.log("身份验证连接已打开", event);
});

twsc.onError((event) => {
    console.log("身份验证错误:", event);
});
```

## 身份验证方法

### 基于令牌的身份验证

```javascript
// 生成并使用 JWT 令牌
const token = generateJWTToken();
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    auth: {
        token: token
    }
});
```

### 自定义身份验证头

```javascript
// 使用自定义身份验证头
const twsc = new TankWebSocket.SocketClient('ws://localhost:8080/ws', {
    auth: {
        headers: {
            'Authorization': 'Bearer your-token',
            'X-Custom-Auth': 'custom-value'
        }
    }
});
```

## 最佳实践

1. **令牌管理**
   - 使用安全的令牌生成
   - 实现令牌过期
   - 处理令牌刷新

2. **安全性**
   - 使用 HTTPS 进行 WebSocket 连接
   - 实现速率限制
   - 验证所有用户输入

3. **错误处理**
   - 处理身份验证失败
   - 实现重试逻辑
   - 记录安全事件

## 安全考虑

1. **令牌存储**
   - 安全存储令牌
   - 使用适当的令牌格式
   - 实现令牌轮换

2. **连接安全**
   - 使用安全的 WebSocket (WSS)
   - 实现连接加密
   - 监控可疑活动

3. **访问控制**
   - 实现基于角色的访问
   - 设置连接限制
   - 监控用户会话

## 下一步

- [客户端连接指南](./client-connection.md)
- [主题订阅](./topic-subscription.md)
- [调试日志](./debug-logging.md) 
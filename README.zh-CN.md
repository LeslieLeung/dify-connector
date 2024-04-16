# dify-connector

[English](./README.md) | 简体中文

[![](https://dcbadge.vercel.app/api/server/WNAMSmTsk8)](https://discord.gg/WNAMSmTsk8)

dify-connector 是一个将 [Dify](https://github.com/langgenius/dify) 发布到各种 IM 平台的工具。

## 特性

- 将 Dify 应用发布到各种 IM 平台(Discord, 钉钉等)
  - ✅Discord
  - ✅钉钉
  - (计划中) Telegram
  - 更多...(欢迎 PR)
- 作为 Dify 的 Go SDK 使用
- (计划中) 管理控制台，用于管理 IM 频道和 Dify 应用
- (计划中) 为 Dify 应用提供内容审查 API

## 部署

### 开始之前

你应该在 Discord 或者钉钉中创建一个机器人，并获取机器人的凭证。

如果你不知道如何获取凭证，互联网和官方文档是你的朋友。

其他前提条件：

- 一个 [Dify](https://github.com/langgenius/dify) 实例(你可以使用官方实例)
- 一个 MySQL 8.0 数据库(你可以使用其他数据库，只要它被 [GORM](https://gorm.io/) 支持)

### Docker Compose(推荐)

你应该在你的服务器上安装 Docker 和 Docker Compose。

```bash
git clone https://github.com/leslieleung/dify-connector.git
docker-compose up -d
```

### Docker

你应该在你的服务器上安装 Docker。并且你应该准备好一个数据库(推荐 MySQL 8.0)。

```bash
docker run -d --name dify-connector -e DATABASE_DSN=<YOUR_DSN> -e BOOTSTRAP_CHANNEL=<YOUR_CHANNEL> leslieleung/dify-connector:latest
```

## 命令

- help: 显示帮助信息
- app: 管理 Dify 应用
  - add: 添加一个 Dify 应用。用法: `app add name type base_url api_key`。
  - list: 列出所有 Dify 应用。
  - remove: 移除一个 Dify 应用。用法: `app remove id`。
  - toggle: 切换一个 Dify 应用。用法: `app toggle id`。
  - use: 使用一个 Dify 应用。用法: `app use id`。

## Dify SDK

### 阻塞模式

```go
package main

import (
  "errors"
  "github.com/google/uuid"
  "github.com/leslieleung/dify-connector/pkg/dify"
  "io"
)

func main() {
    client := dify.New("https://api.dify.ai", "app-xxx")
    client.SetDebug()
    
    text := "Hello, how are you?"
    
    req := dify.CompletionMessageRequest{
        Inputs: map[string]interface{}{
        "query": text,
        },
        User: uuid.New().String(),
    }
    
    resp, err := client.CompletionMessage(req)
    if err != nil {
        panic(err)
    }
    print(resp.Answer)
}
```

### 流式模式

```go
package main

import (
  "errors"
  "github.com/google/uuid"
  "github.com/leslieleung/dify-connector/pkg/dify"
  "io"
)

func main() {
  client := dify.New("https://api.dify.ai", "app-xxx")
  client.SetDebug()

  text := "Hello, how are you?"

  req := dify.CompletionMessageRequest{
    Inputs: map[string]interface{}{
      "query": text,
    },
    User: uuid.New().String(),
  }

  resp, err := client.CompletionMessageStreaming(req)
  if err != nil {
    panic(err)
  }
  for {
    r, err := resp.Recv()
    if err != nil {
      if errors.Is(err, io.EOF) {
        break
      }
    }
    if r.Answer != "" {
      print(r.Answer)
    }
  }
}

```
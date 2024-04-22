# dify-connector

English | [简体中文](./README.zh-CN.md)

[![](https://dcbadge.vercel.app/api/server/WNAMSmTsk8)](https://discord.gg/WNAMSmTsk8)

dify-connector is a tool to publish [Dify](https://github.com/langgenius/dify) apps to various IM platforms.

## Features

- Publish Dify apps to various IM platforms(Discord, DingTalk, etc.)
  - ✅Discord
  - ✅DingTalk
  - (Planned) Telegram
  - more to come...(PRs are welcome)
- Use as a OpenAI compatible API
- Use as a Go SDK for Dify
- (Planned) Management console to manage IM channels and Dify apps
- (Planned) Provide moderation API for Dify apps

## Deployment

### Before you start

You should create a bot in either Discord or DingTalk, and get the bot's credentials.

If you don't know how to obtain the credentials, the Internet and the official documentation are your friends.

Other prerequisites:

- A [Dify](https://github.com/langgenius/dify) Instance(You can use the official instance)
- A MySQL 8.0 Database(You can use other databases, as long as it's supported by [GORM](https://gorm.io/))

### Docker Compose(Recommended)

You should have Docker and Docker Compose installed on your server.

```bash
git clone https://github.com/leslieleung/dify-connector.git
docker-compose up -d
```

### Docker

You should have Docker installed on your server. And you should have a database(MySQL 8.0 is recommended) ready.

```bash
docker run -d --name dify-connector -e DATABASE_DSN=<YOUR_DSN> -e BOOTSTRAP_CHANNEL=<YOUR_CHANNEL> leslieleung/dify-connector:latest
```

## Commands

- help: Display help information
- app: Manage Dify apps
  - add: Add a Dify app. Usage: `app add name type base_url api_key`.
  - list: List all Dify apps.
  - remove: Remove a Dify app. Usage: `app remove id`.
  - toggle: Toggle a Dify app. Usage: `app toggle id`.
  - use: Use a Dify app. Usage: `app use id`.

## Dify SDK

### Blocking mode

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

### Streaming mode

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
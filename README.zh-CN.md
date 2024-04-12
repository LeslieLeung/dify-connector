# dify-connector

[English](./README.md) | 简体中文

[![](https://dcbadge.vercel.app/api/server/WNAMSmTsk8)](https://discord.gg/WNAMSmTsk8)

dify-connector 是一个将 [Dify](https://github.com/langgenius/dify) 发布到各种 IM 平台的工具。

## 特性

- 将 Dify 应用发布到各种 IM 平台(Discord, 钉钉等)
- (计划中) 管理控制台，用于管理 IM 频道和 Dify 应用
- (计划中) 为 Dify 应用提供内容审查 API

## 部署

🚧敬请期待

## 命令

- help: 显示帮助信息
- app: 管理 Dify 应用
  - add: 添加一个 Dify 应用。用法: `app add name type base_url api_key`。
  - list: 列出所有 Dify 应用。
  - remove: 移除一个 Dify 应用。用法: `app remove id`。
  - toggle: 切换一个 Dify 应用。用法: `app toggle id`。
  - use: 使用一个 Dify 应用。用法: `app use id`。
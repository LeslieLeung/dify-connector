# dify-connector

English | [简体中文](./README.zh-CN.md)

[![](https://dcbadge.vercel.app/api/server/WNAMSmTsk8)](https://discord.gg/WNAMSmTsk8)

dify-connector is a tool to publish [Dify](https://github.com/langgenius/dify) apps to various IM platforms.

## Features

- Publish Dify apps to various IM platforms(Discord, DingTalk, etc.)
- (Planned) Management console to manage IM channels and Dify apps
- (Planned) Provide moderation API for Dify apps

## Deployment

🚧TBD

## Commands

- help: Display help information
- app: Manage Dify apps
  - add: Add a Dify app. Usage: `app add name type base_url api_key`.
  - list: List all Dify apps.
  - remove: Remove a Dify app. Usage: `app remove id`.
  - toggle: Toggle a Dify app. Usage: `app toggle id`.
  - use: Use a Dify app. Usage: `app use id`.
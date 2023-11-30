# tgSend

`tgSend` 是一个用于通过 Telegram 发送消息的简单命令行工具。

## 功能

- 通过命令行参数或配置文件发送指定消息到 Telegram 聊天。
- 支持配置文件，方便管理 API 令牌、聊天 ID 等信息。
- 支持代理设置，可通过配置文件指定代理地址。

## 安装

```bash
go get -u github.com/Jisxu/tgsend
```

## 用法

```bash
tgsend -t "你要发送的消息"
```

可用的命令行参数：

- `-t`: 要发送的消息内容，默认为 "hello"。
- `-c`: 配置文件路径，默认为 "config.ini"。

## 配置文件

`config.ini` 文件包含以下配置项：

```ini
[default]
apitoken = YOUR_TELEGRAM_API_TOKEN
chatid = YOUR_TELEGRAM_CHAT_ID
proxy = socks5://127.0.0.1:1080
```

- `apitoken`: Telegram 机器人的 API 令牌。
- `chatid`: 目标聊天的 ID。
- `proxy`: 代理地址，可选配置。默认为 socks5://127.0.0.1:1080。

## 示例

```bash
# 发送消息到默认聊天
tgsend -t "你好，Telegram！"

# 使用自定义配置文件路径
tgsend -t "自定义消息" -c path/to/custom-config.ini
```

## 注意事项

- 请确保 Telegram API 令牌和聊天 ID 的正确性。
- 如果配置文件不存在，会在程序运行时自动生成默认配置文件。

## 许可证

该项目采用 [MIT 许可证](LICENSE)。

请替换文档中的 `YOUR_TELEGRAM_API_TOKEN` 和 `YOUR_TELEGRAM_CHAT_ID` 为你的 Telegram 机器人的 API 令牌和目标聊天的 ID。
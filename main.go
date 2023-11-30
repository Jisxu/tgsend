package main

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	defaultText       = "hello"
	defaultConfigFile = "config.ini"
	defaultProxy      = "socks5://127.0.0.1:1080"
	configKeyApiToken = "apitoken"
	configKeyChatID   = "chatid"
	configKeyProxy    = "proxy"
)

var (
	configSection = ini.DefaultSection
)

func main() {
	// 读取参数
	var text, configFile string
	flag.StringVar(&text, "t", defaultText, "要发送的信息，默认值为hello")
	flag.StringVar(&configFile, "c", defaultConfigFile, "配置文件路径")
	flag.Parse()

	// 读取配置
	cfg, err := loadConfig(configFile)
	if err != nil {
		log.Fatalf("无法加载配置: %v", err)
	}

	// 获取配置值
	apitoken := getConfigValue(cfg, configKeyApiToken, "")
	chatid := getConfigValue(cfg, configKeyChatID, "")
	proxy := getConfigValue(cfg, configKeyProxy, defaultProxy)

	// 初始化机器人
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Fatalf("代理URL解析失败: %v", err)
	}

	bot, err := initializeBot(apitoken, proxyURL)
	if err != nil {
		log.Fatalf("机器人初始化失败: %v", err)
	}

	// 转换聊天ID为int64
	chatID, err := strconv.ParseInt(chatid, 10, 64)
	if err != nil {
		log.Fatalf("转换失败，chatId 为: %s", chatid)
	}

	// 发送消息
	msg := tgbotapi.NewMessage(chatID, text)
	_, err = bot.Send(msg)
	if err != nil {
		log.Fatalf("消息发送失败: %v", err)
	}
}

func loadConfig(configFile string) (*ini.File, error) {
	cfg, err := ini.Load(configFile)
	if err != nil {
		log.Printf("加载配置失败，创建默认配置: %v", err)
		cfg = ini.Empty()
		section := cfg.Section(configSection)
		section.Key(configKeyApiToken).SetValue("")
		section.Key(configKeyChatID).SetValue("")
		section.Key(configKeyProxy).SetValue(defaultProxy)

		if err := cfg.SaveTo(defaultConfigFile); err != nil {
			log.Fatalf("保存默认配置失败: %v", err)
		}

		fmt.Println("未找到配置文件，在当前目录创建新的配置文件 'config.ini'")
		os.Exit(1)
	}
	return cfg, nil
}

func getConfigValue(cfg *ini.File, key, defaultValue string) string {
	section := cfg.Section(configSection)
	value := section.Key(key).String()
	if value == "" {
		return defaultValue
	}
	return value
}

func initializeBot(apiToken string, proxyURL *url.URL) (*tgbotapi.BotAPI, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	bot, err := tgbotapi.NewBotAPIWithClient(apiToken, tgbotapi.APIEndpoint, client)
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	return bot, nil
}

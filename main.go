package main

import (
	"flag"
	"github.com/go-ini/ini"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	//read param
	var text string
	var configFile string
	flag.StringVar(&text, "t", "hello", "要发送的信息，默认值为hello")
	flag.StringVar(&configFile, "c", "config.ini", "配置文件路径")
	flag.Parse()
	// read config
	cfg, err := ini.Load(configFile)
	if err != nil {
		slog.Error("failed:", err)
		cfg := ini.Empty()
		section := cfg.Section(ini.DefaultSection)
		section.Key("apitoken").SetValue("")
		section.Key("chatid").SetValue("")
		section.Key("proxy").SetValue("socks5://127.0.0.1:1080")
		err := cfg.SaveTo("config.ini")
		if err != nil {
			panic(err)
		}
		slog.Info("config not found,create at './config.ini'")
		os.Exit(1)
	}
	section := cfg.Section(ini.DefaultSection)
	apitoken := section.Key("apitoken")
	chatid := section.Key("chatid")
	proxy := section.Key("proxy")
	// init bot
	proxyUrl, err := url.Parse(proxy.String())
	// add proxy
	myClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	bot, err := tgbotapi.NewBotAPIWithClient(apitoken.String(), tgbotapi.APIEndpoint, myClient)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	chatid_Int64, err := strconv.ParseInt(chatid.String(), 10, 64)
	if err != nil {
		slog.Error("convert failed,chatId is:" + chatid.String())
		panic(err)
	}
	msg := tgbotapi.NewMessage(chatid_Int64, text)
	_, err = bot.Send(msg)
	if err != nil {
		panic(err)
	}
}

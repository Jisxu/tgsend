package main

import (
	"flag"
	"github.com/go-ini/ini"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
	"strconv"
)

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	// read config
	section := cfg.Section(ini.DefaultSection)
	apitoken := section.Key("apitoken")
	chatid := section.Key("chatid")
	//read param
	var text string
	flag.StringVar(&text, "t", "hello", "要发送的信息，默认值为hello")
	flag.Parse()
	// init bot
	bot, err := tgbotapi.NewBotAPI(apitoken.String())
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	chatid_Int64, err := strconv.ParseInt(chatid.String(), 10, 64)
	if err != nil {
		slog.Error("convert faild,chatid is:" + chatid.String())
		panic(err)
	}
	msg := tgbotapi.NewMessage(chatid_Int64, text)
	_, err = bot.Send(msg)
	if err != nil {
		panic(err)
	}
}

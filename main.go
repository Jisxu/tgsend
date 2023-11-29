package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	apitoken := cfg.Section(ini.DefaultSection).Key("apitoken")
	chatid := cfg.Section(ini.DefaultSection).Key("chatid")
	fmt.Println(apitoken.String())
	fmt.Println(chatid.String())
}

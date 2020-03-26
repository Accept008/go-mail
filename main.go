package main

import (
	"flag"
	"fmt"
	"github.com/go-mail/config"
	"github.com/go-mail/message"

	"log"
)

var (
	tomlFile = flag.String("config", "toml/qq-mail.toml", "config file")
)

func main() {

	flag.Parse()
	// 解析配置文件
	tomlConfig, err := config.UnmarshalConfig(*tomlFile)
	if err != nil {
		log.Println("解析文件失败", err)
		return
	}

	fmt.Printf("使用的配置文件 -> %v \n", tomlConfig.GetActive())
	fmt.Printf("发件人邮箱 -> %v \n", tomlConfig.GetMailFrom())
	fmt.Printf("收件人邮箱 -> %v \n", tomlConfig.MailTo)

	mail.Send(tomlConfig, "带着天赋空降洛杉矶")
}

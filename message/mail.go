package mail

import (
	"github.com/go-mail/config"
	"gopkg.in/gomail.v2"
	"log"
)

func Send(tomlConfig *config.Config, msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", tomlConfig.MailFrom)
	m.SetHeader("To", tomlConfig.MailTo...)
	m.SetHeader("Subject", "通知邮件")
	m.SetBody("text/html", msg)
	//发送的附件
	//m.Attach("/tmp/sendEmail/FilterLog.2020-03-25.csv")

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewPlainDialer(tomlConfig.GetMailHost(), tomlConfig.GetMailPort(), tomlConfig.GetMailFrom(), tomlConfig.GetMailPassword())
	if err := d.DialAndSend(m); err != nil {
		log.Println("邮件发送失败", err)
		return
	}

	log.Println("邮件发送成功")
}

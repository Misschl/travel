package service

import (
	"gopkg.in/gomail.v2"
)

const (
	emailUser     = "919624032@qq.com"
	emailPassCode = "hhpalbgkzisibedh"
	emailHost     = "smtp.qq.com"
	emailPort     = 25
)

// 发送邮件
func SendMail(to []string, title, body string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", emailUser)
	mail.SetHeader("To", to...)
	mail.SetHeader("Subject", title)
	mail.SetBody("text/html", body)
	return gomail.NewDialer(emailHost, emailPort, emailUser, emailPassCode).DialAndSend(mail)
}

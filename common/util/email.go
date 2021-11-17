package util

import "gopkg.in/gomail.v2"

type EmailInfo struct {
	User     string
	Port     int
	Host     string
	Password string
	To       string
	Subject  string
	Content  string
	Filename string
}

func SendEmail(host, user, password, from, to, content string, port int) error {

	if host == "" {
		return nil
	}
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "主机防护-入侵事件告警")
	m.SetBody("text/plain", content)

	d := gomail.NewDialer(host, port, user, password)
	return d.DialAndSend(m)
}

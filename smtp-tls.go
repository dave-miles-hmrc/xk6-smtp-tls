package smtp_tls

import (
	"strings"

	"go.k6.io/k6/js/modules"
	"gopkg.in/gomail.v2"
)

func init() {
	modules.Register("k6/x/smtp_tls", new(SMTP_TLS))
}

type SMTP_TLS struct{}

type Mail struct {
	From        string   `js:"from"`
	To          []string `js:"to"`
	Cc          []string `js:"cc"`
	Bcc         []string `js:"bcc"`
	Subject     string   `js:"subject"`
	Body        string   `js:"body"`
	Alternative string   `js:"alternative"`
	Attachments []string `js:"attachments"`
}

type SmtpServer struct {
	Host       string `js:"host"`
	Port       int    `js:"port"`
	Account    string `js:"account"`
	Password   string `js:"password"`
	SkipVerify bool   `js:"skipVerify"`
}

func BuildMessage(mail Mail) *gomail.Message {

	m := gomail.NewMessage()

	m.SetHeader("From", mail.From)

	if len(mail.To) > 0 {
		m.SetHeader("To", strings.Join(mail.To, ","))
	}

	if len(mail.Cc) > 0 {
		m.SetHeader("Cc", strings.Join(mail.Cc, ","))
	}

	m.SetHeader("Subject", mail.Subject)

	m.SetBody("text/plain", mail.Body)

	if len(mail.Alternative) > 0 {
		m.AddAlternative("text/html", mail.Alternative)
	}

	if len(mail.Attachments) > 0 {
		for _, file := range mail.Attachments {
			m.Attach(file)
		}
	}

	return m
}

func (*SMTP_TLS) SendMail(smtpServer SmtpServer, mail Mail) {

	email := BuildMessage(mail)

	d := gomail.NewDialer(smtpServer.Host, smtpServer.Port, smtpServer.Account, smtpServer.Password)

	if err := d.DialAndSend(email); err != nil {
		panic(err)
	}

}

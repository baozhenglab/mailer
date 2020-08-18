package mail

import (
	"fmt"
	"net/smtp"
)

type SMTPTransport struct {
	Host     string
	Port     uint64
	Username string
	Password string
	From     string
	FromName string
	Auth     smtp.Auth
	To       []string
	Messages []byte
	Subject  string
}

func (st *SMTPTransport) createAuth() {
	st.Auth = smtp.PlainAuth("", st.Username, st.Password, st.Host)
}

func (st *SMTPTransport) SetFromName(name string) {
	st.FromName = name
}

func (st *SMTPTransport) SetTo(to string) {
	st.To = append(st.To, to)
}

func (st *SMTPTransport) SetSubject(subject string) {
	st.Subject = subject
}

func (st *SMTPTransport) SetMessageHtml(html string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + st.Subject + "\n"
	from := "From: " + st.FromName + " <" + st.From + ">\r\n"
	st.Messages = []byte(from + subject + mime + "\n" + html)
}

func (st *SMTPTransport) SendMail(mailable Mailable) error {
	st.SetSubject(mailable.Subject)
	messages, _ := parseTemplate(mailable.Template, mailable.Data)
	st.SetMessageHtml(messages)
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", st.Host, st.Port),
		st.Auth,
		st.From,
		st.To,
		st.Messages,
	)
	fmt.Println(err)
	return err
}

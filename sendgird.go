package mail

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	endpoint_sendgird = "/v3/mail/send"
	host_sendgird     = "https://api.sendgrid.com"
)

type SendGirdTransport struct {
	KeyApi   string
	From     string
	FromName string
	To       []string
	Messages []byte
	Subject  string
}

func (st *SendGirdTransport) SetTo(to string) {
	st.To = append(st.To, to)
}

func (st *SendGirdTransport) SetSubject(subject string) {
	st.Subject = subject
}

func (st *SendGirdTransport) SetBody(html string) *mail.SGMailV3 {
	m := mail.NewV3Mail()
	from := mail.NewEmail(st.FromName, st.From)
	content := mail.NewContent("text/html", html)
	m.SetFrom(from)
	m.AddContent(content)
	personalization := mail.NewPersonalization()
	for _, to := range st.To {
		personalization.AddTos(mail.NewEmail("User", to))
	}
	personalization.Subject = st.Subject

	// add `personalization` to `m`
	m.AddPersonalizations(personalization)
	return m
}

func (st *SendGirdTransport) SendMail(mailable Mailable) error {
	st.SetSubject(mailable.Subject)
	html, _ := parseTemplate(mailable.Template, mailable.Data)
	message := st.SetBody(html)
	request := sendgrid.GetRequest(st.KeyApi, endpoint_sendgird, host_sendgird)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(message)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return err
}

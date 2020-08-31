package mailer

type MailableContract interface {
	GetTemplate() string
	GetData() map[string]interface{}
	GetSubject() string
}

type TransportContract interface {
	SetTo(to string)
	SendMail(mailable Mailable) error
}

type DriverContract interface {
	SendMail(to string, mail Mailable) error
}

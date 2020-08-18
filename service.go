package mail

import (
	"flag"
	"fmt"

	goservice "github.com/baozhenglab/go-sdk"
)

const (
	KeyMailService = "mail_service"
)

type mailService struct {
	driver   string
	from     string
	fromName string
	//If driver is smtp, we need config host,port,username,password
	host     string
	port     uint64
	username string
	password string
	//If use sendgird, please config key
	keyAPI string
}

type transportDriver struct {
	transport TransportContract
}

func NewMailService() goservice.PrefixConfigure {
	return &mailService{}
}
func (ms *mailService) Name() string {
	return KeyMailService
}

func (ms *mailService) GetPrefix() string {
	return KeyMailService
}

func (ms *mailService) InitFlags() {
	prefix := fmt.Sprintf("%s-", ms.Name())
	flag.StringVar(&ms.driver, prefix+"driver", "", "driver send mail include (smtp , sendgird)")
	flag.StringVar(&ms.from, prefix+"from", "", "Set email send mail")
	flag.StringVar(&ms.fromName, prefix+"from-name", "", "Set Name send mail")
	flag.StringVar(&ms.host, prefix+"host", "", "Host smtp server")
	flag.Uint64Var(&ms.port, prefix+"port", 0, "Port smtp server")
	flag.StringVar(&ms.username, prefix+"username", "", "Username smtp server")
	flag.StringVar(&ms.password, prefix+"password", "", "Password smtp server")
	flag.StringVar(&ms.keyAPI, prefix+"key-api", "", "Key API sendgird if use driver sendgird")
}

func (ms *mailService) Get() interface{} {
	switch ms.driver {
	case "smtp":
		return transportDriver{transport: ms.initSMTP()}
	case "sendgird":
		return transportDriver{transport: ms.initSendgird()}
	}
	panic("Not not mail")
}

func (partTransport transportDriver) SendMail(to string, mail Mailable) error {
	partTransport.transport.SetTo(to)
	return partTransport.transport.SendMail(mail)
}

func (ms *mailService) initSMTP() TransportContract {
	smtpTran := &SMTPTransport{
		Host:     ms.host,
		Port:     ms.port,
		Username: ms.username,
		Password: ms.password,
		From:     ms.from,
		FromName: ms.fromName,
	}
	smtpTran.createAuth()
	return smtpTran
}

func (ms *mailService) initSendgird() TransportContract {
	return &SendGirdTransport{
		KeyApi:   ms.keyAPI,
		From:     ms.from,
		FromName: ms.fromName,
	}

}
func convertToMailable(mail MailableContract) Mailable {
	return Mailable{
		Template: mail.GetTemplate(),
		Subject:  mail.GetSubject(),
		Data:     mail.GetData(),
	}
}

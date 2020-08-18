package mail

import (
	"bytes"
	"strings"
	"text/template"
)

const (
	prefixDirTemplate = "templates/"
	extension         = "html"
	scopePath         = "."
)

type Mailable struct {
	Template string
	Data     map[string]interface{}
	Subject  string
}

func parseTemplate(templateFileName string, data map[string]interface{}) (string, error) {
	if sp := strings.Split(templateFileName, scopePath); len(sp) > 1 {
		templateFileName = strings.Join(sp, "/")
	}
	templateFileName += "." + extension
	t, err := template.ParseFiles(prefixDirTemplate + templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

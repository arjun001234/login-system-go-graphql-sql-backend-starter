package service

import (
	"bytes"
	"log"

	"github.com/arjun001234/E-Commerce-Go-Server/config"
	"github.com/arjun001234/E-Commerce-Go-Server/graph/model"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailTemplatesName string

const (
	Welcome        EmailTemplatesName = "welcome.gohtml"
	ChangePassword EmailTemplatesName = "changePassword.gohtml"
)

type emailParams struct {
	user             model.User
	subject          string
	plainTextContent string
	templateName     EmailTemplatesName
	dataIntoTemplate interface{}
}

type EmailServiceType interface {
	WelcomeEmail(user model.User)
	ChangePasswordEmail(user model.User, token string)
	SendEmail(params emailParams)
}

type email struct {
	from   *mail.Email
	client *sendgrid.Client
	ts     TemplateServiceType
}

func NewEmailService(c *config.Config, ts TemplateServiceType) EmailServiceType {
	from := mail.NewEmail("Arjun Kanojia", c.SenderEmail)
	client := sendgrid.NewSendClient(c.SgApiKey)
	return &email{from, client, ts}
}

func (e *email) SendEmail(params emailParams) {
	fullname := params.user.FirstName + " " + params.user.LastName
	to := mail.NewEmail(fullname, params.user.Email)
	var htmlFile bytes.Buffer
	err := e.ts.GetTemplates().ExecuteTemplate(&htmlFile, string(params.templateName), params.dataIntoTemplate)
	if err != nil {
		log.Printf("email sending error: %v", err)
	}
	message := mail.NewSingleEmail(e.from, params.subject, to, params.plainTextContent, htmlFile.String())
	response, err := e.client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}

func (e *email) WelcomeEmail(user model.User) {
	flname := user.FirstName + " " + user.LastName
	params := emailParams{
		user:             user,
		subject:          "Greetings",
		plainTextContent: "Welcome " + flname,
		templateName:     Welcome,
		dataIntoTemplate: user,
	}
	e.SendEmail(params)
}

func (e *email) ChangePasswordEmail(user model.User, token string) {
	flname := user.FirstName + " " + user.LastName
	params := emailParams{
		user:             user,
		subject:          "Password Change",
		plainTextContent: "Hi," + flname,
		templateName:     ChangePassword,
		dataIntoTemplate: "http://localhost:4000/changePassword?token=" + token,
	}
	e.SendEmail(params)
}

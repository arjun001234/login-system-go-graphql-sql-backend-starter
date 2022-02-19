package service

import "html/template"

type TemplateServiceType interface {
	GetTemplates() *template.Template
}

type templateService struct {
	tpl *template.Template
}

func NewTemplateService() TemplateServiceType {
	tpl := template.Must(template.New("templates").ParseGlob("./templates/*.gohtml"))
	return &templateService{tpl}
}

func (ts *templateService) GetTemplates() *template.Template {
	return ts.tpl
}

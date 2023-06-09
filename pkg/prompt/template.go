package prompt

import (
	"bytes"
	"text/template"
)

type Template struct {
	tmpl *template.Template
}

func MustTemplate(text string) *Template {
	prompt, err := NewTemplate(text)

	if err != nil {
		panic(err)
	}

	return prompt
}

func NewTemplate(text string) (*Template, error) {
	tmpl, err := template.New("prompt").Parse(text)

	if err != nil {
		return nil, err
	}

	return &Template{
		tmpl: tmpl,
	}
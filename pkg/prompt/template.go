package prompt

import (
	"bytes"
	"text/template"
)

type Template struct {
	tmpl *template.Template
}

func MustTemplate(text string) *Template {
	prompt, err := N
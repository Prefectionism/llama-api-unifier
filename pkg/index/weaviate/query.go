package weaviate

import (
	"bytes"
	_ "embed"
	"text/template"
)

var (
	//go:embed query.tmpl
	queryTemplateText string
	queryTemplate     = template.Must(template.New("query"
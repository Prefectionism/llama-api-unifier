package rag

import (
	_ "embed"
)

var (
	//go:embed prompt.tmpl
	promptTemplate string
)

type promptData st
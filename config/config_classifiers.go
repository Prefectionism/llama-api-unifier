package config

import (
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/classifier"
	"github.com/adrianliechti/llama/pkg/classifier/llm"
	"github.com/adrianliechti/llama/pkg/prompt"
	"githu
package openai

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"net/http"

	"github.com/adrianliechti/llama/pkg/provider"

	"github.com/sashabaranov/go-openai"
)

var _ provider.Completer = (*Completer)(nil)

type Completer struct {
	*Config
	client *openai.Client
}

func NewCompleter(options ...Option) (*Completer, error) {
	cfg := &Config{
		model: openai.GPT3Dot5Turbo,
	}

	for _, option := range options {
		option(cfg)
	}

	return &Completer{
		Config: cfg,
		client: cfg.newClient(),
	}, nil
}

f
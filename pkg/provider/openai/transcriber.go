package openai

import (
	"context"

	"github.com/adrianliechti/llama/pkg/provider"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

var _ provider.Transcriber = (*Transcriber)(nil)

type Transcriber struct {
	*Config
	client *openai.Client
}

func NewTranscriber(options ...Option) (*Transcriber, error) {
	cfg := &Config{
		model: openai.Whisper1,
	}

	for _, option := range options {
		option(cfg)
	}

	return &Transcriber{
		Config: cfg,
		client: cfg.newClient(),
	}, nil
}

func (c 
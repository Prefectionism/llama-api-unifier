package openai

import (
	"context"

	"github.com/adrianliechti/llama/pkg/provider"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

var _ provider.Synthesizer = (*Synthesizer)(nil)

type Synthesizer struct {
	*Config
	client *openai.Client
}

func NewSynthesizer(options ...Option) (*Synthesizer, error) {
	cfg := &Config{
		model: string(openai.TTSModel1),
	}

	for _, option := range options {
		option(cfg)
	}

	return &Synthesizer{
		Config: cfg,
		client: cfg.newClient(),
	}, nil
}

f
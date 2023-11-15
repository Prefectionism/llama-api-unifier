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

func (c *Completer) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	req, err := convertCompletionRequest(c.model, messages, options)

	if err != nil {
		return nil, err
	}

	if options.Stream == nil {
		completion, err := c.client.CreateChatCompletion(ctx, *req)

		if err != nil {
			convertError(err)
		}

		choice := completion.Choices[0]

		return &provider.Completion{
			ID:    
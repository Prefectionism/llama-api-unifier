package coqui

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

var (
	_ provider.Synthesizer = (*Synthesizer)(nil)
)

type Synthesizer struct {
	*Config
}

func NewSynthesizer(url string, options ...Option) (*Synthesizer, error) {
	if url == "" {
		return nil, errors.New("invalid url")
	}

	cfg := &Config{
		url: url,

		client: http.DefaultClient,
	}

	for _, option := range options {
		option(cfg)
	}

	return &Synthesizer{
		Config: cfg,
	}, nil
}

func (s 
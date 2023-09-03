package deepl

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

type Translator struct {
	*Config
}

func NewTranslator(url string, options ...Option) (*Translator, error) {
	if url == "" {
		url = "https://api-free.deepl.com"
	}

	cfg := &Config{
		url: url,

		language: "en",

		client: http.DefaultClient,
	}

	for _, option := range options {
		option(cfg)
	}

	return &Translator{
		Config: cfg,
	}, nil
}

func (t *Trans
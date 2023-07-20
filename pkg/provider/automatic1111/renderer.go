package automatic1111

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

var _ provider.Renderer = (*Renderer)(nil)

type Renderer struct {
	*Config
}

func NewRenderer(options ...Option) (*Renderer, error) {
	c := &Config{
		url:    "http://127.0.0.1:7860",
		client: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return &Rende
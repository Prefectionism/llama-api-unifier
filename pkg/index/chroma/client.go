package chroma

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/adrianliechti/llama/pkg/index"

	"github.com/google/uuid"
)

var _ index.Provider = &Client{}

type Client struct {
	url string

	client   *http.Client
	embedder index.Embedder

	namespace string
}

type Option func(*Client)

func New(url, namespace string, options ...Option) (*Client, error) {
	c := &Client{
		url: url,

		client: http.DefaultClient,

		namespace: namespace,
	}

	for _, option := range options {
		option(c)
	}

	if c.embedder == nil {
		return nil, errors.New("embedder is required")
	}

	return c, nil
}

func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

func WithEmbedder(embedder index.Embedder) Option {
	return func(c *Client) {
		c.embedder = embedder
	}
}

func (c *Client) List(ctx context.Context,
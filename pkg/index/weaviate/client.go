
package weaviate

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/adrianliechti/llama/pkg/index"

	"github.com/google/uuid"
)

var _ index.Provider = &Client{}

type Client struct {
	url string

	client   *http.Client
	embedder index.Embedder

	class string
}

type Option func(*Client)

func New(url, namespace string, options ...Option) (*Client, error) {
	c := &Client{
		url: url,

		client: http.DefaultClient,

		class: namespace,
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

func (c *Client) List(ctx context.Context, options *index.ListOptions) ([]index.Document, error) {
	if options == nil {
		options = new(index.ListOptions)
	}

	limit := 50
	cursor := ""

	results := make([]index.Document, 0)

	type pageType struct {
		Objects []Object `json:"objects"`
	}

	for {
		query := url.Values{}
		query.Set("class", c.class)
		query.Set("limit", fmt.Sprintf("%d", limit))

		if cursor != "" {
			query.Set("after", cursor)
		}

		u, _ := url.JoinPath(c.url, "/v1/objects")
		u += "?" + query.Encode()

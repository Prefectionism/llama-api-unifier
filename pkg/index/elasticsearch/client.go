package elasticsearch

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

	client *http.Client

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

	return c, nil
}

func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

func (c *Client) List(ctx context.Context, options *index.ListOptions) ([]index.Document, error) {
	u, _ := url.JoinPath(c.url, "/"+c.namespace+"/_search")

	body := map[string]any{
		"size": 10000,
		"from": 0,
		"query": map[string]any{
			"match_all": map[string]any{},
		},
	}

	req, _ := http.NewRequestWithContext(ctx, "GET", u, jsonReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, convertError(resp)
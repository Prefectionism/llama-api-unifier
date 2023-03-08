
package unstructured

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/adrianliechti/llama/pkg/extractor"
)

var _ extractor.Provider = &Client{}

type Client struct {
	url string

	client *http.Client

	chunkSize    int
	chunkOverlap int
}

type Option func(*Client)

func New(url string, options ...Option) (*Client, error) {
	if url == "" {
		return nil, errors.New("invalid url")
	}

	c := &Client{
		url: url,

		client: http.DefaultClient,

		chunkSize:    4000,
		chunkOverlap: 200,
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

func WithChunkSize(size int) Option {
	return func(c *Client) {
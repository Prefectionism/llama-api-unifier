
package duckduckgo

import (
	"bufio"
	"context"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/adrianliechti/llama/pkg/index"
	"github.com/adrianliechti/llama/pkg/text"
)

var (
	_ index.Provider = (*Client)(nil)
)

type Client struct {
	client *http.Client
}

type Option func(*Client)

func New(options ...Option) (*Client, error) {
	c := &Client{
		client: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return c, nil
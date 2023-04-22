package custom

import (
	"context"
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/index"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	_ index.Provider = (*Client)(nil)
)

type Client struct {
	url string

	client IndexClient
}

type Option func(*Client)

func New(url string, options ...Option) (*Client, error) {
	if url == "" || !strings.HasPrefix(url, "grpc://") {
		return nil, errors.New("invalid url")
	}

	c := &Client{
		url: url,
	}

	for _, option := range options {
		opti
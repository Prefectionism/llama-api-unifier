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
		option(c)
	}

	url = strings.TrimPrefix(c.url, "grpc://")

	conn, err := grpc.Dial(url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	c.client = NewIndexClient(conn)

	return c, nil
}

func (c *Client) List(ctx context.Context, options *index.ListOptions) ([]index.Document, error) {
	return nil, errors.ErrUnsupported
}

func (c *Client) Index(ctx context.Context, documents ...index.Document) error {
	return errors.ErrUnsupported
}

func (c *Client) Delete(ctx context.Context, ids ...string) e
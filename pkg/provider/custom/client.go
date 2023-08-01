
package custom

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/adrianliechti/llama/pkg/to"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	_ provider.Completer = (*Client)(nil)
)

type Client struct {
	url string

	model string

	client CompleterClient
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

	c.client = NewCompleterClient(conn)

	return c, nil
}

func WithModel(model string) Option {
	return func(c *Client) {
		c.model = model
	}
}

func (c *Client) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = &provider.CompleteOptions{}
	}
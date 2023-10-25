package langchain

import (
	"errors"
	"io"
	"net/http"

	"github.com/adrianliechti/llama/pkg/provider"
)

type Client struct {
	provider.Completer
}

func New(url string, options ...Option) (*Client, error) {
	c, err := NewCompleter(url, options...)

	if err != nil {
		return nil, 
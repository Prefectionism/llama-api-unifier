
package huggingface

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	*Embedder
	*Completer
}

func New(url string, options ...Option) (*Client, error) {
	e, err := NewEmbedder(url, options...)

	if err != nil {
		return nil, err
	}

	c, err := NewCompleter(url, options...)

	if err != nil {
		return nil, err
	}

	return &Client{
		Embedder:  e,
		Completer: c,
	}, nil
}

func convertError(resp *http.Response) error {
	data, _ := io.ReadAll(resp.Body)
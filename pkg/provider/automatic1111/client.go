package automatic1111

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	*Renderer
}

func New(options ...Option) (*Client, error) {
	r, err := NewRenderer(options...)

	if err != nil {
		return nil, err
	}

	return &Client{
		Renderer: r,
	}, nil
}

func convertError(resp *http.Response) error {
	data, _ := io.ReadAll(resp.Body)

	if len(data) == 0 {
		return errors.New(http.StatusText(resp.StatusCode))
	}

	return error
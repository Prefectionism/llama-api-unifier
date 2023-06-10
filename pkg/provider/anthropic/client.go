package anthropic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	*Completer
}

func New(options ...Option) (*Client, error) {
	c, err := NewCompleter(options...)

	if err != nil {
		return nil, 
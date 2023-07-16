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
	r, err := NewRenderer(options.
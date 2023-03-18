package chroma

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/adrianliechti/llama/pkg/index"

	"github.com/google/uuid"
)

var _ index.Provider = &Client{}

type Client struct {
	url string

	client   *http.Client
	embedder index.Embedder

	namespace string
}

type Option func(*Client)

func New(url, namespace string, options ...Option) (*Client
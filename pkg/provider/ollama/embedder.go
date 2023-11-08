
package ollama

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
)

var _ provider.Embedder = (*Embedder)(nil)

type Embedder struct {
	*Config
}

func NewEmbedder(url string, options ...Option) (*Embedder, error) {
	if url == "" {
		url = "http://localhost:11434"
	}

	c := &Config{
		url:    url,
		client: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}
package huggingface

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
)

var _ provider.Embedder = (*Embedder)(nil)

type Embedder struct {
	*Config
}

func NewEmbedder(url string, options ...Option) (*Embedder, error) {
	if url 
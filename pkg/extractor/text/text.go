package text

import (
	"context"
	"fmt"
	"io"

	"github.com/adrianliechti/llama/pkg/extractor"
	"github.com/adrianliechti/llama/pkg/text"
)

var _ extractor.Provider = &Provider{}

type Provider struct {
	chunkSize    int
	chunkOverlap int
}

type Option func(*Provider)

func New(options ...Option) (*Provider, error) {
	p := &Provider{
		chunkSize:    4000,
		chunkOverlap: 200,
	}

	for _, option := range options {
		option(p)
	}

	
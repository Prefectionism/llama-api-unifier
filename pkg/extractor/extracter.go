package extractor

import (
	"context"
	"io"
)

type Provider interface {
	Extract(ctx context.Context, input File, options *ExtractOptions) (*Document, error)
}

type ExtractOptions struct {
}

type 
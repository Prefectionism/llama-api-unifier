package provider

import (
	"context"
)

type Embedder interface {
	Embed(ctx context.Context, content s
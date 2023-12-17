package provider

import (
	"context"
)

type Translator interface {
	Translate(ctx context.Context, con
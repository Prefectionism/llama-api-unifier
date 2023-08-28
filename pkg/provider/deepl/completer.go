package deepl

import (
	"context"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

func (t *Translator) Complete(ctx context.Context, messages []provider.Message, 
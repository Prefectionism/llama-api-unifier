
package rag

import (
	"context"
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/chain"
	"github.com/adrianliechti/llama/pkg/classifier"
	"github.com/adrianliechti/llama/pkg/index"
	"github.com/adrianliechti/llama/pkg/prompt"
	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/adrianliechti/llama/pkg/text"
)

var _ chain.Provider = &Chain{}

type Chain struct {
	completer provider.Completer

	template *prompt.Template
	messages []provider.Message

	index index.Provider

	limit       *int
	distance    *float32
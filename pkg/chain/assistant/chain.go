package assistant

import (
	"context"
	"errors"
	"slices"

	"github.com/adrianliechti/llama/pkg/chain"
	"github.com/adrianliechti/llama/pkg/prompt"
	"github.com/adrianliechti/llama/pkg/provider"
)

var _ chain.Provider = &Chain{}

type Chain struct {
	completer provider.Completer

	template *prompt.Template
	messages []provider.Message

	temperature *float32
}

type Option func(*Chain)

func New(options ...Option) (*Chain, error) {
	c := &Chain{}

	for _, option := range options {
		option(c)
	}

	if c.compl
package hermesfn

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/adrianliechti/llama/pkg/adapter"
	"github.com/adrianliechti/llama/pkg/provider"
)

var _ adapter.Provider = &Adapter{}

// https://github.com/NousResearch/Hermes-Function-Calling
type Adapter struct {
	completer provider.Completer
}

func New(completer provider.Completer) (*Adapter, error) {
	a := &Adapter{
		completer: completer,
	}

	return a, nil
}

func (a *Adapter) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	var system string

	if len(messages) > 0 && messages[0].Role == provider.MessageRoleSystem {
		system = messages[0].Content
	}

	prompt, err := convertSystemPrompt(system, options.Functions)

	if err != nil {
		return nil, err
	}

	input := []provider.Message{
		{
			Role:    provider.MessageRoleSystem,
			Content: prompt,
		},
	}

	for _, m := range messages {
		if m.Role == provider.MessageRoleUser {
			input = append(inpu
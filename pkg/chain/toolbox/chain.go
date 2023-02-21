package toolbox

import (
	"context"
	"encoding/json"
	"errors"
	"slices"

	"github.com/adrianliechti/llama/pkg/chain"
	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/adrianliechti/llama/pkg/to"
	"github.com/adrianliechti/llama/pkg/tool"
)

var _ chain.Provider = &Chain{}

type Chain struct {
	completer provider.Completer

	tools map[string]tool.Tool

	temperature *float32
}

type Option func(*Chain)

func New(options ...Option) (*Chain, error) {
	c := &Chain{
		tools: make(map[string]tool.Tool),
	}

	for _, option := range options {
		option(c)
	}

	if c.completer == nil {
		return nil, errors.New("missing completer provider")
	}

	return c, nil
}

func WithCompleter(completer provider.Completer) Option {
	return func(c *Chain) {
		c.completer = completer
	}
}

func WithTools(tool ...tool.Tool) Option {
	return func(c *Chain) {
		for _, t := range tool {
			c.tools[t.Name()] = t
		}
	}
}

func WithTemperature(temperature float32) Option {
	return func(c *Chain) {
		c.temperature = &temperature
	}
}

func (c *Chain) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	if options.Temperature == nil {
		options.Temperature = c.temperature
	}

	functions := make(map[string]provider.Function)

	for _, f := range c.tools {
		functions[f.Name(
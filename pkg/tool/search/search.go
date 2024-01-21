package search

import (
	"context"
	"errors"

	"github.com/adrianliechti/llama/pkg/index"
	"github.com/adrianliechti/llama/pkg/jsonschema"
	"github.com/adrianliechti/llama/pkg/tool"
)

var _ tool.Tool = &Tool{}

type Tool struct {
	name        string
	description string

	index index.Provider
}

func New(index index.Provider) (*Tool, error) {
	t := &Tool{
		name:        "search_tool",
		description: "Get information on recent events from the web.",

		index: index,
	}

	return t, nil
}

type Option func(*Tool)

func WithName(name string) Option {
	return func(t *Tool) {
		t.name = name
	}
}

func WithDescription(description string) Option {
	return func(t *Tool) {
		t.description = description
	}
}

func (t *Tool) Name() string {
	return t.name
}

func (t *Tool) Description() string {
	return t.description
}

func (*Tool) Parameters() jsonschema.Defin
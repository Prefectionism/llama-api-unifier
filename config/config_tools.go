package config

import (
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/index"
	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/adrianliechti/llama/pkg/tool"
	"github.com/adrianliechti/llama/pkg/tool/custom"
	"github.com/adrianliechti/llama/pkg/tool/search"
)

func (c *Config) RegisterTool(id string, val tool.Tool) {
	if c.tools == nil {
		c.tools = make(map[string]tool.Tool)
	}

	c.tools[id] = val
}

type toolContext struct {
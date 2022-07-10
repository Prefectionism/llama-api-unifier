package config

import (
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/index"
	"github.com/adrianliechti/llama/pkg/index/aisearch"
	"github.com/adrianliechti/llama/pkg/index/bing"
	"github.com/adrianliechti/llama/pkg/index/chroma"
	"github.com/adrianliechti/llama/pkg/index/custom"
	"github.com/adrianliechti/llama/pkg/index/duckduckgo"
	"github.com/adrianliechti/llama/pkg/index/elasticsearch"
	"github.com/adrianliechti/llama/pkg/index/memory"
	"github.com/adrianliechti/llama/pkg/index/tavily"
	"github.com/adrianliechti/llama/pkg/index/weaviate"
	"github.com/adrianliechti/llama/pkg/index/wikipedia"
)

func (cfg *Config) RegisterIndex(id string, i index.Provider) {
	if cfg.indexes == nil {
		cfg.indexes = make(map[string]index.Provider)
	}

	cfg.indexes[id] = i
}

type
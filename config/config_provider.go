
package config

import (
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/adrianliechti/llama/pkg/provider/anthropic"
	"github.com/adrianliechti/llama/pkg/provider/automatic1111"
	"github.com/adrianliechti/llama/pkg/provider/azuretranslator"
	"github.com/adrianliechti/llama/pkg/provider/coqui"
	"github.com/adrianliechti/llama/pkg/provider/custom"
	"github.com/adrianliechti/llama/pkg/provider/deepl"
	"github.com/adrianliechti/llama/pkg/provider/groq"
	"github.com/adrianliechti/llama/pkg/provider/huggingface"
	"github.com/adrianliechti/llama/pkg/provider/langchain"
	"github.com/adrianliechti/llama/pkg/provider/llama"
	"github.com/adrianliechti/llama/pkg/provider/mimic"
	"github.com/adrianliechti/llama/pkg/provider/mistral"
	"github.com/adrianliechti/llama/pkg/provider/ollama"
	"github.com/adrianliechti/llama/pkg/provider/openai"
	"github.com/adrianliechti/llama/pkg/provider/whisper"

	"github.com/adrianliechti/llama/pkg/adapter"
	"github.com/adrianliechti/llama/pkg/adapter/hermesfn"
)

func (cfg *Config) RegisterEmbedder(model string, e provider.Embedder) {
	cfg.RegisterModel(model)

	if cfg.embedder == nil {
		cfg.embedder = make(map[string]provider.Embedder)
	}

	cfg.embedder[model] = e
}

func (cfg *Config) RegisterCompleter(model string, c provider.Completer) {
	cfg.RegisterModel(model)

	if cfg.completer == nil {
		cfg.completer = make(map[string]provider.Completer)
	}

	cfg.completer[model] = c
}

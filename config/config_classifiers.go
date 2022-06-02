package config

import (
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/classifier"
	"github.com/adrianliechti/llama/pkg/classifier/llm"
	"github.com/adrianliechti/llama/pkg/prompt"
	"github.com/adrianliechti/llama/pkg/provider"
)

func (cfg *Config) RegisterClassifier(model string, c classifier.Provider) {
	if cfg.classifiers == nil {
		cfg.classifiers = make(map[string]classifier.Provider)
	}

	cfg.classifiers[model]
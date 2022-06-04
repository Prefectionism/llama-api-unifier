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

	cfg.classifiers[model] = c
}

type classifierContext struct {
	Completer provider.Completer

	Template *prompt.Template
	Messages []provider.Message
}

func (cfg *Config) registerClassifiers(f *configFile) error {
	for id, c := range f.Classifiers {
		var err error

		context := classifierContext{}

		if c.Model != "" {
			if context.Completer, err = cfg.Completer(c.Model); err != nil {
				return err
			}
		}

		if c.Template != "" {
			if context.Template, err = parseTemplate(c.Template); err != nil {
				return err
			}
		}

		if c.Me
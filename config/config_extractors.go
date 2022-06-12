package config

import (
	"errors"
	"strings"

	"github.com/adrianliechti/llama/pkg/extractor"
	"github.com/adrianliechti/llama/pkg/extractor/code"
	"github.com/adrianliechti/llama/pkg/extractor/tesseract"
	"github.com/adrianliechti/llama/pkg/extractor/text"
	"github.com/adrianliechti/llama/pkg/extractor/unstructured"
)

func (cfg *Config) RegisterExtractor(model string, e extractor.Provider) {
	if cfg.extractors == nil {
		cfg.extractors = make(map[strin
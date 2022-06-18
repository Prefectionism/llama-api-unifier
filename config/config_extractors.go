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
		cfg.extractors = make(map[string]extractor.Provider)
	}

	cfg.extractors[model] = e
}

func (cfg *Config) registerExtractors(f *configFile) error {
	for id, c := range f.Extractors {
		extractor, err := createExtractor(c)

		if err != nil {
			return err
		}

		cfg.RegisterExtractor(id, extractor)
	}

	return nil
}

func createExtractor(cfg extractorConfig) (extractor.Provider, error) {
	switch strings.ToLower(cfg.Type) {
	case "text":
		return textExtractor(cfg)

	case "code":
		return codeExtractor(cfg)

	case "tesseract":
		return tesseractExtractor(cfg)

	case "unstructured":
		return unstructuredExtractor(cfg)

	default:
		return nil, errors.New("invalid extractor type: " + cfg.Type)
	}
}

func textExtractor(cfg extractorConfig) (extractor.Provider, error) {
	var options []text.Option

	if cfg.ChunkSize != nil {
		options = append(options, text.WithChunkSize(*cfg.ChunkSize))
	}

	if cfg.ChunkOverlap != nil {
		options = append(options, text.WithChunkOverlap(*cfg.ChunkOverlap))
	}

	return text.New(options...)
}

func codeExtractor(
package config

import (
	"errors"
	"os"
	"strings"

	"github.com/adrianliechti/llama/pkg/prompt"
	"github.com/adrianliechti/llama/pkg/provider"
	"gopkg.in/yaml.v3"
)

func parseFile(path string) (*configFile, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	data = []byte(os.ExpandEnv(string(data)))

	var config configFile

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

type configFile struct {
	Authorizers []authorizerConfig `yaml:"authorizers"`

	Providers []providerConfig `yaml:"providers"`

	Indexes     map[string]indexConfig      `yaml:"indexes"`
	Extractors  map[string]extractorConfig  `yaml:"extractors"`
	Classifiers map[string]classifierConfig `yaml:"classifiers"`

	Tools  map[string]toolConfig  `yaml:"tools"`
	Chains map[string]chainConfig `yaml:"chains"`
}

type authorizerConfig struct {
	Type string `yaml:"type"`

	Token string `yaml:"token"`

	Issuer   string `yaml:"issuer"`
	Audience string `yaml:"audience"`
}

type providerConfig struct {
	Type string `yaml:"type"`

	URL   string `yaml:"url"`
	Token string `yaml:"token"`

	Models map[string]modelConfig `yaml:"models"`
}

type ModelType string

const (
	ModelTypeCompleter   ModelType = "completer"
	ModelTypeEmbedder    ModelType = "embedder"
	ModelTypeTranslator  ModelType = "translator"
	ModelTypeTranscriber ModelType = "transcriber"
)

type modelConfig struct {
	ID string `yaml:"id"`

	Type ModelType `yaml:"type"`

	Name        string `yaml:"name"`
	Description string `yaml:"description"`

	Adapter string `yaml:"adapter"`
}

type indexConfig struct {
	Type string `yaml:"type"`

	URL   string `yaml:"url"`
	Token string `yaml:"token"`

	Namespace string `yaml:"namespace"`
	Embedding string `yaml:"embedding"`
}

type extractorConfig struct {
	Type string `yaml:"type"`

	URL   string `yaml:"url"`
	Token string `yaml:"token"`

	ChunkSize    *int `yaml:"chunkSize"`
	ChunkOverlap *int `yaml:"chunkOverlap"`
}

type classifierConfig struct {
	Type string `yaml:"type"`

	Model string `yaml:"model"`

	Template string    `yaml:"template"`
	Messages []message `yaml:"messages"`

	Classes map[string]string `yaml:"classes"`
}

type chainConfig struct {
	Type string `yaml:"type"`

	Model     string `yaml:"model"`
	Index     string `yaml:"index"`
	Embedding string `yaml:"embedd
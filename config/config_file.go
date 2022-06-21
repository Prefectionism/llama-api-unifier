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
	data, err := os.ReadFile(pa
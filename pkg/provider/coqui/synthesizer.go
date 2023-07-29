package coqui

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

var (
	_ provider.Synthesizer = (*Synthesizer)(nil)
)

type Synthesizer struct {
	*Config
}

func NewSynthesizer(url string, options ...Option) (*Synthesi
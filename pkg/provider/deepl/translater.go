package deepl

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

type Translator struct {
	*Config
}

func NewTranslator(url string, options ...Option) (*Translator, error) {
	if url == "" {
		url = "https
package huggingface

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

var _ provider.Completer = (*Completer)(nil)

type Completer struct {
	*Config
}

func NewCompleter(url string, options ...Option) (*Completer, error) {
	if url == "" {
		return nil, errors.New("invalid url")
	}

	url = strings.TrimRight(url,
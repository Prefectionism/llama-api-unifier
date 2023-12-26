package whisper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

var (
	_ provider.Transcriber = (*Transcriber)(nil)
)

type Transcriber struct {
	*Config
}

func NewTranscriber(url string, options ...Option) (*Transcriber, error) {
	if url == "" {
		return nil, errors.New("invalid url")
	}

	cfg := &Config{
		url: url,

		client: http
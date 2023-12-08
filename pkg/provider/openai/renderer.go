package openai

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"

	"github.com/sashabaranov/go-openai"
)

var _ provider.Renderer = (*Renderer)(nil)

type Renderer struct {
	*Config
	client *openai.Client

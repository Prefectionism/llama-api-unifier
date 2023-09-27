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

	url = strings.TrimRight(url, "/")
	url = strings.TrimSuffix(url, "/v1")

	cfg := &Config{
		url: url,

		token: "-",
		model: "tgi",

		client: http.DefaultClient,
	}

	for _, option := range options {
		option(cfg)
	}

	return &Completer{
		Config: cfg,
	}, nil
}

func (c *Completer) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	id := uuid.NewString()

	url, _ := url.JoinPath(c.url, "/v1/chat/completions")
	body, err := convertChatRequest(c.model, messages, options)

	if err != nil {
		return nil, err
	}

	if options.Stream == nil {
		req, _ := http.NewRequestWithContext(ctx, "POST", url, jsonReader(body))
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.client.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, convertError(resp)
		}

		var completion ChatCompletion

		if err := json.NewDecoder(resp.Body).Decode(&completion); err != nil {
			return nil, err
		}

		return &provider.Completion{
			ID:     id,
			Reason: provider.CompletionReasonStop,

			Message: provider.Message{
				Role:    provider.MessageRole(completion.Choices[0].Message.Role),
				Content: completion.Choices[0].Message.Content,
			},
		}, nil
	} else {
		defer close(options.Stream)

		req, _ := http.NewRequestWithContext(ctx, "POST", url, jsonReader(body))
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/x-ndjson")

		resp, err := c.client.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, convertError(resp)
		}

		reader := bufio.NewReader(resp.Body)

		var resultText strings.Builder
		var resultRole provider.MessageRole
		var resultReason provid
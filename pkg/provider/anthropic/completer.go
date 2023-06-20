package anthropic

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"github.com/adrianliechti/llama/pkg/provider"
)

var _ provider.Completer = (*Completer)(nil)

type Completer struct {
	*Config
}

func NewCompleter(options ...Option) (*Completer, error) {
	c := &Config{
		url:    "https://api.anthropic.com",
		client: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return &Completer{
		Config: c,
	}, nil
}

func (c *Completer) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	url, _ := url.JoinPath(c.url, "/v1/messages")
	body, err := convertChatRequest(c.model, messages, options)

	if err != nil {
		return nil, err
	}

	if options.Stream == nil {
		req, _ := http.NewRequestWithContext(ctx, "POST", url, jsonReader(body))
		req.Header.Set("x-api-key", c.token)
		req.Header.Set("anthropic-version", "2023-06-01")
		req.Header.Set("content-type", "application/json")

		resp, err := c.client.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, convertError(resp)
		}

		var response MessagesResponse

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, err
		}

		if response.Role != MessageRoleAssistant || len(response.Content) != 1 {
			return nil, errors.New("invalid complete response")
		}

		role := provider.MessageRoleAssistant
		content := response.Content[0].Text

		return &provider.Completion{
			ID:     response.ID,
			Reason: provider.CompletionReasonStop,

			Message: provider.Message{
				Role:    role,
				Content: content,
			},
		}, nil
	} else {
		defer close(options.Stream)

		req, _ := http.NewRequestWithContext(ctx, "POST", url, jsonReader(body))
		req.Header.Set("x-api-key", c.token)
		req.Header.Set("anthropic-version", "2023-06-01")
		req.Header.Set("content-type", "application/json")

		resp, err := c.client.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp
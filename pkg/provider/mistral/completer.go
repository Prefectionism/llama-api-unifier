
package mistral

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/adrianliechti/llama/pkg/provider"
)

var _ provider.Completer = (*Completer)(nil)

type Completer struct {
	*Config
}

func NewCompleter(options ...Option) (*Completer, error) {
	cfg := &Config{
		url: "https://api.mistral.ai",

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

	url, _ := url.JoinPath(c.url, "/v1/chat/completions")
	body, err := convertCompletionRequest(c.model, messages, options)

	if err != nil {
		return nil, err
	}

	if options.Stream == nil {
		req, _ := http.NewRequestWithContext(ctx, "POST", url, jsonReader(body))
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "text/event-stream")

		resp, err := c.client.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, convertError(resp)
		}

		var completion ChatCompletionResponse

		if err := json.NewDecoder(resp.Body).Decode(&completion); err != nil {
			return nil, err
		}

		choice := completion.Choices[0]

		return &provider.Completion{
			ID:     completion.ID,
			Reason: toCompletionReason(choice.FinishReason),

			Message: provider.Message{
				Role:    toMessageRole(choice.Message.Role),
				Content: choice.Message.Content,
			},
		}, nil
	} else {
		defer close(options.Stream)

		req, _ := http.NewRequestWithContext(ctx, "POST", url, jsonReader(body))
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "text/event-stream")

		resp, err := c.client.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, convertError(resp)
		}

		reader := bufio.NewReader(resp.Body)

		result := &provider.Completion{
			Message: provider.Message{
				Role: provider.MessageRoleAssistant,
			},
		}

		for i := 0; ; i++ {
			data, err := reader.ReadBytes('\n')

			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}

				return nil, err
			}

			data = bytes.TrimSpace(data)

			if bytes.HasPrefix(data, []byte("event:")) {
				continue
			}

			data = bytes.TrimPrefix(data, []byte("data:"))
			data = bytes.TrimSpace(data)

			if len(data) == 0 {
				continue
			}

			if bytes.HasPrefix(data, []byte("[DONE]")) {
				break
			}

			var completion ChatCompletionResponse

			if err := json.Unmarshal([]byte(data), &completion); err != nil {
				return nil, err
			}

			choice := completion.Choices[0]

			result.ID = completion.ID
			result.Reason = toCompletionReason(choice.FinishReason)

			result.Message.Role = toMessageRole(choice.Delta.Role)
			result.Message.Content += choice.Delta.Content

			options.Stream <- provider.Completion{
				ID:     result.ID,
				Reason: result.Reason,

				Message: provider.Message{
					Content: choice.Delta.Content,
				},
			}
		}

		return result, nil
	}
}

func convertCompletionRequest(model string, messages []provider.Message, options *provider.CompleteOptions) (*ChatCompletionRequest, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	req := &ChatCompletionRequest{
		Model: model,

		Stream: options.Stream != nil,
	}

	for _, m := range messages {
		message := Message{
			Role:    convertMessageRole(m.Role),
			Content: m.Content,
		}

		req.Messages = append(req.Messages, message)
	}

	return req, nil
}

type ChatCompletionRequest struct {
	Model string `json:"model"`

	Stream bool `json:"stream"`

	Messages []Message `json:"messages"`
}

type ChatCompletionResponse struct {
	ID    string `json:"id"`
	Model string `json:"model"`

	Choices []ChatCompletionChoice `json:"choices"`
}

type ChatCompletionChoice struct {
	Index int `json:"index"`

	Delta   *Message `json:"delta,omitempty"`
	Message *Message `json:"message,omitempty"`

	FinishReason string `json:"finish_reason"`
}

type MessageRole string

var (
	MessageRoleSystem    MessageRole = "system"
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleTool      MessageRole = "tool"
)

type Message struct {
	Role    MessageRole `json:"role"`
	Content string      `json:"content"`
}

func convertMessageRole(role provider.MessageRole) MessageRole {
	switch role {

	case provider.MessageRoleSystem:
		return MessageRoleSystem

	case provider.MessageRoleUser:
		return MessageRoleUser

	case provider.MessageRoleAssistant:
		return MessageRoleAssistant

	case provider.MessageRoleFunction:
		return MessageRoleTool

	default:
		return ""
	}
}

func toMessageRole(role MessageRole) provider.MessageRole {
	switch role {

	case MessageRoleSystem:
		return provider.MessageRoleSystem

	case MessageRoleUser:
		return provider.MessageRoleUser

	case MessageRoleAssistant:
		return provider.MessageRoleAssistant

	case MessageRoleTool:
		return provider.MessageRoleFunction

	default:
		return ""
	}
}

func toCompletionReason(val string) provider.CompletionReason {
	switch val {
	case "stop":
		return provider.CompletionReasonStop

	case "length", "model_length":
		return provider.CompletionReasonLength

	case "tool_calls":
		return provider.CompletionReasonFunction
	}

	return ""
}
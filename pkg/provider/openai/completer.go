package openai

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"net/http"

	"github.com/adrianliechti/llama/pkg/provider"

	"github.com/sashabaranov/go-openai"
)

var _ provider.Completer = (*Completer)(nil)

type Completer struct {
	*Config
	client *openai.Client
}

func NewCompleter(options ...Option) (*Completer, error) {
	cfg := &Config{
		model: openai.GPT3Dot5Turbo,
	}

	for _, option := range options {
		option(cfg)
	}

	return &Completer{
		Config: cfg,
		client: cfg.newClient(),
	}, nil
}

func (c *Completer) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	req, err := convertCompletionRequest(c.model, messages, options)

	if err != nil {
		return nil, err
	}

	if options.Stream == nil {
		completion, err := c.client.CreateChatCompletion(ctx, *req)

		if err != nil {
			convertError(err)
		}

		choice := completion.Choices[0]

		return &provider.Completion{
			ID:     completion.ID,
			Reason: toCompletionResult(choice.FinishReason),

			Message: provider.Message{
				Role:    toMessageRole(choice.Message.Role),
				Content: choice.Message.Content,

				FunctionCalls: toFunctionCalls(choice.Message.ToolCalls),
			},
		}, nil
	} else {
		defer close(options.Stream)

		stream, err := c.client.CreateChatCompletionStream(ctx, *req)

		if err != nil {
			convertError(err)
		}

		result := provider.Completion{
			Message: provider.Message{
				Role: provider.MessageRoleAssistant,
			},
		}

		for {
			completion, err := stream.Recv()

			if errors.Is(err, io.EOF) {
				break
			}

			if err != nil {
				return nil, err
			}

			choice := completion.Choices[0]

			role := toMessageRole(choice.Delta.Role)

			if role == "" {
				role = provider.MessageRoleAssistant
			}

			result.ID = completion.ID
			result.Reason = toCompletionResult(choice.FinishReason)

			result.Message.Role = role
			result.Message.Content += choice.Delta.Content
			result.Message.FunctionCalls = toFunctionCalls(choice.Delta.ToolCalls)

			options.Stream <- provider.Completion{
				ID:     result.ID,
				Reason: result.Reason,

				Message: provider.Message{
					Role:    role,
					Content: choice.Delta.Content,

					FunctionCalls: toFunctionCalls(choice.Delta.ToolCalls),
				},
			}

			if choice.FinishReason != "" {
				break
			}
		}

		return &result, nil
	}
}

func convertCompletionRequest(model string, messages []provider.Message, options *provider.CompleteOptions) (*openai.ChatCompletionRequest, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	req := &openai.ChatCompletionRequest{
		Model: model,
	}

	if options.Format == provider.CompletionFormatJSON {
		req.ResponseFormat = &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		}
	}

	if model == "gpt-4-vision-preview" || model == "gpt-4-1106-vision-preview" {
		req.MaxTokens = 4096
	}

	if options.Stop != nil {
		req.Stop = options.Stop
	}

	for _, f := range options.Functions {
		tool := openai.Tool{
			Type: openai.ToolTypeFunction,

			Function: &openai.FunctionDefinition{
				Name:       f.Name,
				
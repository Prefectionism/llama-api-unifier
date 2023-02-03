package hermesfn

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/adrianliechti/llama/pkg/adapter"
	"github.com/adrianliechti/llama/pkg/provider"
)

var _ adapter.Provider = &Adapter{}

// https://github.com/NousResearch/Hermes-Function-Calling
type Adapter struct {
	completer provider.Completer
}

func New(completer provider.Completer) (*Adapter, error) {
	a := &Adapter{
		completer: completer,
	}

	return a, nil
}

func (a *Adapter) Complete(ctx context.Context, messages []provider.Message, options *provider.CompleteOptions) (*provider.Completion, error) {
	if options == nil {
		options = new(provider.CompleteOptions)
	}

	var system string

	if len(messages) > 0 && messages[0].Role == provider.MessageRoleSystem {
		system = messages[0].Content
	}

	prompt, err := convertSystemPrompt(system, options.Functions)

	if err != nil {
		return nil, err
	}

	input := []provider.Message{
		{
			Role:    provider.MessageRoleSystem,
			Content: prompt,
		},
	}

	for _, m := range messages {
		if m.Role == provider.MessageRoleUser {
			input = append(input, m)
		}

		if m.Role == provider.MessageRoleAssistant {
			input = append(input, m)
		}

		if m.Role == provider.MessageRoleFunction {
			if m.Function == "" {
				return nil, errors.New("function is required")
			}

			prompt, err := convertToolPrompt(m.Function, m.Content)

			if err != nil {
				return nil, err
			}

			input = append(input, provider.Message{
				Role:    provider.MessageRoleFunction,
				Content: prompt,
			})
		}
	}

	completion, err := a.completer.Complete(ctx, input, &provider.CompleteOptions{
		Stream: options.Stream,

		Stop: options.Stop,

		MaxTokens:   options.MaxTokens,
		Temperature: options.Temperature,

		Format: options.Format,
	})

	if err != nil {
		return nil, err
	}

	if call, err := extractToolCall(completion.Message); err == nil {
		completion = &provider.Completion{
			ID: completion.ID,

			Reason: provider.CompletionReasonFunction,

			Message: provider.Message{
				Role: provider.MessageRoleFunction,

				Function: call.Name,

				FunctionCalls: []provider.FunctionCall{
					*call,
				},
			},
		}
	}

	return completion, nil
}

func convertSystemPrompt(prompt string, functions []provider.Function) (string, error) {
	var result string

	if prompt != "" {
		result += strings.TrimSpace(prompt) + "\n"
	}

	result += "You are a function calling AI model. "
	result += `You are provided with function signatures within <tools></tools> XML tags. `
	result += `You may call one or more functions to assist with the user query. `
	result += `Don't make assumptions about what values to plug into functions. `

	result += `Here are the available tools:\n`
	result += `<tools>\n`

	for _, f := range functions {
		if f.Name == "" {
			return "", errors.New("function name is required")
		}

		if f.Description == "" {
			return "", errors.New("function description is required")
		}

		if len(f.Parameters.Properties) == 0 {
			return "", errors.New("function parameters are required")
		}

		tool := Tool{
			Type: "function",

			Function: &ToolFunction{
				Name:        f.Name,
				Description: f.Description,
				Parameters:  f.Parameters,
			},
		}

		data, err := encodeJSON(tool)

		if err != nil {
			return "", err
		}

		result += data
package provider

import (
	"context"
)

type Completer interface {
	Complete(ctx context.Context, messages []Message, options *CompleteOptions) (*Completion, error)
}

type Message struct {
	Role    MessageRole
	Content string

	Files []File

	Function      string
	FunctionCalls []FunctionCall
}

type MessageRole string

const (
	MessageRoleSystem    MessageRole = "system"
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleFunction  MessageRole = "
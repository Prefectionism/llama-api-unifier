
package oai

import (
	"encoding/json"
	"errors"
)

// https://platform.openai.com/docs/api-reference/models/object
type Model struct {
	Object string `json:"object"` // "model"

	ID      string `json:"id"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

// https://platform.openai.com/docs/api-reference/models
type ModelList struct {
	Object string `json:"object"` // "list"

	Models []Model `json:"data"`
}

// https://platform.openai.com/docs/api-reference/embeddings/create
type EmbeddingsRequest struct {
	Input any    `json:"input"`
	Model string `json:"model"`

	// encoding_format string: float, base64
	// dimensions int
	// user string
}

// https://platform.openai.com/docs/api-reference/embeddings/object
type Embedding struct {
	Object string `json:"object"` // "embedding"

	Index     int       `json:"index"`
	Embedding []float32 `json:"embedding"`
}

// https://platform.openai.com/docs/api-reference/embeddings/create
type EmbeddingList struct {
	Object string `json:"object"` // "list"

	Model string      `json:"model"`
	Data  []Embedding `json:"data"`

	// model string

	// usage {
	//   prompt_tokens int
	//   total_tokens int
	// }
}

type MessageRole string

var (
	MessageRoleSystem    MessageRole = "system"
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleTool      MessageRole = "tool"
)

type ResponseFormat string

var (
	ResponseFormatText ResponseFormat = "text"
	ResponseFormatJSON ResponseFormat = "json_object"
)

// // https://platform.openai.com/docs/api-reference/chat/object
type FinishReason string

var (
	FinishReasonStop   FinishReason = "stop"
	FinishReasonLength FinishReason = "length"

	FinishReasonToolCalls     FinishReason = "tool_calls"
	FinishReasonContentFilter FinishReason = "content_filter"
)

// https://platform.openai.com/docs/api-reference/chat/create
type ChatCompletionRequest struct {
	Model string `json:"model"`

	Messages []ChatCompletionMessage `json:"messages"`

	Stream bool   `json:"stream,omitempty"`
	Stop   any    `json:"stop,omitempty"`
	Tools  []Tool `json:"tools,omitempty"`

	MaxTokens   *int     `json:"max_tokens,omitempty"`
	Temperature *float32 `json:"temperature,omitempty"`

	ResponseFormat *ChatCompletionResponseFormat `json:"response_format,omitempty"`

	// frequency_penalty *float32
	// presence_penalty *float32

	// logit_bias
	// logprobs *bool
	// top_logprobs *int

	// n *int

	// seed *int

	// top_p *float32

	// tool_choice string: none, auto

	// user string
}

// https://platform.openai.com/docs/api-reference/chat/create
type ChatCompletionResponseFormat struct {
	Type ResponseFormat `json:"type"`
}

// https://platform.openai.com/docs/api-reference/chat/object
type ChatCompletion struct {
	Object string `json:"object"` // "chat.completion" | "chat.completion.chunk"

	ID string `json:"id"`

	Model   string `json:"model"`
	Created int64  `json:"created"`

	Choices []ChatCompletionChoice `json:"choices"`

	// system_fingerprint string

	// usage {
	//   completion_tokens int
	//   prompt_tokens int
	//   total_tokens int
	// }
}

// https://platform.openai.com/docs/api-reference/chat/object
type ChatCompletionChoice struct {
	Index int `json:"index"`

	Delta   *ChatCompletionMessage `json:"delta,omitempty"`
	Message *ChatCompletionMessage `json:"message,omitempty"`

	FinishReason *FinishReason `json:"finish_reason"`
}

// https://platform.openai.com/docs/api-reference/chat/object
type ChatCompletionMessage struct {
	Role MessageRole `json:"role,omitempty"`

	Content  string           `json:"content"`
	Contents []MessageContent `json:"-"`

	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}

type MessageContent struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`

	ImageURL *MessageContentURL `json:"image_url,omitempty"`
}

type MessageContentURL struct {
	URL string `json:"url"`
}

func (m *ChatCompletionMessage) MarshalJSON() ([]byte, error) {
	if m.Content != "" && m.Contents != nil {
		return nil, errors.New("cannot have both content and contents")
	}

	if len(m.Contents) > 0 {
		type2 := struct {
			Role MessageRole `json:"role"`

			Content  string           `json:"-"`
			Contents []MessageContent `json:"content,omitempty"`

			ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
			ToolCallID string     `json:"tool_call_id,omitempty"`
		}(*m)

		return json.Marshal(type2)
	} else {
		type1 := struct {
			Role MessageRole `json:"role"`

			Content  string           `json:"content"`
			Contents []MessageContent `json:"-"`

			ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
			ToolCallID string     `json:"tool_call_id,omitempty"`
		}(*m)

		return json.Marshal(type1)
	}
}

func (m *ChatCompletionMessage) UnmarshalJSON(data []byte) error {
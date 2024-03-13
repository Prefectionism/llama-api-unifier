
package oai

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type TestContext struct {
	Context context.Context
	Client  *openai.Client

	Model     string
	Embedding openai.EmbeddingModel
}

func newTestContext() *TestContext {
	config := openai.DefaultConfig("")
	config.BaseURL = "http://localhost:8080/oai/v1"

	client := openai.NewClientWithConfig(config)

	return &TestContext{
		Context: context.Background(),
		Client:  client,

		Model:     openai.GPT3Dot5Turbo,
		Embedding: openai.AdaEmbeddingV2,
	}
}

func TestModels(t *testing.T) {
	c := newTestContext()

	resp, err := c.Client.ListModels(c.Context)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Models)

	for _, model := range resp.Models {
		assert.NotEmpty(t, model.ID)
		assert.NotEmpty(t, model.CreatedAt)
		assert.Equal(t, "model", model.Object)
	}

}

func TestEmbedding(t *testing.T) {
	c := newTestContext()

	resp, err := c.Client.CreateEmbeddings(c.Context, &openai.EmbeddingRequest{
		Model: c.Embedding,
		Input: "The food was delicious and the waiter...",

		EncodingFormat: openai.EmbeddingEncodingFormatFloat,
	})

	assert.NoError(t, err)
	assert.Equal(t, "list", resp.Object)
	assert.NotEmpty(t, resp.Model)
	assert.Len(t, resp.Data, 1)

	if len(resp.Data) == 0 {
		t.FailNow()
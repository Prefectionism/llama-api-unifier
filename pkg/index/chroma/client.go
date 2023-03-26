package chroma

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/adrianliechti/llama/pkg/index"

	"github.com/google/uuid"
)

var _ index.Provider = &Client{}

type Client struct {
	url string

	client   *http.Client
	embedder index.Embedder

	namespace string
}

type Option func(*Client)

func New(url, namespace string, options ...Option) (*Client, error) {
	c := &Client{
		url: url,

		client: http.DefaultClient,

		namespace: namespace,
	}

	for _, option := range options {
		option(c)
	}

	if c.embedder == nil {
		return nil, errors.New("embedder is required")
	}

	return c, nil
}

func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

func WithEmbedder(embedder index.Embedder) Option {
	return func(c *Client) {
		c.embedder = embedder
	}
}

func (c *Client) List(ctx context.Context, options *index.ListOptions) ([]index.Document, error) {
	col, err := c.createCollection(c.namespace)

	if err != nil {
		return nil, err
	}

	u, _ := url.JoinPath(c.url, "/api/v1/collections/"+col.ID+"/get")

	body := map[string]any{}

	resp, err := c.client.Post(u, "application/json", jsonReader(body))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, convertError(resp)
	}

	var result getResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	results := make([]index.Document, 0)

	for i := range result.IDs {
		r := index.Document{
			ID: result.IDs[i],

			Content:  result.Documents[i],
			Metadata: result.Metadatas[i],
		}

		results = append(results, r)
	}

	return results, nil
}

func (c *Client) Index(ctx context.Context, documents ...index.Document) error {
	if len(documents) == 0 {
		return nil
	}

	col, err := c.createCollection(c.namespace)

	if err != nil {
		return err
	}

	u, _ := url.JoinPath(c.url, "/api/v1/collections/"+col.ID+"/upsert")

	body := embeddings{
		IDs: make([]string, len(documents)),

		Embeddings: make([][]float32, len(documents)),

		Documents: make([]string, len(documents)),
		Metadatas: make([]map[string]string, len(documents)),
	}

	for i, d := range documents {
		if d.ID == "" {
			d.ID = uuid.NewString()
		}

		if len(d.Embedding) == 0 && c.embedder != nil {
			embedding, err := c.embedder.Embed(ctx, d.Content)

			if err != nil {
				
package automatic1111

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/adrianliechti/llama/pkg/provider"
	"github.com/google/uuid"
)

var _ provider.Renderer = (*Renderer)(nil)

type Renderer struct {
	*Config
}

func NewRenderer(options ...Option) (*Renderer, error) {
	c := &Config{
		url:    "http://127.0.0.1:7860",
		client: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return &Renderer{
		Config: c,
	}, nil
}

func (r *Renderer) Render(ctx context.Context, input string, options *provider.RenderOptions) (*provider.Image, error) {
	body := Text2ImageRequest{
		Prompt: strings.TrimSpace(input),
		//Steps:  20,
	}

	u, _ := url.JoinPath(r.url, "/sdapi/v1/txt2img")
	resp, err := r.client.Post(u, "application/json", jsonReader(body))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, convertError(resp)
	}


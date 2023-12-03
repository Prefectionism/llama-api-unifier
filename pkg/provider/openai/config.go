package openai

import (
	"net/http"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type Config struct {
	url string

	token string
	model string

	client *http.Client
}

type Option func(*Config)

func WithClient(client *http.Client) Optio
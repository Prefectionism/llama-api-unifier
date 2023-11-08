package ollama

import (
	"net/http"
)

type Config struct {
	url string

	token string
	model string

	client *http.Client
}

type Option func(*Config)

func WithClient(client *http.Client) Option {
	r
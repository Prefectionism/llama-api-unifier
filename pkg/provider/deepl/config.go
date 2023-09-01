package deepl

import (
	"net/http"
)

type Config struct {
	url string

	token    string
	language string

	client *http.Client
}

type Option func(*Config)

func WithClient(client *http.Client) Option {
	return func(c *Config) {
		c.client = client
	}

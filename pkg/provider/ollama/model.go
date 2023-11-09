package ollama

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

func (c *Config) ensureModel() error {
	body := ModelRequest{
		Name: c.model,
	}

	u, _ := url.JoinPath(c.url, "/api/show")

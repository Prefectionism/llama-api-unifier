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
	resp, err := c.client.Post(u, "application/json", jsonReader(body))

	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return c.pullModel()
}

func (c *Config) pullModel() error {
	body := PullRequest{
		Name:   c.model,
		Stream: true,
	}

	u, _ := url.JoinPath(c.url, "/api/pull")
	resp, err := c.client.Post(u, "application/json", jsonReader(body))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return convertError(resp)
	}


package oai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/adrianliechti/llama/pkg/jsonschema"
	"github.com/adrianliechti/llama/pkg/provider"

	"github.com/google/uuid"
)

func (s *Server) handleChatCompletion(w http.ResponseWriter, r *http.Request) {
	var req ChatCompletionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	completer, err := s.Completer(req.Model)

	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	messages, err := toMessages(req.Messages)

	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	functions, err := toFunctions(req.Tools)

	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var stops []string

	switch v := req.Stop.(type) {
	case string:
		stops = []string{v}
	case []string:
		stops = v
	}

	options := &provider.CompleteOptions{
		Stop:      stops,
		Functions: functions,

		MaxTokens:   req.MaxTokens,
		Temperature: req.Temperature,
	}

	if req.ResponseFormat != nil {
		if req.ResponseFormat.Type == ResponseFormatJSON {
			options.Format = provider.CompletionFormatJSON
		}
	}

	if req.Stream {
		w.Header().Set("Content-Type", "text/event-stream")

		done := make(chan error)
		stream := make(chan provider.Completion)

		go func() {
			options.Stream = stream

			completion, err := completer.Complete(r.Context(), messages, options)

			select {
			case <-stream:
				break
			default:
				if completion != nil {
					stream <- *completion
				}

				close(stream)
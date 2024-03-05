package oai

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"path"

	"github.com/adrianliechti/llama/pkg/provider"
)

func (s *Server) handleImageGeneration(w http.ResponseWriter, r *http.Request) {
	var req ImageCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	renderer, err := s.Renderer(req.Model)

	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	options := &provider.RenderOptions{}

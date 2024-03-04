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
	var req ImageCreateReque
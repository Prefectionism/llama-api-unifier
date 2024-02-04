package api

import (
	"net/http"

	"github.com/adrianliechti/llama/pkg/extractor"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleExtract(w http.ResponseWriter, r *http.Request) {
	e, er
package api

import (
	"encoding/json"
	"net/http"

	"github.com/adrianliechti/llama/pkg/index"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleIndexIngest(w http.ResponseWriter, r *http.Request) {
	i, err := s.Index(chi.URLParam(r, "index"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request []Document

	if err := 
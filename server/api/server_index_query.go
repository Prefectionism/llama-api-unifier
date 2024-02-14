package api

import (
	"encoding/json"
	"net/http"

	"github.com/adrianliechti/llama/pkg/index"
	"github.com/adrianliechti/llama/pkg/to"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleIndexQuery(w http.ResponseWriter, r *http.Request) {
	i, err := s.Index(chi.URLParam(r, "index"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var query Query

	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, e
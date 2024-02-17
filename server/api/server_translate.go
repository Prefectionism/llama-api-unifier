package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleTranslate(w http.ResponseWriter, r *http.Request) {
	t, err := s.Translator(chi.URLParam(r, "translator"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request Document

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http
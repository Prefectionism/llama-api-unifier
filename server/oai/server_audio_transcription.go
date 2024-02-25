package oai

import (
	"net/http"

	"github.com/adrianliechti/llama/pkg/provider"
)

func (s *Server) handleAudioTranscription(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	model := r.FormValue("model")

	transcriber, err := s.Transcriber(model)

	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	prompt := r.FormValue("p
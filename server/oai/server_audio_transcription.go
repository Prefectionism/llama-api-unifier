package oai

import (
	"net/http"

	"github.com/adrianliechti/llama/pkg/provider"
)

func (s *Server) handleAudioTranscription(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultip
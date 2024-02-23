package oai

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/adrianliechti/llama/pkg/provider"
)

func (s *Server) handleAudioSpeech(w http.ResponseWriter, r *http.Request) {
	var req SpeechReque
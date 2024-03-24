package ollama

import "net/http"

func (s *Server) handleHeartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusO
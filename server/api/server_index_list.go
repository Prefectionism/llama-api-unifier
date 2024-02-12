
package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleIndexList(w http.ResponseWriter, r *http.Request) {
	i, err := s.Index(chi.URLParam(r, "index"))
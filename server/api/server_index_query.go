package api

import (
	"encoding/json"
	"net/http"

	"github.com/adrianliechti/llama/pkg/index"
	"github.com/adrianliechti/llama/pkg/to"

	"github.com/go-chi/chi/v5"
)

func (s *Server) handleIndexQuery(w http.ResponseWri
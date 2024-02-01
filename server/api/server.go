
package api

import (
	"encoding/json"
	"mime"
	"net/http"
	"path"

	"github.com/adrianliechti/llama/config"
	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	*config.Config
	http.Handler
}

func New(cfg *config.Config) (*Server, error) {
	r := chi.NewRouter()

	s := &Server{
		Config:  cfg,
		Handler: r,
	}

	r.Post("/extract/{extractor}", s.handleExtract)
	r.Post("/translate/{translator}", s.handleTranslate)

	r.Get("/index/{index}", s.handleIndexList)

	r.Post("/index/{index}", s.handleIndexIngest)
	r.Delete("/index/{index}", s.handleIndexDeletion)

	r.Post("/index/{index}/query", s.handleIndexQuery)
	r.Post("/index/{index}/{extractor}", s.handleIndexWithExtractor)

	return s, nil
}

func writeJson(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	enc.Encode(v)
}

func writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
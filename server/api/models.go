package api

type Document struct {
	ID string `json:"id,omitempty"`

	Content string `json:"content,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`
}

type Result struct {
	Document `jso
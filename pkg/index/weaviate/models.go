package weaviate

type Object struct {
	ID string `json:"id"`

	Created int64 `json:"creationTimeUnix"`
	Upda
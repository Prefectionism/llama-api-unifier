package tavily

type SearchResult struct {
	Query string `json:"query"`

	Answer string `json:"answer"`

	
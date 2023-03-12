package aisearch

type Results struct {
	Value []Result `json:"value"`
}

type Result map[string]any

func (r Result) String(name string) string {
	val, ok := r[name]

	if !ok {
		return ""
	}

	data, ok := val.(string)

	if !ok {
		return ""
	}

	return data
}

func (r Result) ID() string
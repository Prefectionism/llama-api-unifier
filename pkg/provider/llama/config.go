package llama

import (
	"net/http"

	"github.com/adrianliechti/llama/pkg/provider/openai"
)

type Config struct {
	options []openai.Option
}

typ
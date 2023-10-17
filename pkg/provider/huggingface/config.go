package huggingface

import (
	"net/http"
)

type Config struct {
	url string

	token string
	model string

	client *http.Client
}

type Option
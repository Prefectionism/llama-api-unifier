package main

import (
	"context"
	"encoding/base64"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	config := openai.DefaultConfig("")
	config.BaseURL = "http://localhost:8080/oai/v1"

	model := "stable-diffusion"
	prompt := "a cute baby sea otter"

	client := openai.NewClie
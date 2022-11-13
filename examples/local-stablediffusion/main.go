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

	client := openai.NewClientWithConfig(config)

	image, err := client.CreateImage(context.Background(), openai.ImageRequest{
		Model:  model,
		Prompt: prompt,

		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	})

	if err != nil {
		panic(err)
	}

	da
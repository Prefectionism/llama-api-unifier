
package main

import (
	"context"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

func main() {
	ctx := context.Background()

	config := openai.DefaultConfig("")
	config.BaseURL = "http://localhost:8080/oai/v1"

	completionModel := "mistral-7b-instruct"
	synthesizerModel := "tts-1"
	transcriptionModel := "whisper-1"

	client := openai.NewClientWithConfig(config)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a voice assistant. Answer short and concise!",
		},
	}

	println("Press enter to start / stop recording")

	for {
		waitForEnter()
		println("Recording...")

		recordCtx, recordCancel := context.WithCancel(ctx)

		go func() {
			output, err := record(recordCtx)

			if err != nil {
				println(err.Error())
				return
			}

			defer os.Remove(output)

			transcription, err := transcribe(ctx, client, transcriptionModel, output)

			if err != nil {
				println(err.Error())
				return
			}

			prompt := transcription.Text
			println("> " + prompt)

			if prompt == "" || prompt == "[BLANK_AUDIO]" {
				return
			}

			if transcription.Language != "" {
				println("Language: " + transcription.Language)
				prompt = strings.TrimRight(prompt, ".") + ". Answer in " + transcription.Language + "."
			}

			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			})

			completion, err := complete(ctx, client, completionModel, messages)

			if err != nil {
				println(err.Error())
				return
			}

			answer := completion.Choices[0].Message.Content

			answer = regexp.MustCompile(`\[.*?\]|\(.*?\)`).ReplaceAllString(answer, "")
			answer = regexp.MustCompile(`\[.*?\]?|\(.*?\)?`).ReplaceAllString(answer, "")
			answer = strings.Split(answer, "\n")[0]

			println("< " + answer)

			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: answer,
			})

			data, err := synthesize(ctx, client, synthesizerModel, transcription.Language, answer)

			if err != nil {
				println(err.Error())
				return
			}

			f, err := os.CreateTemp("", "voicebot-*.wav")

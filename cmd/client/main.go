package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	urlFlag := flag.String("url", "http://localhost:8080/oai/v1", "server url")
	tokenFlag := flag.String("token", "", "server token")
	modelFlag := flag.String("model", "", "model id")

	flag.Parse()

	ctx := context.Background()

	reader := bufio.NewReader(os.Stdin)
	output := os.Stdout

	config := openai.DefaultConfig(*tokenFlag)
	config.BaseURL = *urlFlag

	client := openai.NewClientWithConfig(config)
	model := *modelFlag

	if model == "" {
		list, err := client.ListModels(ctx)

		if err != nil {
			panic(err)
		}

		sort.SliceStable(list.Models, func(i, j int) bool {
			return list.Models[i].ID < list.Models[j].ID
		})

		for i, m := range list.Models {
			output.WriteString(fmt.Sprintf("%2d) ", i+1))
			output.WriteString(m.ID)
			output.WriteString("\n")
		}

		output.WriteString(" >  ")
		sel, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		idx, err := strconv.Atoi(strings.TrimSpace(sel))

		if err != nil {
			panic(err)
		}

		model = list.Models[idx-1].ID
		output.WriteString("\n")
	}

	var messages []openai.ChatCompletionMessage

LOOP:
	for {
		output.WriteString(">>> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		input = strings.TrimSpace(input)

		i
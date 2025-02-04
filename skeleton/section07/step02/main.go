package main

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	var message string
	fmt.Print(">")
	fmt.Scanln(&message)

	// TODO: OpenAIのAPIクライアントを生成する

	completion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	})

	if err != nil {
		return err
	}

	// TODO: "AI:"のあとに空白区切りでAIのコメントを出力する
	// AIのコメントはcompletion.Choices[0].Message.Contentから取れる

	return nil
}

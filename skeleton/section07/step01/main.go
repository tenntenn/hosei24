package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	const url = "https://api.openai.com/v1/models"
	// TODO: GETメソッドのリクエストを生成する（コンテキストを使用する）

	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	// TODO: http.DefaultClientでリクエストを送る

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 標準出力にダンプする
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		return err
	}

	return nil
}

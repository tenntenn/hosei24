package main

import (
	"fmt"
	"os"

	"github.com/tenntenn/hosei24/section07/step05/aichat"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	chat, err := aichat.New(":8080")
	if err != nil {
		return err
	}

	if err := chat.Start(); err != nil {
		return err
	}

	return nil
}

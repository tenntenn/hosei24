package main

import (
	"flag"
	"fmt"
	"os"
	_ "image/png"
	_ "image/jpeg"

	imgcheck "github.com/tenntenn/hosei24/section06/step07"
)

var (
	flagFormat string
	flagMaxHeight int
	flagMaxWidth int
)

func init() {
	flag.StringVar(&flagFormat, "format", "", "allow image format")
	flag.IntVar(&flagMaxHeight, "height", -1, "max height")
	flag.IntVar(&flagMaxWidth, "width", -1, "max width")
}

func main() {
	flag.Parse()

	var rules []imgcheck.Rule
	if flagFormat != "" {
		rules = append(rules, imgcheck.Format(flagFormat))
	}
	if flagMaxHeight > 0 {
		rules = append(rules, imgcheck.MaxHeight(flagMaxHeight))
	}
	if flagMaxWidth > 0 {
		rules = append(rules, imgcheck.MaxWidth(flagMaxWidth))
	}

	if err := imgcheck.ValidateDir(flag.Arg(0), rules...); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}

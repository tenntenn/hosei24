package main

import (
	"os"

	eventcal "github.com/tenntenn/hosei24/section06/step02"
)

func main() {
	cli := eventcal.CLI{
		Calendar: eventcal.NewCalendar(),
	}
	os.Exit(cli.Main())
}

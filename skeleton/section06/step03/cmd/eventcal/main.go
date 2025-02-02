package main

import (
	"os"

	eventcal "github.com/tenntenn/hosei24/section06/step03"
)

func main() {
	cli := eventcal.CLI{
		Calendar: eventcal.NewCalendar(),
	}
	os.Exit(cli.Main())
}

package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	_ "go.uber.org/automaxprocs"
)

const (
	ExitCode    = 1
	ExitMessage = "unable to execute apollgo"
)

var Version = "unknown"

func main() {
	app := cli.NewApp()
	app.Version = Version

	app.Name = "apollgo"
	app.Usage = "disturbed radio-streaming integration server"
	app.Action = func(context *cli.Context) error {
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(ExitMessage)
		os.Exit(ExitCode)
	}
}

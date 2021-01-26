package main

import (
	"go-third-cli-tool/cmds/remove_dir_file"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const Version = "0.1.0"

func main() {
	app := &cli.App{
		Name:    "go-third-cli-tool",
		Usage:   "a tool help you forward to success",
		Version: Version,
		Commands: []*cli.Command{
			remove_dir_file.Cmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

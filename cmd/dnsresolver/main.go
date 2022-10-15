package main

import (
	"github.com/parinpan/dns-resolver/app/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Usage: "to resolve dns with a given hostname",
		Action: func(ctx *cli.Context) error {
			return server.Start(ctx.Context, ":80")
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "Healthchecker",
		Usage: "A tiny tool that check whether a website is running or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")

			if ctx.String(port) == "" {
				port = "80"
			}

			status := Check(ctx.String("domain"), port)

			fmt.Println(status)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

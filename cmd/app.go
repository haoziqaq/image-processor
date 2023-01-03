package main

import (
	"github.com/haoziqaq/image-processor/internal/actions"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func setupApp() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "resize",
				Usage: "Resize images",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Usage:   "image workspace dirname",
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "image output dirname",
					},
					&cli.IntFlag{
						Name:    "width",
						Aliases: []string{"w"},
						Usage:   "image width",
						Value:   0,
					},
					&cli.IntFlag{
						Name:    "height",
						Aliases: []string{"he"},
						Usage:   "image height",
						Value:   0,
					},
				},
				Action: func(context *cli.Context) error {
					actions.Resize(context)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

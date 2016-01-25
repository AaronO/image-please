package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	// App meta-data
	app := cli.NewApp()
	app.Version = "0.3.0"
	app.Name = "image-please"
	app.Author = "Aaron O'Mullan"
	app.Email = "aaron@gitbook.com"
	app.Usage = "Images best served cold"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "port, p",
			Value:  "9999",
			Usage:  "Port to listen on",
			EnvVar: "PORT",
		},
	}

	// Main app code
	app.Action = func(ctx *cli.Context) {
		// Log port
		log.Println("Listening on", ctx.String("port"))
		// Run
		if err := RunServer(ctx.String("port")); err != nil {
			log.Fatal(err)
		}
	}

	// Parse cli args and run :)
	app.Run(os.Args)
}

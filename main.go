package main

import (
	"fmt"
	"os"

	"github.com/anatolio-deb/gosplash/actions"
	"github.com/urfave/cli"
)

func main() {
	var searchTerm string
	app := &cli.App{
		Name: "gosplash",
		// HelpName:    "GoSplash",
		// Usage:       "gosplash random -query batman",
		Version:     "0.1a",
		Description: "GoSplash is a tool for changing wallpapers randomly on Pantheon desktop environment of elementaryOS.",
		Commands: []cli.Command{
			{
				Name:      "random",
				ShortName: "r",
				Aliases:   []string{"any"},
				// Usage:       "random -search batman",
				Description: "Get a random wallpapper",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "search",
						// Usage:       "-search batman",
						Required:    false,
						Destination: &searchTerm,
					},
				},
				Action: func(c *cli.Context) error { return actions.GetRandomPhoto(searchTerm) },
			},
		},
	}

	if len(os.Args) == 0 {
		return
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		return
	}
}

package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/anatolio-deb/gosplash/common"
	"github.com/anatolio-deb/gosplash/random"
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
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
	var p *common.Photo

	photos, err := random.Get(searchTerm)

	if err != nil {
		log.Fatal(err)
	}

	p = &photos[0]

	link, err := common.GetDownloadURL(p.Links.DownloadLocation)
	if err != nil {
		log.Fatal(err)
	}
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	var filepath string
	fm := u.Query().Get("fm")
	if len(fm) == 0 {
		fm = "jpg"
	}
	filepath, err = download(u.String(), p.ID+fmt.Sprintf(".%s", fm))
	if err != nil {
		log.Fatal(err)
	}
	err = setWallpaper(filepath)
	if err != nil {
		log.Fatal(err)
	}
}

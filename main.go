package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/anatolio-deb/gosplash/common"
	"github.com/anatolio-deb/gosplash/random"
	"github.com/anatolio-deb/gosplash/searching"
	"github.com/anatolio-deb/gosplash/topics"
)

func main() {
	randomCMD := flag.NewFlagSet("random", flag.ContinueOnError)
	tptr := randomCMD.String("topics", "", "a comma separated list of Unsplash topics, e.g. nature,wallpapers")
	searchCMD := flag.NewFlagSet("random", flag.ContinueOnError)
	sptr := searchCMD.String("query", "", "a keyphrase to search")
	cptr := searchCMD.String("color", "", "a color of images to search")
	flag.Usage = func() {
		fmt.Println(`GoSplash is a tool for changing wallpapers randomly on Pantheon desktop environment of elementaryOS.

Usage:
	gosplash random -topics=nature,wallpapers
	gosplash search -query=batman -color=dark
	gosplash random`)
	}
	flag.Parse()
	var p common.Photo
	var err error
	slugs := strings.Split(*tptr, ",")
	ids := make([]string, len(slugs))
	query := *sptr
	color := *cptr
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == randomCMD.Name() && len(slugs) > 1 {
			for i := 0; i < len(slugs); i++ {
				s := slugs[i]
				t, err := topics.Get(s)
				if err != nil {
					log.Println(err)
				} else {
					ids[i] = t.ID
				}
			}
		} else if args[0] == searchCMD.Name() {
			if len(query) > 1 {
				p, err = searching.Get(query, color)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				p, err = random.Get(ids)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	if len(p.ID) > 0 {
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
	} else {
		flag.Usage()
	}
}

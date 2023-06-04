package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/anatolio-deb/gosplash/random"
)

func main() {
	ptr := flag.String("topics", "", "a comma separated list of Unsplash topics, e.g. nature,ocean")
	flag.Parse()
	slugs := strings.Split(*ptr, ",")
	ids := make([]string, len(slugs))
	for i := 0; i < len(slugs); i++ {
		s := slugs[i]
		t, err := topics.Get(s)
		if err != nil {
			log.Println(err)
		} else {
			ids[i] = t.ID
		}
	}
	log.Println(ids)
	p, err := random.Get(ids)
	if err != nil {
		log.Fatal(err)
	}
	link, err := random.GetDownloadURL(p.Links.DownloadLocation)
	if err != nil {
		log.Fatal(err)
	}
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	filepath, err := download(link, p.ID+fmt.Sprintf(".%s", u.Query()["fm"][0]))
	if err != nil {
		log.Fatal(err)
	}
	err = setWallpaper(filepath)
	if err != nil {
		log.Fatal(err)
	}
}

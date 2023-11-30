package actions

import (
	"fmt"
	"net/url"

	"github.com/anatolio-deb/gosplash/common"
	"github.com/anatolio-deb/gosplash/random"
)

func GetRandomPhoto(searchTerm string) error {
	var p *common.Photo

	photos, err := random.Get(searchTerm)

	if err != nil {
		return err
	}

	p = &photos[0]

	link, err := common.GetDownloadURL(p.Links.DownloadLocation)
	if err != nil {
		return err
	}
	u, err := url.Parse(link)
	if err != nil {
		return err
	}
	var filepath string
	fm := u.Query().Get("fm")
	if len(fm) == 0 {
		fm = "jpg"
	}
	filepath, err = download(u.String(), p.ID+fmt.Sprintf(".%s", fm))
	if err != nil {
		return err
	}
	return setWallpaper(filepath)
}

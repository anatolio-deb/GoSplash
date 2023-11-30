package random

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/anatolio-deb/gosplash/common"
)

func Get(searchTerm string) ([]common.Photo, error) {
	var photos []common.Photo
	url, err := url.Parse(common.APIURL)
	if err != nil {
		return photos, err
	}
	url.Path = "/photos/random/"
	params := url.Query()
	params.Add("content_filter", "high")
	params.Add("orientation", "landscape")
	params.Add("count", "1")
	if len(searchTerm) > 0 {
		params.Add("query", searchTerm)
	}
	url.RawQuery = params.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url.String(), strings.NewReader(params.Encode()))

	if err != nil {
		return photos, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", common.AccessKey))

	if err != nil {
		return photos, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return photos, err
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return photos, err
	}

	return photos, json.Unmarshal(b, &photos)
}

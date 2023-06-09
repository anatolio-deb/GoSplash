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

func Get(topicIDs []string) (common.Photo, error) {
	var p common.Photo
	url, err := url.Parse(common.APIURL)
	if err != nil {
		return p, err
	}
	url.Path = "/photos/random/"
	params := url.Query()
	params.Add("content_filter", "high")
	params.Add("orientation", "landscape")
	if len(topicIDs) > 0 {
		params.Add("topics", strings.Join(topicIDs, ","))
	}
	url.RawQuery = params.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url.String(), strings.NewReader(params.Encode()))

	if err != nil {
		return p, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", common.AccessKey))

	if err != nil {
		return p, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return p, err
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return p, err
	}

	return p, json.Unmarshal(b, &p)
}

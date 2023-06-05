package searching

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/anatolio-deb/gosplash/common"
)

func Get(query string, color string) (common.Photo, error) {
	var p common.Photo
	var response struct {
		Total      int            `json:"total"`
		TotalPages int            `json:"total_pages"`
		Results    []common.Photo `json:"results"`
	}
	url, err := url.Parse(common.APIURL)
	if err != nil {
		return p, err
	}
	url.Path = "/photos/search/"
	params := url.Query()
	params.Add("query", query)
	params.Add("color", color)
	params.Add("content_filter", "high")
	params.Add("orientation", "landscape")
	params.Add("per_page", "1")
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

	err = json.Unmarshal(b, &response)
	if err != nil {
		return p, err
	}
	p = response.Results[0]
	return p, err
}

package random

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type DownloadURL struct {
	URL string `json:"url"`
}

type Links struct {
	Self             string `json:"self"`
	HTML             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}

type Photo struct {
	ID    string `json:"id"`
	Links Links  `json:"links"`
}

func Get(topicIDs []string) (Photo, error) {
	var p Photo
	url, err := url.Parse("https://api.unsplash.com")
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

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", ""))

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

func GetDownloadURL(url string) (string, error) {
	var r DownloadURL
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return r.URL, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", ""))
	resp, err := client.Do(req)
	if err != nil {
		return r.URL, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return r.URL, err
	}

	return r.URL, json.Unmarshal(b, &r)
}

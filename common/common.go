package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const APIURL = "https://api.unsplash.com"
const AccessKey = "MFSCHMjRZROiBxyaShpoIy6aW0zq3KSmgWe3j82FtGw"

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

func GetDownloadURL(url string) (string, error) {
	var r DownloadURL
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return r.URL, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", AccessKey))
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

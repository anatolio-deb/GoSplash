package topics

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Topic struct {
	ID string `json:"id"`
}

func Get(slug string) (Topic, error) {
	var t Topic
	url, err := url.Parse("")
	if err != nil {
		return t, err
	}
	url.Path = fmt.Sprintf("/topics/%s", slug)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)

	if err != nil {
		return t, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", ""))

	if err != nil {
		return t, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return t, err
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return t, err
	}

	return t, json.Unmarshal(b, &t)
}

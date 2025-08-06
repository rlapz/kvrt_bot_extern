package util

import (
	"io"
	"net/http"
)

func FetchGet(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

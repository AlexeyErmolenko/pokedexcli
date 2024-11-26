package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

var baseURL = "https://pokeapi.co/api/v2/"

func requestError(err error) error { return fmt.Errorf("request error: %w", err) }

func sendingError(err error) error {
	return fmt.Errorf("sending request error: %w", err)
}

func readResponseError(err error) error {
	return fmt.Errorf("reading response error: %w", err)
}

func parseError(err error) error {
	return fmt.Errorf("parsing error: %w", err)
}

func (c *Client) makeGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, requestError(err)
	}

	req.Header.Set("Accept", "application/json")
	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, sendingError(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, readResponseError(err)
	}

	return body, nil
}

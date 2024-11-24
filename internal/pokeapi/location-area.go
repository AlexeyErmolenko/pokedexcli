package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Name string
	URL  string
}

func (c *Client) GetLocationAreas(url string) ([]byte, error) {
	if len(url) == 0 {
		url = baseURL + "location-area/"
	}

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

func (c *Client) ParseLocationAreas(body []byte) (*Pagination[LocationArea], error) {
	var resStruct Pagination[LocationArea]

	if err := json.Unmarshal(body, &resStruct); err != nil {
		return nil, parseError(err)
	}

	return &resStruct, nil
}

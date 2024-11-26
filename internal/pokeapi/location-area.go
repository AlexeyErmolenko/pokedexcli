package pokeapi

import (
	"encoding/json"
)

type LocationArea struct {
	Name string
	URL  string
}

func (c *Client) GetLocationAreas(url string) ([]byte, error) {
	if len(url) == 0 {
		url = baseURL + "location-area/"
	}

	return c.makeGetRequest(url)
}

func (c *Client) ParseLocationAreas(body []byte) (*Pagination[LocationArea], error) {
	var resStruct Pagination[LocationArea]

	if err := json.Unmarshal(body, &resStruct); err != nil {
		return nil, parseError(err)
	}

	return &resStruct, nil
}

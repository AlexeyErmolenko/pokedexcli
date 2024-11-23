package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Name string
	URL  string
}

func (c *Client) GetLocationAreas(curPage int) (*Pagination[LocationArea], error) {
	path := "location-area/"
	query := fmt.Sprintf("?offset=%d&limit=%d", curPage*defaultLimit, defaultLimit)
	req, err := http.NewRequest(http.MethodGet, baseURL+path+query, nil)

	if err != nil {
		return &Pagination[LocationArea]{}, requestError(err)
	}

	req.Header.Set("Accept", "application/json")
	res, err := c.httpClient.Do(req)

	if err != nil {
		return &Pagination[LocationArea]{}, sendingError(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return &Pagination[LocationArea]{}, readResponseError(err)
	}

	var resStruct Pagination[LocationArea]
	err = parseLocationAreas(body, &resStruct)

	if err != nil {
		return &Pagination[LocationArea]{}, err
	}

	return &resStruct, nil
}

func parseLocationAreas(body []byte, resStruct *Pagination[LocationArea]) error {
	if err := json.Unmarshal(body, resStruct); err != nil {
		return parseError(err)
	}

	return nil
}

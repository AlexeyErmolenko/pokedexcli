package api

import "fmt"

var baseURL = "https://pokeapi.co/api/v2/"
var defaultLimit = 20

func requestError(err error) error {
	return fmt.Errorf("request error: %w", err)
}

func sendingError(err error) error {
	return fmt.Errorf("sending request error: %w", err)
}

func readResponseError(err error) error {
	return fmt.Errorf("reading response error: %w", err)
}

func parseError(err error) error {
	return fmt.Errorf("parsing error: %w", err)
}

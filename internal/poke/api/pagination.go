package api

type Pagination[T any] struct {
	Results  []T    `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Count    int    `json:"count"`
}

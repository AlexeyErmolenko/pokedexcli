package pokeapi

import (
	"encoding/json"
	"fmt"
)

type PokemonDetail struct {
	Abilities              []interface{} `json:"abilities"`
	BaseExperience         int           `json:"base_experience"`
	Cries                  interface{}   `json:"cries"`
	Forms                  []interface{} `json:"forms"`
	GameIndices            []interface{} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []interface{} `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int           `json:"order"`
	PastAbilities          []interface{} `json:"past_abilities"`
	PastTypes              []interface{} `json:"past_types"`
	Species                interface{}   `json:"species"`
	Sprites                interface{}   `json:"sprites"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (c *Client) GetPokemon(pokemonName string) ([]byte, error) {
	if len(pokemonName) == 0 {
		return nil, pokemonNameEmptyError()
	}
	url := baseURL + "/pokemon/" + pokemonName

	return c.makeGetRequest(url)
}

func (c *Client) ParsePokemon(body []byte) (*PokemonDetail, error) {
	var pokemonDetail PokemonDetail

	if err := json.Unmarshal(body, &pokemonDetail); err != nil {
		return nil, parseError(err)
	}

	return &pokemonDetail, nil
}

func pokemonNameEmptyError() error {
	return fmt.Errorf("pokemon name is empty")
}

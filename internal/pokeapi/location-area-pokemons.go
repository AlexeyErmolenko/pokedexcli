package pokeapi

import (
	"encoding/json"
	"fmt"
)

type AreaDetail struct {
	EncounterMethodRates []interface{} `json:"encounter_method_rates"`
	GameIndex            int           `json:"game_index"`
	ID                   int           `json:"id"`
	Location             LocationArea  `json:"location"`
	Name                 string        `json:"name"`
	Names                []interface{} `json:"names"`
	PokemonEncounters    []struct {
		Pokemon        Pokemon       `json:"pokemon"`
		VersionDetails []interface{} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
type Pokemon struct {
	Name string
	URL  string
}

func (c *Client) GetPokemonsLocationArea(areaName string) ([]byte, error) {
	if len(areaName) == 0 {
		return nil, areaIsEmptyError()
	}

	url := baseURL + "location-area/" + areaName

	return c.makeGetRequest(url)
}

func (c *Client) ParsePokemonsLocationArea(body []byte) (*[]Pokemon, error) {
	var areaDetail AreaDetail

	if err := json.Unmarshal(body, &areaDetail); err != nil {
		return nil, parseError(err)
	}

	var pokemons []Pokemon

	for _, i := range areaDetail.PokemonEncounters {
		pokemons = append(pokemons, i.Pokemon)
	}

	return &pokemons, nil
}

func areaIsEmptyError() error {
	return fmt.Errorf("area is empty")
}

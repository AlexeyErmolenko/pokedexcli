package main

import (
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
)

func commandExplore(conf *config, args []string) error {
	if len(args[0]) == 0 {
		return fmt.Errorf("area shouldn't be empty")
	}

	areaName := args[0]
	body, ok := conf.pokeCache.Get(areaName)
	if !ok {
		res, err := conf.pokeApiClient.GetPokemonsLocationArea(areaName)

		if err != nil {
			return err
		}
		body = res
		conf.pokeCache.Add(areaName, body)
	}

	data, err := conf.pokeApiClient.ParsePokemonsLocationArea(body)
	if err != nil {
		return err
	}

	showPokemons(data)

	return nil
}

func showPokemons(data *[]pokeapi.Pokemon) {
	fmt.Printf("%s\n", escSQ+tealColor+end+"Found Pokemon:"+escSQ+regularStyle+end)
	for _, poke := range *data {
		fmt.Printf("\t%s\n", escSQ+tealColor+end+"- "+poke.Name+escSQ+regularStyle+end)
	}
}

package main

import (
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config) error {
	page, err := conf.pokeApiClient.GetLocationAreas(conf.nextPage)
	if err != nil {
		return err
	}

	showLocationAreas(&page.Results)

	if conf.nextPage != 0 {
		conf.prevPage++
	}

	conf.nextPage++

	return nil
}

func commandMapB(conf *config) error {
	if conf.prevPage == 0 {
		return fmt.Errorf("you're on the first page")
	}
	conf.prevPage--
	conf.nextPage--

	page, err := conf.pokeApiClient.GetLocationAreas(conf.prevPage)
	if err != nil {
		return err
	}

	showLocationAreas(&page.Results)

	return nil
}

func showLocationAreas(locationAreas *[]pokeapi.LocationArea) {
	for _, locationArea := range *locationAreas {
		if locationArea.Name == "" {
			continue
		}

		fmt.Printf("%s\n", escSQ+tealColor+end+locationArea.Name+escSQ+regularStyle+end)
	}
}

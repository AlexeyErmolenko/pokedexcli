package main

import (
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config) error {
	page, err := conf.pokeApiClient.GetLocationAreas(conf.nextLocationUrl)
	if err != nil {
		return err
	}

	showLocationAreas(&page.Results)

	conf.nextLocationUrl = page.Next
	conf.prevLocationUrl = page.Previous

	return nil
}

func commandMapB(conf *config) error {
	page, err := conf.pokeApiClient.GetLocationAreas(conf.prevLocationUrl)
	if err != nil {
		return err
	}

	showLocationAreas(&page.Results)
	conf.nextLocationUrl = page.Next
	conf.prevLocationUrl = page.Previous

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

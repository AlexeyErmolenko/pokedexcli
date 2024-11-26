package main

import (
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
)

func commandMap(conf *config, args []string) error {
	return getLocation(conf, conf.nextLocationUrl)
}

func commandMapB(conf *config, args []string) error {
	return getLocation(conf, conf.prevLocationUrl)
}

func getLocation(conf *config, url string) error {
	body, ok := conf.pokeCache.Get(url)

	if ok != true {
		data, err := conf.pokeApiClient.GetLocationAreas(url)

		if err != nil {
			return err
		}

		body = data

		conf.pokeCache.Add(url, body)
	}

	page, err := conf.pokeApiClient.ParseLocationAreas(body)
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

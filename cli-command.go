package main

import (
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/poke/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var currenAreaLocationPage int

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world.",
			callback:    commandNextLocations,
		},
		"mapb": {
			name:        "mapb",
			description: "The mapb displays the previous 20 locations. It's a way to go back.",
			callback:    commandPreviousLocations,
		},
	}
}

func commandHelp() error {
	fmt.Println(fmt.Sprintf("%s\nWelcome to the Pokedex!", escseq+tealColor+end))
	fmt.Print("Usage: \n\n")

	for _, command := range getCliCommands() {
		fmt.Printf("%s%s: %s%s\n", escseq+tealColor+end, command.name, command.description, escseq+regularStyle+end)
	}

	fmt.Print("\n")

	return nil
}

func commandExit() error {
	return fmt.Errorf("exit status 1")
}

func commandNextLocations() error {
	if currenAreaLocationPage != 0 {
		currenAreaLocationPage++
	}

	page, err := api.GetLocationAreas(currenAreaLocationPage)

	if err != nil {
		return err
	}

	showLocationAreas(&page.Results)

	if currenAreaLocationPage == 0 {
		currenAreaLocationPage++
	}

	return nil
}

func commandPreviousLocations() error {
	if currenAreaLocationPage == 0 {
		return fmt.Errorf("no locations available")
	}

	currenAreaLocationPage--

	if currenAreaLocationPage == 0 {
		return fmt.Errorf("no locations available")
	}

	page, err := api.GetLocationAreas(currenAreaLocationPage)
	if err != nil {
		return err
	}

	showLocationAreas(&page.Results)
	return nil
}

func showLocationAreas(locationAreas *[]api.LocationArea) {
	for _, locationArea := range *locationAreas {
		if locationArea.Name == "" {
			continue
		}

		fmt.Printf("%s\n", escseq+tealColor+end+locationArea.Name+escseq+regularStyle+end)
	}
}

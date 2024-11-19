package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage: \n\n")

	for _, command := range getCliCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Print("\n")

	return nil
}

func commandExit() error {
	return fmt.Errorf("exiting Pokedex")
}

func commandNextLocations() error {
	return nil
}

func commandPreviousLocations() error {
	return nil
}

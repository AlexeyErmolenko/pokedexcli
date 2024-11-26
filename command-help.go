package main

import "fmt"

func commandHelp(conf *config, args []string) error {
	fmt.Println(prepareWelcomeLine())
	fmt.Print("Usage: \n\n")

	for _, command := range getCliCommands() {
		fmt.Printf(prepareHelpCommandLine(command))
	}

	fmt.Println()

	return nil
}

func prepareWelcomeLine() string {
	return fmt.Sprintf("%s\nWelcome to the Pokedex!", escSQ+tealColor+end)
}

func prepareHelpCommandLine(command cliCommand) string {
	return fmt.Sprintf("%s%s: %s%s\n", escSQ+tealColor+end, command.name, command.description, escSQ+regularStyle+end)
}

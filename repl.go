package main

import (
	"bufio"
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokecache"
	"os"
)

type config struct {
	pokeApiClient   pokeapi.Client
	pokeCache       *pokecache.Cache
	nextLocationUrl string
	prevLocationUrl string
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()

	for {
		fmt.Printf(preComL())
		ok := scanner.Scan()
		if !ok {
			break
		}

		commandName := scanner.Text()
		command, ok := cliCommands[commandName]
		if !ok {
			continue
		}

		err := command.callback(conf)

		if err != nil {
			fmt.Printf(preErrL(err))
		}
	}
}

func preComL() string {
	return fmt.Sprintf("%sPokedex %s> %s", escSQ+tealColor+end, escSQ+blueColor+end, escSQ+greenColor+end)
}

func preErrL(err error) string {
	return fmt.Sprintf("%s%w%s\n", escSQ+redColor+end, err, escSQ+regularStyle+end)
}

package main

import (
	"bufio"
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokecache"
	"os"
	"strings"
)

type config struct {
	pokeApiClient   pokeapi.Client
	pokeCache       *pokecache.Cache
	nextLocationUrl string
	prevLocationUrl string
	caughtPokemon   map[string]pokeapi.PokemonDetail
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

		commandLine := scanner.Text()
		arguments := strings.Split(commandLine, " ")
		if len(arguments) < 1 {
			fmt.Printf(preErrL(fmt.Errorf("command not given")))
		}

		commandName := arguments[0]
		command, ok := cliCommands[commandName]
		if !ok {
			continue
		}

		commandArgs := arguments[1:]
		err := command.callback(conf, commandArgs)

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

package main

import (
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokecache"
	"time"
)

const (
	escSQ        = "\033["
	regularStyle = "0"
	end          = "m"
	divider      = ";"
)

const (
	tealColor  = "36"
	blueColor  = "34"
	greenColor = "32"
	redColor   = "31"
)

func main() {
	pokeCache := pokecache.NewCache(60)
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		pokeApiClient: pokeClient,
		pokeCache:     pokeCache,
		caughtPokemon: map[string]pokeapi.PokemonDetail{},
	}

	startRepl(conf)
}

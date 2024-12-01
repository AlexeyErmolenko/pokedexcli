package main

import (
	"fmt"
	"github.com/AlexeyErmolenko/pokedexcli/internal/pokeapi"
	"strconv"
)

func commandInspect(conf *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("a name shoudn't be empty")
	}

	name := args[0]
	pokemon, ok := conf.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("no such pokemon")
	}

	showPokemonDetails(pokemon)

	return nil
}

func showPokemonDetails(pokemon pokeapi.PokemonDetail) {
	fmt.Printf("%s\n", escSQ+tealColor+end+"Name: "+pokemon.Name+escSQ+regularStyle+end)
	fmt.Printf("%s\n", escSQ+tealColor+end+"Height: "+strconv.Itoa(pokemon.Height)+escSQ+regularStyle+end)
	fmt.Printf("%s\n", escSQ+tealColor+end+"Weight: "+strconv.Itoa(pokemon.Weight)+escSQ+regularStyle+end)

	fmt.Printf("%s\n", escSQ+tealColor+end+"Stats:"+escSQ+regularStyle+end)
	for _, s := range pokemon.Stats {
		fmt.Printf("\t%s\n", escSQ+tealColor+end+"-"+s.Stat.Name+": "+strconv.Itoa(s.BaseStat)+escSQ+regularStyle+end)
	}

	fmt.Printf("%s\n", escSQ+tealColor+end+"Types:"+escSQ+regularStyle+end)
	for _, t := range pokemon.Types {
		fmt.Printf("\t%s\n", escSQ+tealColor+end+"-"+t.Type.Name+escSQ+regularStyle+end)
	}
}

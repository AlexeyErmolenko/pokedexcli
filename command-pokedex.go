package main

import "fmt"

func commandPokedex(conf *config, args []string) error {
	if len(conf.caughtPokemon) == 0 {
		return fmt.Errorf("you don't have any caught pokemon")
	}

	fmt.Printf("%s\n", escSQ+tealColor+end+"Your Pokedex:"+escSQ+regularStyle+end)

	for name, _ := range conf.caughtPokemon {
		fmt.Printf("\t%s\n", escSQ+tealColor+end+"- "+name+escSQ+regularStyle+end)
	}

	return nil
}

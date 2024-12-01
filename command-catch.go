package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("you should provide a pokemon name")
	}

	pokemonName := args[0]
	res, err := conf.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	pokemon, err := conf.pokeApiClient.ParsePokemon(res)
	if err != nil {
		return err
	}

	catch := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if catch > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	conf.caughtPokemon[pokemon.Name] = *pokemon

	return nil
}

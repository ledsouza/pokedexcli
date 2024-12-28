package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemonResp.BaseExperience)
	threshold := pokemonResp.BaseExperience / 2

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if chance > threshold {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.pokedex[pokemonName] = pokemonResp
	return nil
}

package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]
	exploreLocationResp, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	pokemonEncounters := exploreLocationResp.PokemonEncounters
	for _, pokemonEncounter := range pokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}

	return nil
}

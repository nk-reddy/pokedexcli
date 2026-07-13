package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("need to pass in an area to explore!")
	}

	fmt.Printf("Exploring %s...\n", args[0])
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, object := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", object.Pokemon.Name)
	}

	return nil
}

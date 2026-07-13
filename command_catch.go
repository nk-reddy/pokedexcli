package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("need to pass in a pokemon to catch!")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemonResp, err := cfg.pokeapiClient.ListPokemonInfo(args[0])
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemonResp.BaseExperience)

	if chance <= 40 {
		cfg.pokedex[args[0]] = pokemonResp
		fmt.Println(args[0], "was caught!")
	} else {
		fmt.Println(args[0], "escaped!")
	}

	return nil
}

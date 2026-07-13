package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("need to pass in a pokemon to inspect!")
	}

	resp, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", resp.Name)
	fmt.Println("Height:", resp.Height)
	fmt.Println("Weight:", resp.Weight)
	fmt.Println("Stats:")
	for _, stat := range resp.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.Score)
	}
	fmt.Println("Types:")
	for _, types := range resp.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}

	return nil
}

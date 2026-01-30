package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, argv []string) error {
	fmt.Printf("Your Pokedex:\n")
	if len(c.caughtPokemon) == 0 {
		return errors.New("your Pokedex is empty")
	}
	for k, _ := range c.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}

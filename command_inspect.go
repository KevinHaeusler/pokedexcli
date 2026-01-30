package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args []string) error {
	if len(args) < 1 {
		return errors.New("No pokemon was specified")
	}
	pokemon := args[0]
	_, ok := c.caughtPokemon[pokemon]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", c.caughtPokemon[pokemon].Name)
	fmt.Printf("Height: %d\n", c.caughtPokemon[pokemon].Height)
	fmt.Printf("Weight: %d\n", c.caughtPokemon[pokemon].Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range c.caughtPokemon[pokemon].Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types\n")
	for _, typ := range c.caughtPokemon[pokemon].Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	return nil
}

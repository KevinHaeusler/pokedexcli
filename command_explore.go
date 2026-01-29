package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return errors.New("no location specified")
	}
	areaName := args[0]
	path := "location-area/" + areaName

	area, err := cfg.pokeapiClient.GetLocationArea(path)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

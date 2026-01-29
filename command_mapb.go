package main

import (
	"fmt"
	"strings"
)

func commandMapb(cfg *config, args []string) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	fullURL := *cfg.prevLocationsURL
	const base = "https://pokeapi.co/api/v2/"
	path := strings.TrimPrefix(fullURL, base)

	resp, err := cfg.pokeapiClient.ListLocationAreas(path)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

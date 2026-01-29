package main

import (
	"fmt"
	"strings"
)

func commandMap(cfg *config, args []string) error {
	var path string

	if cfg.nextLocationsURL == nil {
		path = "location-area"
	} else {
		fullURL := *cfg.nextLocationsURL
		const base = "https://pokeapi.co/api/v2/"
		path = strings.TrimPrefix(fullURL, base)
	}

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

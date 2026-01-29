package main

import "github.com/KevinHaeusler/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

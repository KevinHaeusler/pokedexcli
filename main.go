package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/KevinHaeusler/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient:    pokeapi.NewClient(5 * time.Second),
		nextLocationsURL: nil,
		prevLocationsURL: nil,
		caughtPokemon:    make(map[string]pokeapi.Pokemon),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean := cleanInput(scanner.Text())
		if len(clean) == 0 {
			continue
		}

		commandName := clean[0]
		args := clean[1:]
		command, exists := getCommands()[commandName]
		if exists {

			if err := command.callback(cfg, args); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

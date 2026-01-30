package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists Pokemon in specified Area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch an existing Pokemon",
			callback:    commandCatch,
		},
	}
}

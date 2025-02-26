package main

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

type Config struct {
	Previous string
	Next     string
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Getting next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Getting previous 20 locations, nothing if first page",
			callback:    commandMapb,
		},
	}
}

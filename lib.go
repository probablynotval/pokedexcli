package main

import (
	"net/url"
	"path"
	"strings"

	"github.com/probablynoval/pokedexcli/api"
)

type config struct {
	apiClient api.Client
	location  api.RespDeepLocations
	pokedex   map[string]api.RespPokemon
	next      *string
	prev      *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokémon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore a map for Pokémon",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Shows this help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokémon's stats",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Show the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous page of locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all caught Pokémon",
			callback:    commandPokedex,
		},
	}
}

func sanitizeInput(input string) []string {
	inputLower := strings.ToLower(input)
	words := strings.Fields(inputLower)
	return words
}

func locationId(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	idStr := path.Base(u.Path)
	return idStr, nil
}

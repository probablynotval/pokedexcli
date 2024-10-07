package main

import (
	"bufio"
	"fmt"
	"os"
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

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
	args        *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Shows this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
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
		"explore": {
			name:        "explore",
			description: "Explore a map for Pokémon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokémon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokémon's stats",
			callback:    commandInspect,
		},
	}
}

func sanitizeInput(input string) []string {
	inputLower := strings.ToLower(input)
	words := strings.Fields(inputLower)
	return words
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := sanitizeInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		commandArgs := []string{}
		if len(input) > 1 {
			commandArgs = input[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(conf, commandArgs...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}
	}
}

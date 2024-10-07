package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
	args        *string
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

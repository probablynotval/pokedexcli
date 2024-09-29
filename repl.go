package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func sanitizeInput(input string) []string {
	inputLower := strings.ToLower(input)
	words := strings.Fields(inputLower)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		switch strings.ToLower(input) {
		case "exit":
			fmt.Println("Bye!")
			return
		case "help":
			fmt.Printf("\nUsage:\nhelp: Displays this help message\nexit: Exits the program")
		default:
			fmt.Printf("\nUsage:\nhelp: Displays this help message\nexit: Exits the program")
		}
	}
}

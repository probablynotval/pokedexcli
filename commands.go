package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}

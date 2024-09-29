package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Printf("\nUsage:\nhelp: Displays this help message\nexit: Exits the program")
	return nil
}

func commandExit() error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}

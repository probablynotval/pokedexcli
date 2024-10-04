package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
)

func commandHelp(conf *config, args ...string) error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(conf *config, args ...string) error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}

func commandMap(conf *config, args ...string) error {
	locationResp, err := conf.apiClient.ListLocations(conf.Next)
	if err != nil {
		return err
	}

	conf.Next = locationResp.Next
	conf.Prev = locationResp.Prev

	for i, location := range locationResp.Results {
		locationId, err := locationId(locationResp.Results[i].URL)
		if err != nil {
			return err
		}
		fmt.Printf("%s. %s\n", locationId, location.Name)
	}
	return nil
}

func commandMapb(conf *config, args ...string) error {
	if conf.Prev == nil {
		return errors.New("You're on the first page silly!")
	}

	location, err := conf.apiClient.ListLocations(conf.Prev)
	if err != nil {
		return err
	}

	conf.Next = location.Next
	conf.Prev = location.Prev

	for i, loc := range location.Results {
		locationId, err := locationId(location.Results[i].URL)
		if err != nil {
			return err
		}

		fmt.Printf("%s. %s\n", locationId, loc.Name)
	}
	return nil
}

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Incorrect number of arguments, please enter one location's name to explore")
	}
	mapName := args[0]

	location, err := conf.apiClient.ExploreLocation(mapName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, p := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}

func locationId(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	idStr := path.Base(u.Path)
	return idStr, nil
}

package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
)

func commandHelp(conf *config) error {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(conf *config) error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}

func commandMap(conf *config) error {
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

func commandMapb(conf *config) error {
	if conf.Prev == nil {
		return errors.New("You're on the first page silly!")
	}

	locationResp, err := conf.apiClient.ListLocations(conf.Prev)
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

func locationId(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	idStr := path.Base(u.Path)
	return idStr, nil
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
	_ "math/rand"
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
	locationResp, err := conf.apiClient.ListLocations(conf.next)
	if err != nil {
		return err
	}

	conf.next = locationResp.Next
	conf.prev = locationResp.Prev

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
	if conf.prev == nil {
		return errors.New("You're on the first page silly!")
	}

	location, err := conf.apiClient.ListLocations(conf.prev)
	if err != nil {
		return err
	}

	conf.next = location.Next
	conf.prev = location.Prev

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

	location, err := conf.apiClient.GetLocation(mapName)
	if err != nil {
		return err
	}

	conf.location = location

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, p := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}

	return nil
}

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Incorrect number of arguments, please enter one Pokémon's name")
	}

	pokemonName := args[0]
	pokemonFound := false
	var pokemonIndex int

	for i, p := range conf.location.PokemonEncounters {
		if p.Pokemon.Name == pokemonName {
			pokemonFound = true
			pokemonIndex = i
			break
		}
	}

	if !pokemonFound {
		return errors.New("Incorrect argument, the Pokémon does not exist at this location")
	}

	pokemonURL := conf.location.PokemonEncounters[pokemonIndex].Pokemon.URL
	pokemon, err := conf.apiClient.GetPokemon(pokemonURL)
	if err != nil {
		return err
	}

	catchChance := 500 - pokemon.BaseExperience
	caught := rand.Intn(500) < catchChance

	if !caught {
		fmt.Printf("%s has fled!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	conf.pokedex[pokemon.Name] = pokemon

	return nil
}

func commandInspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Incorrect number of arguments, please enter one Pokémon's name")
	}

	pokemonName := args[0]

	pokemon, caught := conf.pokedex[pokemonName]
	if !caught {
		return errors.New("You have not caught that Pokémon")
	}

	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(conf *config, args ...string) error {
	if len(args) != 0 {
		fmt.Println()
		fmt.Println("Arguments ignored, command takes no arguments")
		fmt.Println()
	}

	if len(conf.pokedex) == 0 {
		return errors.New("You have not caught any Pokémon")
	}

	for _, pokemon := range conf.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
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

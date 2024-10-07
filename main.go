package main

import (
	"time"

	"github.com/probablynoval/pokedexcli/api"
)

func main() {
	client := api.NewClient(5*time.Second, 5*time.Minute)
	pokedex := map[string]api.RespPokemon{}
	conf := &config{
		apiClient: client,
		pokedex:   pokedex,
	}

	startRepl(conf)
}

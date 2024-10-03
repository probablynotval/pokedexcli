package main

import (
	"time"

	"github.com/probablynoval/pokedexcli/api"
)

func main() {
	client := api.NewClient(5 * time.Second)
	conf := &config{
		apiClient: client,
	}

	startRepl(conf)
}

package main

import (
	"time"

	"github.com/probablynoval/pokedexcli/api"
)

func main() {
	client := api.NewClient(5*time.Second, 5*time.Minute)
	conf := &config{
		apiClient: client,
	}

	startRepl(conf)
}

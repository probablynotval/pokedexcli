package api

const baseURL = "https://pokeapi.co/api/v2"

type RespShallowLocations struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

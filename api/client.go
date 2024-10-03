package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/probablynoval/pokedexcli/cache"
)

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}

func NewClient(httpTimeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: httpTimeout,
		},
		cache: cache.NewCache(cacheInterval),
	}
}

func getData(c *Client, url string) ([]byte, error) {
	cacheData, exists := c.cache.Get(url)
	if exists {
		return cacheData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)

	return data, nil
}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := getData(c, url)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationResp, nil
}

package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: []byte("data"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/420",
			val: []byte("moredata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected a key but found none")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected a value but found none")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte("data"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
	if !ok {
		t.Errorf("Expected a key but found none")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area")
	if ok {
		t.Errorf("Expected no key but found one")
		return
	}
}

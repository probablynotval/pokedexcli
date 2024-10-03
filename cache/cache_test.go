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
			val: []byte(`{"count":1054,"next":"https://pokeapi.co/api/v2/location-area?offset=20&limit=20","previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"},{"name":"eterna-city-area","url":"https://pokeapi.co/api/v2/location-area/2/"},{"name":"pastoria-city-area","url":"https://pokeapi.co/api/v2/location-area/3/"},{"name":"sunyshore-city-area","url":"https://pokeapi.co/api/v2/location-area/4/"},{"name":"sinnoh-pokemon-league-area","url":"https://pokeapi.co/api/v2/location-area/5/"},{"name":"oreburgh-mine-1f","url":"https://pokeapi.co/api/v2/location-area/6/"},{"name":"oreburgh-mine-b1f","url":"https://pokeapi.co/api/v2/location-area/7/"},{"name":"valley-windworks-area","url":"https://pokeapi.co/api/v2/location-area/8/"},{"name":"eterna-forest-area","url":"https://pokeapi.co/api/v2/location-area/9/"},{"name":"fuego-ironworks-area","url":"https://pokeapi.co/api/v2/location-area/10/"},{"name":"mt-coronet-1f-route-207","url":"https://pokeapi.co/api/v2/location-area/11/"},{"name":"mt-coronet-2f","url":"https://pokeapi.co/api/v2/location-area/12/"},{"name":"mt-coronet-3f","url":"https://pokeapi.co/api/v2/location-area/13/"},{"name":"mt-coronet-exterior-snowfall","url":"https://pokeapi.co/api/v2/location-area/14/"},{"name":"mt-coronet-exterior-blizzard","url":"https://pokeapi.co/api/v2/location-area/15/"},{"name":"mt-coronet-4f","url":"https://pokeapi.co/api/v2/location-area/16/"},{"name":"mt-coronet-4f-small-room","url":"https://pokeapi.co/api/v2/location-area/17/"},{"name":"mt-coronet-5f","url":"https://pokeapi.co/api/v2/location-area/18/"},{"name":"mt-coronet-6f","url":"https://pokeapi.co/api/v2/location-area/19/"},{"name":"mt-coronet-1f-from-exterior","url":"https://pokeapi.co/api/v2/location-area/20/"}]}`),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/420",
			val: []byte(`{"encounter_method_rates":[{"encounter_method":{"name":"surf","url":"https://pokeapi.co/api/v2/encounter-method/5/"},"version_details":[{"rate":4,"version":{"name":"ruby","url":"https://pokeapi.co/api/v2/version/7/"}},{"rate":4,"version":{"name":"sapphire","url":"https://pokeapi.co/api/v2/version/8/"}},{"rate":4,"version":{"name":"emerald","url":"https://pokeapi.co/api/v2/version/9/"}}]}],"game_index":96,"id":420,"location":{"name":"hoenn-route-126","url":"https://pokeapi.co/api/v2/location/474/"},"name":"hoenn-route-126-underwater","names":[{"language":{"name":"en","url":"https://pokeapi.co/api/v2/language/9/"},"name":"Road 126 (underwater)"},{"language":{"name":"fr","url":"https://pokeapi.co/api/v2/language/5/"},"name":"Route 126 (sous l'eau)"}],"pokemon_encounters":[{"pokemon":{"name":"chinchou","url":"https://pokeapi.co/api/v2/pokemon/170/"},"version_details":[{"encounter_details":[{"chance":30,"condition_values":[],"max_level":30,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":20}],"max_chance":30,"version":{"name":"ruby","url":"https://pokeapi.co/api/v2/version/7/"}},{"encounter_details":[{"chance":30,"condition_values":[],"max_level":30,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":20}],"max_chance":30,"version":{"name":"sapphire","url":"https://pokeapi.co/api/v2/version/8/"}},{"encounter_details":[{"chance":30,"condition_values":[],"max_level":30,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":20}],"max_chance":30,"version":{"name":"emerald","url":"https://pokeapi.co/api/v2/version/9/"}}]},{"pokemon":{"name":"clamperl","url":"https://pokeapi.co/api/v2/pokemon/366/"},"version_details":[{"encounter_details":[{"chance":60,"condition_values":[],"max_level":30,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":20},{"chance":5,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30}],"max_chance":65,"version":{"name":"ruby","url":"https://pokeapi.co/api/v2/version/7/"}},{"encounter_details":[{"chance":60,"condition_values":[],"max_level":30,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":20},{"chance":5,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30}],"max_chance":65,"version":{"name":"sapphire","url":"https://pokeapi.co/api/v2/version/8/"}},{"encounter_details":[{"chance":60,"condition_values":[],"max_level":30,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":20},{"chance":5,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30}],"max_chance":65,"version":{"name":"emerald","url":"https://pokeapi.co/api/v2/version/9/"}}]},{"pokemon":{"name":"relicanth","url":"https://pokeapi.co/api/v2/pokemon/369/"},"version_details":[{"encounter_details":[{"chance":4,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30},{"chance":1,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30}],"max_chance":5,"version":{"name":"ruby","url":"https://pokeapi.co/api/v2/version/7/"}},{"encounter_details":[{"chance":4,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30},{"chance":1,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30}],"max_chance":5,"version":{"name":"sapphire","url":"https://pokeapi.co/api/v2/version/8/"}},{"encounter_details":[{"chance":4,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30},{"chance":1,"condition_values":[],"max_level":35,"method":{"name":"seaweed","url":"https://pokeapi.co/api/v2/encounter-method/27/"},"min_level":30}],"max_chance":5,"version":{"name":"emerald","url":"https://pokeapi.co/api/v2/version/9/"}}]}]}`),
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
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte(`{"count":1054,"next":"https://pokeapi.co/api/v2/location-area?offset=20&limit=20","previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"},{"name":"eterna-city-area","url":"https://pokeapi.co/api/v2/location-area/2/"},{"name":"pastoria-city-area","url":"https://pokeapi.co/api/v2/location-area/3/"},{"name":"sunyshore-city-area","url":"https://pokeapi.co/api/v2/location-area/4/"},{"name":"sinnoh-pokemon-league-area","url":"https://pokeapi.co/api/v2/location-area/5/"},{"name":"oreburgh-mine-1f","url":"https://pokeapi.co/api/v2/location-area/6/"},{"name":"oreburgh-mine-b1f","url":"https://pokeapi.co/api/v2/location-area/7/"},{"name":"valley-windworks-area","url":"https://pokeapi.co/api/v2/location-area/8/"},{"name":"eterna-forest-area","url":"https://pokeapi.co/api/v2/location-area/9/"},{"name":"fuego-ironworks-area","url":"https://pokeapi.co/api/v2/location-area/10/"},{"name":"mt-coronet-1f-route-207","url":"https://pokeapi.co/api/v2/location-area/11/"},{"name":"mt-coronet-2f","url":"https://pokeapi.co/api/v2/location-area/12/"},{"name":"mt-coronet-3f","url":"https://pokeapi.co/api/v2/location-area/13/"},{"name":"mt-coronet-exterior-snowfall","url":"https://pokeapi.co/api/v2/location-area/14/"},{"name":"mt-coronet-exterior-blizzard","url":"https://pokeapi.co/api/v2/location-area/15/"},{"name":"mt-coronet-4f","url":"https://pokeapi.co/api/v2/location-area/16/"},{"name":"mt-coronet-4f-small-room","url":"https://pokeapi.co/api/v2/location-area/17/"},{"name":"mt-coronet-5f","url":"https://pokeapi.co/api/v2/location-area/18/"},{"name":"mt-coronet-6f","url":"https://pokeapi.co/api/v2/location-area/19/"},{"name":"mt-coronet-1f-from-exterior","url":"https://pokeapi.co/api/v2/location-area/20/"}]}`))

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

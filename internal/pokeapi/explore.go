package pokeapi

import (
	"github.com/adamsma/pokedexcli/internal/pokecache"
)

func ExploreLocation(name string, c *pokecache.Cache) ([]Pokemon, error) {

	var exploreInfo Explore

	path := "location-area/" + name

	err := do(path, &exploreInfo, c)

	if err != nil {
		return nil, err
	}

	foundPoke := make([]Pokemon, len(exploreInfo.PokemonEncounters))

	for i, encounter := range exploreInfo.PokemonEncounters {
		foundPoke[i].Name = encounter.Pokemon.Name
		foundPoke[i].URL = encounter.Pokemon.URL
	}

	return foundPoke, nil
}

package pokeapi

import "github.com/adamsma/pokedexcli/internal/pokecache"

func GetPokemonInfo(name string, c *pokecache.Cache) (Pokemon, error) {

	var pokeInfo PokeData

	path := "pokemon/" + name

	err := do(path, &pokeInfo, c)

	if err != nil {
		return Pokemon{}, err
	}

	pkmn := Pokemon{
		Name: name,
		URL:  path,
		Data: pokeInfo,
	}

	return pkmn, nil
}

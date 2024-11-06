package pokeapi

import "github.com/adamsma/pokedexcli/internal/pokecache"

func GetLocations(qry string, c *pokecache.Cache) Location {

	var locations Location

	path := "location?" + qry

	do(path, &locations, c)
	return locations

}

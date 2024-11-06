package main

import (
	"fmt"
	"net/url"

	"github.com/adamsma/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {

	if cfg.Next == nil {
		return fmt.Errorf("<<End of Locations>>")
	}

	u, err := url.Parse(*cfg.Next)
	if err != nil {
		return err
	}

	locations := pokeapi.GetLocations(u.RawQuery, &cfg.pokeCache)

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}

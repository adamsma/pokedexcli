package main

import (
	"fmt"
	"net/url"

	"github.com/adamsma/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *config, params []string) error {

	if cfg.Previous == nil {
		return fmt.Errorf("<<Start of Locations>>")
	}

	u, err := url.Parse(*cfg.Previous)
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

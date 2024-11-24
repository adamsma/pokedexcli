package main

import (
	"fmt"

	"github.com/adamsma/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, params ...string) error {

	if len(params) != 1 {
		return fmt.Errorf("you must provide a location name, and only a location name")
	}

	fmt.Printf("Exploreing %s...\n", params[0])
	foundPoke, err := pokeapi.ExploreLocation(params[0], &cfg.pokeCache)

	if err != nil {
		fmt.Println("Unable to find location")
		return nil
	}

	fmt.Println("Found Pokemon")
	for _, pokemon := range foundPoke {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil

}

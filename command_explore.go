package main

import (
	"fmt"

	"github.com/adamsma/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, params []string) error {

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

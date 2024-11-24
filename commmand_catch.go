package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/adamsma/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, params ...string) error {

	if len(params) < 1 {
		return fmt.Errorf("you must provide the name of a pokemon you wish to catch")
	}

	name := params[0]

	pokemon, err := pokeapi.GetPokemonInfo(name, &cfg.pokeCache)

	if err != nil {
		return fmt.Errorf("failed to throw a pokeball or unknown pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	// at time of creation max base experience was 608
	throwSkill := rand.IntN(700)
	if throwSkill <= pokemon.Data.BaseExperience {
		fmt.Printf("%s escapsed!\n", name)
	} else {
		cfg.pokedex[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
	}

	return nil
}

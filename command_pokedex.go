package main

import "fmt"

func commandPokedex(cfg *config, params ...string) error {

	if len(params) > 0 {
		fmt.Println("ignoring additional parameters...")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}

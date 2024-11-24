package main

import "fmt"

func commandInspect(cfg *config, params ...string) error {

	if len(params) < 1 {
		return fmt.Errorf("you must provide the name of a pokemon to inspect")
	}

	if len(params) > 1 {
		fmt.Println("ignoring additional parameters...")
	}

	name := params[0]

	entry, caught := cfg.pokedex[name]
	info := entry.Data

	if !caught {
		return fmt.Errorf("no information found: you have not caught a %s", name)
	}

	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Height: %d\n", info.Height)
	fmt.Printf("Weight: %d\n", info.Weight)

	fmt.Printf("Stats:\n")
	for _, val := range info.Stats {
		fmt.Printf("  -%s: %d\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, val := range info.Types {
		fmt.Printf("  - %s\n", val.Type.Name)
	}

	return nil
}

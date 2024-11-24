package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adamsma/pokedexcli/internal/pokecache"
)

type config struct {
	Next      *string
	Previous  *string
	pokeCache pokecache.Cache
}

func main() {

	var cfg config
	ep := "https://pokeapi.co/api/v2/location?offset=0&limit=20"
	cfg.Next = &ep
	cacheRefresh, _ := time.ParseDuration("10s")
	cfg.pokeCache = pokecache.NewCache(cacheRefresh)

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := strings.ToLower(reader.Text())
		words := strings.Fields(input)

		if len(words) == 0 {
			continue
		}

		cmd := words[0]
		params := []string{}

		if len(words) > 1 {
			params = words[1:]
		}

		cliCmd, ok := getCommands()[cmd]
		if ok {

			err := cliCmd.callback(&cfg, params...)
			if err != nil {
				fmt.Println(err)
			}
			continue

		} else {

			fmt.Printf("Unknown command: %s\n", cmd)
			continue

		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the names of 20 loaction areas in the Pokemon world",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 loaction areas in the Pokemon world",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore <location_name>",
			description: "Takes an area name as a parameter and displays list of pokemon found in the given area;",
			callback:    commandExplore,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

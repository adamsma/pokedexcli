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
		input = strings.Fields(input)[0]
		if len(input) == 0 {
			continue
		}

		cliCmd, ok := getCommands()[input]
		if ok {

			err := cliCmd.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue

		} else {

			fmt.Printf("Unknown command: %s\n", input)
			continue

		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

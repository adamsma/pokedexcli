package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := strings.ToLower(reader.Text())
		input = strings.Fields(input)[0]
		if len(input) == 0 {
			continue
		}

		// fmt.Printf("Command Recieved: %s\n", input)
		cliCmd, ok := getCommands()[input]
		if ok {

			err := cliCmd.callback()
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
	callback    func() error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

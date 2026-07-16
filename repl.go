package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shnavo/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
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
			description: "Displays all available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the 20 locations from Pokemon. Call the command again to see the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations from Pokemon. Call the command again to see the previous 20 locations.",
			callback:    commandMapBack,
		},
	}
}

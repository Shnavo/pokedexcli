package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
			err := command.callback()
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
	callback    func() error
	config		*Config
}

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var config Config = Config{
	Next: "https://pokeapi.co/api/v2/location-area/",
} 

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays all available commands",
			callback:    commandHelp,
			config:      &config,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
			config:      &config,
		},
		"map": {
			name:        "map",
			description: "Displays the 20 locations from Pokemon. Call the command again to see the next 20 locations.",
			callback:    commandMap,
			config:      &config,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations from Pokemon. Call the command again to see the previous 20 locations.",
			callback:    commandMapBack,
			config:      &config,
		},
	}
}

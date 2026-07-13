package main

import (
	"strings"
    "fmt"
    "bufio"
    "os"
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

    }
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}
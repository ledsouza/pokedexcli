package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ledsouza/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")
		scanner.Scan()

		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		inputCommand := cleaned[0]

		commands := getCommands()
		command, ok := commands[inputCommand]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays a list of the next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays a list of the previous locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

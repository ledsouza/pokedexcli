package main

import "fmt"

func commandHelp(cfg *config) error {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, cmd := range commands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

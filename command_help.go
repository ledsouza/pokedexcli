package main

import "fmt"

func commandHelp() {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range commands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
}

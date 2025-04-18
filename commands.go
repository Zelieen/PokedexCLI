package main

import (
	"fmt"
	"internal/pokeAPI"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(params *config) error
}

type config struct {
	next     string
	previous string
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
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}

func commandHelp(params *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(params *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(params *config) error {
	nex, prev := pokeAPI.GetLocationAreas(params.next)
	params.next = nex
	params.previous = prev
	return nil
}

func commandMapBack(params *config) error {
	if params.previous == "" {
		fmt.Println("you're on the first page, you can not go back")
		return nil
	}
	nex, prev := pokeAPI.GetLocationAreas(params.previous)
	params.next = nex
	params.previous = prev
	return nil
}

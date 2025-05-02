package main

import (
	"fmt"
	pokeAPI "internal/pokeAPI"
	pokecache "internal/pokeCache"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(input []string, params *config) error
}

type config struct {
	next     string
	previous string
	cache    pokecache.Cache
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
		"input": {
			name:        "input",
			description: "Prints the registered input for testing",
			callback:    commandInput,
		},
		"explore": {
			name:        "explore",
			description: "lists all the Pokemon of the specified location (explore <location from map>)",
			callback:    commandExplore,
		},
	}
}

func commandHelp(input []string, params *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(input []string, params *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(input []string, params *config) error {
	url := params.next
	return mapCommmandLogic(params, url)
}

func commandMapBack(input []string, params *config) error {
	url := params.previous
	if url == "" {
		fmt.Println("you're on the first page, you can not go back")
		return nil
	}
	return mapCommmandLogic(params, url)
}

func mapCommmandLogic(params *config, url string) error {
	cached, ok := params.cache.Get(url)
	if !ok {
		//fmt.Println("not cached, trying API call")
		new, err := pokeAPI.MakeAPICall(url)
		if err != nil {
			return err
		}
		params.cache.Add(url, new)
		//fmt.Println("adding cache from API call")
		cached = new
	}

	area := pokeAPI.GetAreas(cached)
	params.next = area.Next
	params.previous = area.Previous

	pokeAPI.PrintLocationAreas(area)
	return nil
}

func commandInput(input []string, params *config) error {
	fmt.Println("These input words were received:")
	for _, w := range input {
		fmt.Println(w)
	}
	return nil
}

func commandExplore(input []string, params *config) error {
	if len(input) < 2 {
		fmt.Println("please provide a location after the explore command")
		return nil
	}
	fmt.Println(pokeAPI.ConstructURL("location") + input[1] + "/")
	return nil
}

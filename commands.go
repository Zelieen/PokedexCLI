package main

import (
	"fmt"
	"internal/pokeAPI"
	pokecache "internal/pokeCache"
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
	url := params.next
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

func commandMapBack(params *config) error {
	url := params.previous
	if url == "" {
		fmt.Println("you're on the first page, you can not go back")
		return nil
	}
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

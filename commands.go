package main

import (
	"fmt"
	pokeAPI "internal/pokeAPI"
	pokecache "internal/pokeCache"
	"math/rand"
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
	dex      map[string]pokeAPI.Pokemon
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
			name:        "explore <location from map>",
			description: "lists all the Pokemon of the specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Throws a pokeball at the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Print information on a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all pokemon in your pokedex",
			callback:    commandPokedex,
		},
		"save": {
			name:        "save",
			description: "Saves your progress",
			callback:    commandSave,
		},
		"load": {
			name:        "load",
			description: "Loads your progress",
			callback:    commandLoad,
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
	body := useCache(params, url)

	area := pokeAPI.GetAreas(body)
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
	//fmt.Println(pokeAPI.ConstructURL("location") + input[1] + "/")
	url := pokeAPI.ConstructURL("location") + input[1] + "/"
	area := useCache(params, url)
	if area == nil {
		fmt.Println("no such location found")
		return nil
	}
	list := pokeAPI.GetPokemonList(area)
	fmt.Printf("Pokemon in %s:", input[1])
	pokeAPI.PrintList(list)
	return nil
}

func useCache(params *config, url string) []byte {
	cached, ok := params.cache.Get(url)
	if !ok {
		//fmt.Println("not cached, trying API call")
		new, err := pokeAPI.MakeAPICall(url)
		if err != nil {
			return nil
		}
		params.cache.Add(url, new)
		//fmt.Println("adding cache from API call")
		cached = new
	}
	return cached
}

func commandCatch(input []string, params *config) error {
	if len(input) < 2 {
		fmt.Println("please provide a pokemon name")
		return nil
	}
	//fmt.Println(pokeAPI.ConstructURL("pokemon") + input[1] + "/")
	url := pokeAPI.ConstructURL("pokemon") + input[1] + "/"
	poke := useCache(params, url)
	if poke == nil {
		fmt.Println("there is no such pokemon")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", input[1])
	pokemon := pokeAPI.GetPokemonInfo(poke)
	// Catchrate: BaseExperience / (BaseExperience + 50) (min/max BaseExperinece: 36/608)
	cr := float32(pokemon.BaseExperience) / float32(pokemon.BaseExperience+80)
	chance := rand.Float32()
	if chance > cr {
		fmt.Printf("The %s was caught!!\n", pokemon.Name)
		//fmt.Printf("Caught: %v vs %v\n", chance, cr)
		(params.dex)[pokemon.Name] = pokemon
	} else {
		fmt.Printf("Aww... %s broke free.\n", pokemon.Name)
		//fmt.Printf("Aww.. (%v vs %v)\n", chance, cr)
	}
	return nil
}

func commandInspect(input []string, params *config) error {
	if len(input) < 2 {
		fmt.Println("please provide a pokemon name")
		return nil
	}
	pokemon, ok := (params.dex)[input[1]]
	if !ok {
		fmt.Printf("There is no data on %s\n", input[1])
		return nil
	}
	pokeAPI.PrintPokemon(pokemon)

	return nil
}

func commandPokedex(input []string, params *config) error {
	if len(params.dex) < 1 {
		fmt.Println("Your pokedex is empty.")
		return nil
	}
	for poke := range params.dex {
		fmt.Println(poke)
	}
	return nil
}

func commandSave(input []string, params *config) error {
	//save the params_config to disk
	fmt.Println("Saved your progress")
	return nil
}

func commandLoad(input []string, params *config) error {
	//load the params_config to disk
	fmt.Println("Loading your progress")
	return nil
}

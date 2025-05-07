package main

import (
	"bufio"
	"fmt"
	"internal/pokeAPI"
	pokecache "internal/pokeCache"
	"os"
	"strings"
	"time"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	params_config := config{
		cache: pokecache.NewCache(3 * time.Minute),
	}
	dex := CreatePokedex()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleanedInput := cleanInput(input)
		if len(cleanedInput) < 1 {
			continue
		}
		//fmt.Printf("Your command was: %s\n", command)
		command, ok := getCommands()[cleanedInput[0]]
		if ok {
			ExecuteCommand(command.callback, cleanedInput, &params_config, &dex)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	var cleanedInput []string
	cleanedWords := strings.Split(strings.ToLower(text), " ")
	for _, word := range cleanedWords {
		if word != "" {
			cleanedInput = append(cleanedInput, word)
		}
	}
	return cleanedInput
}

func ExecuteCommand(f func(input []string, params *config, dex *map[string]pokeAPI.Pokemon) error, input []string, params_pointer *config, dex *map[string]pokeAPI.Pokemon) error {
	err := f(input, params_pointer, dex)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

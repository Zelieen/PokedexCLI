package main

import (
	"bufio"
	"fmt"
	pokecache "internal/pokeCache"
	"os"
	"strings"
	"time"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	params_config := config{
		cache: pokecache.NewCache(3 * time.Minute),
		Dex:   CreatePokedex(),
	}

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
			ExecuteCommand(command.callback, cleanedInput, &params_config)
		} else {
			fmt.Println("Unknown command. Type 'help' for available commands.")
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

func ExecuteCommand(f func(input []string, params *config) error, input []string, params_pointer *config) error {
	err := f(input, params_pointer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

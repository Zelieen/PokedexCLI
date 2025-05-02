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
		cache: pokecache.NewCache(1 * time.Minute),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		commandInput := cleanInput(input)[0]
		//fmt.Printf("Your command was: %s\n", command)
		command, ok := getCommands()[commandInput]
		if ok {
			ExecuteCommand(command.callback, &params_config)
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

func ExecuteCommand(f func(params *config) error, params_pointer *config) error {
	err := f(params_pointer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

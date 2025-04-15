package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		firstWord := cleanInput(input)[0]
		fmt.Printf("Your command was: %s\n", firstWord)
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

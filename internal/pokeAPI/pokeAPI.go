package pokeAPI

import (
	"fmt"
	"io"
	"net/http"
	//"encoding/json"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func PokeLocationArea() {
	fmt.Println("But PokeAPI function says hi.")

	//area := LocationArea{}

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/?limit=10&offset=0")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode > 299 {
		fmt.Printf("Response failed. Status code is %d\nThe response body:\n%s", res.StatusCode, body)
		return
	}
	fmt.Printf("The response body:\n%s\n", body)
}

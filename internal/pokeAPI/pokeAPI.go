package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func MakeAPICall(url string) ([]byte, error) {
	getUrl := url
	if getUrl == "" {
		getUrl = baseURL + "/location-area/?limit=20&offset=0"
	}

	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if res.StatusCode > 299 {
		fmt.Printf("Response failed. Status code is %d\n", res.StatusCode)
		return nil, err
	}

	return body, nil
}

func GetAreas(body []byte) LocationArea {
	area := LocationArea{}
	err := json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
		return area
	}

	return area
}

func PrintLocationAreas(areas LocationArea) {
	for _, place := range areas.Results {
		fmt.Println(place.Name)
	}
}

func ConstructURL(target string) string {
	endpoints := map[string]string{
		"location": "/location-area/",
		"pokemon":  "/pokemon/",
	}

	_, ok := endpoints[target]
	if !ok {
		return ""
	}
	return baseURL + endpoints[target]
}

func GetPokemonList(body []byte) []string {
	area := DetailedArea{}
	err := json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	pokeList := make([]string, 1)
	for _, pokemon := range area.PokemonEncounters {
		pokeList = append(pokeList, pokemon.Pokemon.Name)
	}
	return pokeList
}

func GetPokemonInfo(body []byte) Pokemon {
	pokemon := Pokemon{}
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println(err)
		return pokemon
	}

	return pokemon
}

func PrintList(list []string) {
	for _, item := range list {
		fmt.Println(item)
	}
}

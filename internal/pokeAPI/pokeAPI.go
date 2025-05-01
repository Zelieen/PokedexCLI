package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	//pokecache "internal/pokeCache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
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

func GetLocationAreasAPI(url string) LocationArea {
	area := LocationArea{}
	getUrl := url
	if getUrl == "" {
		getUrl = baseURL + "/location-area/?limit=20&offset=0"
	}

	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println(err)
		return area
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return area
	}
	if res.StatusCode > 299 {
		fmt.Printf("Response failed. Status code is %d\n", res.StatusCode)
		return area
	}

	err = json.Unmarshal(body, &area)
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

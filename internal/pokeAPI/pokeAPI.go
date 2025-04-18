package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func PokeLocationArea(next_url string) {
	fmt.Println("Not fully implemented. But PokeAPI function says hi.")
	getUrl := "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	if next_url != "" {
		getUrl = next_url
	}

	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.StatusCode > 299 {
		fmt.Printf("Response failed. Status code is %d\n", res.StatusCode)
		return
	}

	area := LocationArea{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The next area url:\n")
	fmt.Println(area.Next)
}

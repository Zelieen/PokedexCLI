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

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (next string, previous string) {
	getUrl := url
	if getUrl == "" {
		getUrl = baseURL + "/location-area/?limit=20&offset=0"
	}

	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	if res.StatusCode > 299 {
		fmt.Printf("Response failed. Status code is %d\n", res.StatusCode)
		return "", ""
	}

	area := LocationArea{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	for _, place := range area.Results {
		fmt.Println(place.Name)
	}
	return area.Next, area.Previous
}

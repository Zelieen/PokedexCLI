package main

import (
	pokeAPI "internal/pokeAPI"
)

func CreatePokedex() map[string]pokeAPI.Pokemon {
	return make(map[string]pokeAPI.Pokemon)
}

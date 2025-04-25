module github.com/zelieen/pokedexcli

go 1.23.2

require internal/pokeAPI v0.0.0-unpublished
replace internal/pokeAPI v0.0.0-unpublished => ./internal/pokeAPI

require internal/pokeCache v0.0.0-unpublished
replace internal/pokeCache v0.0.0-unpublished => ./internal/pokeCache
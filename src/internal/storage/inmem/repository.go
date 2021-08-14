package inmem

import (
	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
)

type repository struct{}

// NewRepository initialize repository
func NewRepository() gopoke.PokemonRepo {
	return &repository{}
}

// GetPokemons fetch data from memory
func (r *repository) GetPokemons(limit, offset int) ([]gopoke.PokemonSimple, error) {
	return []gopoke.PokemonSimple{
		gopoke.NewPokemonSimple(
			"bulbasaur",
			"https://pokeapi.co/api/v2/pokemon/1/",
		),
		gopoke.NewPokemonSimple(
			"charmander",
			"https://pokeapi.co/api/v2/pokemon/4/",
		),
		gopoke.NewPokemonSimple(
			"squrtle",
			"https://pokeapi.co/api/v2/pokemon/7/",
		),
	}, nil
}

// GetPokemonById fetch data from memory
func (r *repository) GetPokemonById(pokemonID int) (gopoke.Pokemon, error) {
	return gopoke.Pokemon(
		gopoke.NewPokemon(
			25,
			"pikachu",
			112,
			4,
			true,
			35,
			60,
		),
	), nil
}

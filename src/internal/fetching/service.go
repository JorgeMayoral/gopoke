package fetching

import (
	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
)

// Service provides pokemon fetching operations
type Service interface {
	// FetchPokemons fetch all pokemons from repository
	FetchPokemons(limit, offset int) ([]gopoke.PokemonSimple, error)
	// FetchByID fetch the pokemon that match the given id
	FetchByID(id int) (gopoke.Pokemon, error)
}

type service struct {
	pR gopoke.PokemonRepo
}

// NewService creates an adding service with the necessary dependencies
func NewService(r gopoke.PokemonRepo) Service {
	return &service{r}
}

func (s *service) FetchPokemons(limit, offset int) ([]gopoke.PokemonSimple, error) {
	return s.pR.GetPokemons(limit, offset)
}

func (s *service) FetchByID(id int) (gopoke.Pokemon, error) {
	return s.pR.GetPokemonById(id)
}

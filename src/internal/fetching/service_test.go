package fetching_test

import (
	"errors"
	"testing"

	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
	"github.com/JorgeMayoral/gopoke/src/internal/fetching"
	mock "github.com/JorgeMayoral/gopoke/src/internal/storage/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFetchById(t *testing.T) {
	tests := map[string]struct {
		repo  gopoke.PokemonRepo
		input int
		want  int
		err   error
	}{
		"valid pokemon":         {repo: buildMockPokemons(), input: 25, want: 25, err: nil},
		"not found pokemon":     {repo: buildMockPokemons(), input: 999999, err: errors.New("error")},
		"error with repository": {repo: buildMockError(), err: errors.New("error")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			service := fetching.NewService(tc.repo)

			b, err := service.FetchByID(tc.input)

			if tc.err != nil {
				assert.Error(t, err)
			}

			if tc.err == nil {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.want, b.PokemonID)
		})
	}
}

func buildMockPokemons() gopoke.PokemonRepo {
	mockedRepo := &mock.PokemonRepoMock{
		GetPokemonByIdFunc: func(pokemonID int) (gopoke.Pokemon, error) {
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
		},
	}

	return mockedRepo
}

func buildMockError() gopoke.PokemonRepo {
	mockedRepo := &mock.PokemonRepoMock{
		GetPokemonByIdFunc: func(pokemonID int) (gopoke.Pokemon, error) {
			return gopoke.Pokemon{}, errors.New("error")
		},
	}

	return mockedRepo
}

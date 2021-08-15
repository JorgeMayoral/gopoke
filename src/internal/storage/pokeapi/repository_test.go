package pokeapi_test

import (
	"testing"

	"github.com/JorgeMayoral/gopoke/src/internal/storage/pokeapi"
)

func BenchmarkGetPokemons(b *testing.B) {
	repo := pokeapi.NewPokeapiRepository()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		repo.GetPokemons(2000, 0)
	}
}

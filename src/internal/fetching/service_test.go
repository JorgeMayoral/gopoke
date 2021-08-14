package fetching_test

import (
	"testing"

	"github.com/JorgeMayoral/gopoke/src/internal/fetching"
	"github.com/JorgeMayoral/gopoke/src/internal/storage/inmem"
)

func TestFetchById(t *testing.T) {
	repo := inmem.NewRepository()
	service := fetching.NewService(repo)

	expected := 25
	b, err := service.FetchByID(expected)
	if err != nil {
		t.Fatalf("expected %d, got an error %v", expected, err)
	}

	if b.PokemonID != expected {
		t.Fatalf("expected %d, got: %d", expected, b.PokemonID)
	}
}

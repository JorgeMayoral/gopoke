package fetching_test

import (
	"errors"
	"testing"

	"github.com/JorgeMayoral/gopoke/src/internal/fetching"
	"github.com/JorgeMayoral/gopoke/src/internal/storage/inmem"
)

func TestFetchById(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
		err   error
	}{
		"valid pokemon":     {input: 25, want: 25, err: nil},
		"not found pokemon": {input: 999999, err: errors.New("error")},
	}
	repo := inmem.NewRepository()
	service := fetching.NewService(repo)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := service.FetchByID(tc.input)
			if err != nil && tc.err == nil {
				t.Fatalf("not expected any errors and got %v", err)
			}

			if err == nil && tc.err != nil {
				t.Error("expected an error and got nil")
			}

			if b.PokemonID != tc.want {
				t.Fatalf("expected %d, got: %d", tc.want, b.PokemonID)
			}
		})
	}
}

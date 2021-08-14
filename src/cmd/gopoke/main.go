package main

import (
	_ "flag"

	"github.com/JorgeMayoral/gopoke/src/internal/fetching"
	"github.com/JorgeMayoral/gopoke/src/internal/storage/pokeapi"

	gopoke "github.com/JorgeMayoral/gopoke/src/internal"
	"github.com/JorgeMayoral/gopoke/src/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	var repo gopoke.PokemonRepo = pokeapi.NewPokeapiRepository()

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "gopoke"}
	rootCmd.AddCommand(cli.InitPokemonsCmd(fetchingService))
	rootCmd.Execute()
}
